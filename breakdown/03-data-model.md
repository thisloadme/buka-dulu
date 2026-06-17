# Data Model & Schema — BukaDulu MVP
## Versi 1.0 | SQLite (Dev) → PostgreSQL (Prod)

---

## 1. Entity-Relationship Overview

```
users 1──N ventures 1──N missions 1──N evidences 1──N evidence_reviews
                    │                                       │
                    │ 1──N menus                             │
                    │ 1──N ingredients                        │
                    │ 1──N customer_segments                  │
                    │ 1──N scores                             │
                    │ 1──N decisions                           │
                    │ 1──N venture_versions                    │
                    │ 1──N notifications                       │
                    │ 1──N audit_logs                          │
```

---

## 2. Full Schema (PostgreSQL)

### 2.1 Users & Roles

```sql
CREATE TYPE user_role AS ENUM ('founder', 'mentor', 'admin');

CREATE TABLE users (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    role            user_role NOT NULL DEFAULT 'founder',
    full_name       VARCHAR(255) NOT NULL,
    email           VARCHAR(255) UNIQUE,
    phone           VARCHAR(20) UNIQUE,
    password_hash   VARCHAR(255) NOT NULL,
    avatar_url      VARCHAR(512),
    status          VARCHAR(20) NOT NULL DEFAULT 'active'
                        CHECK (status IN ('active', 'suspended', 'deleted')),
    last_login_at   TIMESTAMPTZ,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_users_email ON users(email) WHERE email IS NOT NULL;
CREATE INDEX idx_users_phone ON users(phone) WHERE phone IS NOT NULL;
```

### 2.2 Ventures

```sql
CREATE TYPE venture_stage AS ENUM (
    'draft',
    'idea_defined',
    'customer_defined',
    'sku_focused',
    'cost_evaluated',
    'mission_active',
    'evidence_submitted',
    'evidence_reviewed',
    'ready_to_decide',
    'continue',
    'repeat',
    'pivot',
    'stop'
);

CREATE TABLE ventures (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    owner_user_id   UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name            VARCHAR(255) NOT NULL,
    category        VARCHAR(100),
    region          VARCHAR(100),
    stage           venture_stage NOT NULL DEFAULT 'draft',
    current_version INT NOT NULL DEFAULT 1,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_ventures_owner ON ventures(owner_user_id);
CREATE INDEX idx_ventures_stage ON ventures(stage);

CREATE TABLE venture_versions (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    venture_id      UUID NOT NULL REFERENCES ventures(id) ON DELETE CASCADE,
    version         INT NOT NULL,
    snapshot        JSONB NOT NULL,  -- full snapshot of idea/customer/menu at that version
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(venture_id, version)
);
```

### 2.3 Idea & Customer

```sql
CREATE TABLE ideas (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    venture_id      UUID NOT NULL REFERENCES ventures(id) ON DELETE CASCADE,
    raw_input       TEXT NOT NULL,
    one_line_concept VARCHAR(500),
    target_customer TEXT,
    value_proposition TEXT,
    key_assumptions TEXT,
    early_risks     TEXT,
    version         INT NOT NULL DEFAULT 1,
    is_locked       BOOLEAN NOT NULL DEFAULT FALSE,
    ai_raw_input    JSONB,  -- save AI raw request for debugging
    ai_raw_output   JSONB,  -- save AI raw response for debugging
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(venture_id, version)
);

CREATE TABLE customer_segments (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    venture_id      UUID NOT NULL REFERENCES ventures(id) ON DELETE CASCADE,
    name            VARCHAR(255) NOT NULL,
    age_range       VARCHAR(50),
    buy_context     TEXT,
    budget_range    VARCHAR(100),
    location        VARCHAR(255),
    consumption_moment TEXT,
    is_too_general  BOOLEAN NOT NULL DEFAULT FALSE,
    is_locked       BOOLEAN NOT NULL DEFAULT FALSE,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

### 2.4 Menu

```sql
CREATE TYPE menu_status AS ENUM ('candidate', 'active', 'deferred', 'dropped');

CREATE TABLE menus (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    venture_id      UUID NOT NULL REFERENCES ventures(id) ON DELETE CASCADE,
    name            VARCHAR(255) NOT NULL,
    description     TEXT,
    status          menu_status NOT NULL DEFAULT 'candidate',
    is_hero         BOOLEAN NOT NULL DEFAULT FALSE,
    complexity_score DECIMAL(5,2),
    complexity_factors JSONB,  -- {ingredient, prep_time, spoilage, process_variation}
    sort_order      INT NOT NULL DEFAULT 0,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_menus_venture ON menus(venture_id);

-- Enforce max 3 active SKU per venture
CREATE OR REPLACE FUNCTION check_active_sku_limit()
RETURNS TRIGGER AS $$
BEGIN
    IF (NEW.status = 'active') THEN
        IF (SELECT COUNT(*) FROM menus
            WHERE venture_id = NEW.venture_id AND status = 'active' AND id != NEW.id
        ) >= 3 THEN
            RAISE EXCEPTION 'Maximum 3 active SKU per venture';
        END IF;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_menu_active_limit
    AFTER INSERT OR UPDATE ON menus
    FOR EACH ROW WHEN (NEW.status = 'active')
    EXECUTE FUNCTION check_active_sku_limit();
```

### 2.5 Cost/Ingredients

```sql
CREATE TABLE ingredients (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    venture_id      UUID NOT NULL REFERENCES ventures(id) ON DELETE CASCADE,
    menu_id         UUID NOT NULL REFERENCES menus(id) ON DELETE CASCADE,
    name            VARCHAR(255) NOT NULL,
    unit            VARCHAR(50) NOT NULL,   -- gram, pcs, ml, sdm, etc.
    quantity        DECIMAL(10,2) NOT NULL,
    unit_price      DECIMAL(12,2) NOT NULL,  -- price per unit
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_ingredients_menu ON ingredients(menu_id);

CREATE TABLE packaging_costs (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    venture_id      UUID NOT NULL REFERENCES ventures(id) ON DELETE CASCADE,
    menu_id         UUID NOT NULL REFERENCES menus(id) ON DELETE CASCADE,
    name            VARCHAR(255) NOT NULL,
    unit_price      DECIMAL(12,2) NOT NULL,
    quantity_per_pcs DECIMAL(10,2) NOT NULL DEFAULT 1,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE cost_summaries (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    venture_id      UUID NOT NULL REFERENCES ventures(id) ON DELETE CASCADE,
    menu_id         UUID NOT NULL REFERENCES menus(id) ON DELETE CASCADE,
    hpp_per_porsi   DECIMAL(12,2) NOT NULL,
    suggested_price DECIMAL(12,2) NOT NULL,
    target_margin   DECIMAL(5,2) NOT NULL,  -- percentage
    gross_margin    DECIMAL(5,2) NOT NULL,  -- percentage
    margin_status   VARCHAR(20) NOT NULL CHECK (margin_status IN ('sehat', 'tipis', 'berbahaya')),
    break_even_unit INT NOT NULL,  -- minimum units per month to break even
    labor_per_unit  DECIMAL(12,2),
    overhead_per_unit DECIMAL(12,2),
    is_locked       BOOLEAN NOT NULL DEFAULT FALSE,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(venture_id, menu_id)
);
```

### 2.6 Missions

```sql
CREATE TYPE mission_priority AS ENUM ('high', 'medium', 'low');
CREATE TYPE mission_status AS ENUM ('pending', 'accepted', 'in_progress', 'completed', 'skipped');
CREATE TYPE mission_type AS ENUM ('polling', 'pre_order', 'sampling', 'titip_jual', 'interview', 'harga_test', 'observation');

CREATE TABLE missions (
    id                  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    venture_id          UUID NOT NULL REFERENCES ventures(id) ON DELETE CASCADE,
    title               VARCHAR(255) NOT NULL,
    description         TEXT NOT NULL,
    mission_type        mission_type NOT NULL,
    priority            mission_priority NOT NULL DEFAULT 'medium',
    status              mission_status NOT NULL DEFAULT 'pending',
    due_at              TIMESTAMPTZ,
    evidence_required   TEXT,  -- description of what evidence is needed
    estimated_minutes   INT,
    created_by          UUID REFERENCES users(id),
    sort_order          INT NOT NULL DEFAULT 0,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_missions_venture ON missions(venture_id);
CREATE INDEX idx_missions_status ON missions(status);

CREATE TABLE mission_templates (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            VARCHAR(255) NOT NULL,
    description     TEXT NOT NULL,
    mission_type    mission_type NOT NULL,
    priority        mission_priority NOT NULL DEFAULT 'medium',
    estimated_minutes INT,
    evidence_required TEXT,
    applicable_stages venture_stage[],
    risk_profile_filters JSONB,  -- filter criteria
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

### 2.7 Evidence

```sql
CREATE TYPE evidence_type AS ENUM ('image', 'text', 'link');
CREATE TYPE verdict AS ENUM ('valid', 'weak', 'invalid', 'suspicious', 'pending');
CREATE TYPE review_source AS ENUM ('ai', 'human_override');

CREATE TABLE evidences (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    venture_id      UUID NOT NULL REFERENCES ventures(id) ON DELETE CASCADE,
    mission_id      UUID NOT NULL REFERENCES missions(id) ON DELETE CASCADE,
    uploader_id     UUID NOT NULL REFERENCES users(id),
    evidence_type   evidence_type NOT NULL,
    storage_url     VARCHAR(1024),
    text_content    TEXT,
    thumbnail_url   VARCHAR(1024),
    file_size_bytes INT,
    mime_type       VARCHAR(100),
    submitted_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_evidences_mission ON evidences(mission_id);
CREATE INDEX idx_evidences_venture ON evidences(venture_id);

CREATE TABLE evidence_reviews (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    evidence_id     UUID NOT NULL REFERENCES evidences(id) ON DELETE CASCADE,
    reviewer_type   review_source NOT NULL,
    verdict         verdict NOT NULL DEFAULT 'pending',
    score           DECIMAL(5,2),  -- 0-100
    rationale       TEXT NOT NULL,
    next_action     VARCHAR(50),  -- continue, repeat, pivot
    overridden_by   UUID REFERENCES users(id),  -- if human override
    overridden_at   TIMESTAMPTZ,
    processing_time_ms INT,  -- how long AI took
    ai_raw_output   JSONB,  -- raw AI response for debugging
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_reviews_evidence ON evidence_reviews(evidence_id);
```

### 2.8 Scores & Decisions

```sql
CREATE TABLE scores (
    id                  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    venture_id          UUID NOT NULL REFERENCES ventures(id) ON DELETE CASCADE,
    clarity_score       DECIMAL(5,2) NOT NULL DEFAULT 0,
    focus_score         DECIMAL(5,2) NOT NULL DEFAULT 0,
    economics_score     DECIMAL(5,2) NOT NULL DEFAULT 0,
    execution_score     DECIMAL(5,2) NOT NULL DEFAULT 0,
    evidence_score      DECIMAL(5,2) NOT NULL DEFAULT 0,
    market_response_score DECIMAL(5,2) NOT NULL DEFAULT 0,
    total_score         DECIMAL(5,2) NOT NULL DEFAULT 0,
    is_final            BOOLEAN NOT NULL DEFAULT FALSE,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_scores_venture ON scores(venture_id);

CREATE TYPE decision_type AS ENUM ('continue', 'repeat', 'pivot', 'stop');

CREATE TABLE decisions (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    venture_id      UUID NOT NULL REFERENCES ventures(id) ON DELETE CASCADE,
    decision        decision_type NOT NULL,
    rationale       TEXT NOT NULL,
    score_snapshot  JSONB NOT NULL,  -- copy of score components at decision time
    triggered_by    VARCHAR(100) NOT NULL DEFAULT 'system',  -- 'system', 'mentor', 'admin'
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_decisions_venture ON decisions(venture_id);
```

### 2.9 Notifications & Audit

```sql
CREATE TYPE notification_type AS ENUM (
    'mission_new', 'mission_deadline', 'review_complete',
    'score_update', 'decision_made', 'mentor_comment',
    'reactivation_prompt'
);

CREATE TABLE notifications (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id         UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    venture_id      UUID REFERENCES ventures(id) ON DELETE CASCADE,
    type            notification_type NOT NULL,
    title           VARCHAR(255) NOT NULL,
    body            TEXT,
    is_read         BOOLEAN NOT NULL DEFAULT FALSE,
    metadata        JSONB,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_notifications_user ON notifications(user_id);
CREATE INDEX idx_notifications_unread ON notifications(user_id, is_read) WHERE is_read = FALSE;

CREATE TABLE audit_logs (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    venture_id      UUID REFERENCES ventures(id),
    user_id         UUID REFERENCES users(id),
    action          VARCHAR(255) NOT NULL,
    entity_type     VARCHAR(100),
    entity_id       UUID,
    old_value       JSONB,
    new_value       JSONB,
    ip_address      VARCHAR(45),
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_audit_venture ON audit_logs(venture_id);
CREATE INDEX idx_audit_created ON audit_logs(created_at DESC);
```

### 2.10 Mentor

```sql
CREATE TABLE mentor_assignments (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    mentor_id       UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    founder_id      UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    venture_id      UUID NOT NULL REFERENCES ventures(id) ON DELETE CASCADE,
    status          VARCHAR(20) NOT NULL DEFAULT 'active'
                        CHECK (status IN ('active', 'archived')),
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(mentor_id, venture_id)
);

CREATE TABLE mentor_comments (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    mentor_id       UUID NOT NULL REFERENCES users(id),
    venture_id      UUID NOT NULL REFERENCES ventures(id) ON DELETE CASCADE,
    mission_id      UUID REFERENCES missions(id),
    evidence_id     UUID REFERENCES evidences(id),
    content         TEXT NOT NULL,
    is_deleted      BOOLEAN NOT NULL DEFAULT FALSE,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

---

## 3. SQLite Adaptation Notes

Untuk development (SQLite):

| PostgreSQL | SQLite |
|---|---|
| `UUID PRIMARY KEY DEFAULT gen_random_uuid()` | `TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(16))))` atau pake Go (generate UUID di kode) |
| `TIMESTAMPTZ` | `TEXT` (ISO 8601) atau `INTEGER` (Unix timestamp) |
| `ENUM` | `TEXT` dengan CHECK constraint |
| `JSONB` | `TEXT` (store JSON string) |
| `CREATE INDEX .. WHERE` | Tidak didukung — buat index biasa saja |
| `ON DELETE CASCADE` | Didukung |
| `gen_random_uuid()` | Tidak ada — generate di Go dengan `google/uuid` |

**Rekomendasi:** Generate UUID di Go layer (pakai `github.com/google/uuid`) daripada di DB — kompatibel SQLite & PG.

---

## 4. Migration Plan

Gunakan `golang-migrate`:

```
migrations/
├── 001_create_users.up.sql
├── 001_create_users.down.sql
├── 002_create_ventures.up.sql
├── 002_create_ventures.down.sql
├── 003_create_ideas.up.sql
├── 003_create_ideas.down.sql
├── 004_create_customer_segments.up.sql
├── 004_create_customer_segments.down.sql
├── 005_create_menus.up.sql
├── 005_create_menus.down.sql
├── 006_create_ingredients.up.sql
├── 006_create_ingredients.down.sql
├── 007_create_missions.up.sql
├── 007_create_missions.down.sql
├── 008_create_evidences.up.sql
├── 008_create_evidences.down.sql
├── 009_create_scores.up.sql
├── 009_create_scores.down.sql
├── 010_create_notifications.up.sql
├── 010_create_notifications.down.sql
├── 011_create_audit.up.sql
├── 011_create_audit.down.sql
└── 012_create_mentor.up.sql
└── 012_create_mentor.down.sql
```

Untuk SQLite dev: 
- Generate migration SQL yang kompatibel SQLite
- Atau gunakan driver SQLite `golang-migrate` dengan migration PG yang disesuaikan
