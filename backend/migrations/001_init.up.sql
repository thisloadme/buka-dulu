-- Auth
CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY,
    role TEXT NOT NULL DEFAULT 'founder' CHECK(role IN ('founder','mentor','admin')),
    full_name TEXT NOT NULL,
    email TEXT UNIQUE,
    phone TEXT UNIQUE,
    password_hash TEXT NOT NULL,
    status TEXT NOT NULL DEFAULT 'active' CHECK(status IN ('active','suspended','deleted')),
    last_login_at TEXT,
    created_at TEXT NOT NULL DEFAULT (CURRENT_TIMESTAMP),
    updated_at TEXT NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

-- Ventures
CREATE TABLE IF NOT EXISTS ventures (
    id TEXT PRIMARY KEY,
    owner_user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    category TEXT,
    region TEXT,
    stage TEXT NOT NULL DEFAULT 'draft' CHECK(stage IN (
        'draft','idea_defined','customer_defined','sku_focused','cost_evaluated',
        'mission_active','evidence_submitted','evidence_reviewed','ready_to_decide',
        'continue','repeat','pivot','stop'
    )),
    current_version INTEGER NOT NULL DEFAULT 1,
    created_at TEXT NOT NULL DEFAULT (CURRENT_TIMESTAMP),
    updated_at TEXT NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);
CREATE INDEX IF NOT EXISTS idx_ventures_owner ON ventures(owner_user_id);
CREATE INDEX IF NOT EXISTS idx_ventures_stage ON ventures(stage);

-- Venture versions
CREATE TABLE IF NOT EXISTS venture_versions (
    id TEXT PRIMARY KEY,
    venture_id TEXT NOT NULL REFERENCES ventures(id) ON DELETE CASCADE,
    version INTEGER NOT NULL,
    snapshot TEXT NOT NULL, -- JSON
    created_at TEXT NOT NULL DEFAULT (CURRENT_TIMESTAMP),
    UNIQUE(venture_id, version)
);

-- Ideas
CREATE TABLE IF NOT EXISTS ideas (
    id TEXT PRIMARY KEY,
    venture_id TEXT NOT NULL REFERENCES ventures(id) ON DELETE CASCADE,
    raw_input TEXT NOT NULL,
    one_line_concept TEXT,
    target_customer TEXT,
    value_proposition TEXT,
    key_assumptions TEXT,
    early_risks TEXT,
    version INTEGER NOT NULL DEFAULT 1,
    is_locked BOOLEAN NOT NULL DEFAULT FALSE,
    status TEXT NOT NULL DEFAULT 'pending' CHECK(status IN ('pending','processing','done','failed')),
    ai_raw_input TEXT,
    ai_raw_output TEXT,
    created_at TEXT NOT NULL DEFAULT (CURRENT_TIMESTAMP),
    updated_at TEXT NOT NULL DEFAULT (CURRENT_TIMESTAMP),
    UNIQUE(venture_id, version)
);

-- Customer segments
CREATE TABLE IF NOT EXISTS customer_segments (
    id TEXT PRIMARY KEY,
    venture_id TEXT NOT NULL REFERENCES ventures(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    age_range TEXT,
    buy_context TEXT,
    budget_range TEXT,
    location TEXT,
    consumption_moment TEXT,
    is_too_general BOOLEAN NOT NULL DEFAULT FALSE,
    is_locked BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TEXT NOT NULL DEFAULT (CURRENT_TIMESTAMP),
    updated_at TEXT NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

-- Menu focus
CREATE TABLE IF NOT EXISTS menus (
    id TEXT PRIMARY KEY,
    venture_id TEXT NOT NULL REFERENCES ventures(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    description TEXT,
    status TEXT NOT NULL DEFAULT 'candidate' CHECK(status IN ('candidate','active','deferred','dropped')),
    is_hero BOOLEAN NOT NULL DEFAULT FALSE,
    complexity_score REAL,
    complexity_factors TEXT, -- JSON
    sort_order INTEGER NOT NULL DEFAULT 0,
    created_at TEXT NOT NULL DEFAULT (CURRENT_TIMESTAMP),
    updated_at TEXT NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);
CREATE INDEX IF NOT EXISTS idx_menus_venture ON menus(venture_id);

-- Ingredients
CREATE TABLE IF NOT EXISTS ingredients (
    id TEXT PRIMARY KEY,
    venture_id TEXT NOT NULL REFERENCES ventures(id) ON DELETE CASCADE,
    menu_id TEXT NOT NULL REFERENCES menus(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    unit TEXT NOT NULL,
    quantity REAL NOT NULL,
    unit_price REAL NOT NULL,
    created_at TEXT NOT NULL DEFAULT (CURRENT_TIMESTAMP),
    updated_at TEXT NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);
CREATE INDEX IF NOT EXISTS idx_ingredients_menu ON ingredients(menu_id);

-- Packaging costs
CREATE TABLE IF NOT EXISTS packaging_costs (
    id TEXT PRIMARY KEY,
    venture_id TEXT NOT NULL REFERENCES ventures(id) ON DELETE CASCADE,
    menu_id TEXT NOT NULL REFERENCES menus(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    unit_price REAL NOT NULL,
    quantity_per_pcs REAL NOT NULL DEFAULT 1,
    created_at TEXT NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

-- Cost summaries
CREATE TABLE IF NOT EXISTS cost_summaries (
    id TEXT PRIMARY KEY,
    venture_id TEXT NOT NULL REFERENCES ventures(id) ON DELETE CASCADE,
    menu_id TEXT NOT NULL REFERENCES menus(id) ON DELETE CASCADE,
    hpp_per_porsi REAL NOT NULL,
    suggested_price REAL NOT NULL,
    target_margin REAL NOT NULL,
    gross_margin REAL NOT NULL,
    margin_status TEXT NOT NULL CHECK(margin_status IN ('sehat','tipis','berbahaya')),
    break_even_unit INTEGER NOT NULL,
    labor_per_unit REAL,
    overhead_per_unit REAL,
    is_locked BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TEXT NOT NULL DEFAULT (CURRENT_TIMESTAMP),
    updated_at TEXT NOT NULL DEFAULT (CURRENT_TIMESTAMP),
    UNIQUE(venture_id, menu_id)
);

-- Missions
CREATE TABLE IF NOT EXISTS missions (
    id TEXT PRIMARY KEY,
    venture_id TEXT NOT NULL REFERENCES ventures(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    mission_type TEXT NOT NULL CHECK(mission_type IN ('polling','pre_order','sampling','titip_jual','interview','harga_test','observation')),
    priority TEXT NOT NULL DEFAULT 'medium' CHECK(priority IN ('high','medium','low')),
    status TEXT NOT NULL DEFAULT 'pending' CHECK(status IN ('pending','accepted','in_progress','completed','skipped')),
    due_at TEXT,
    evidence_required TEXT,
    estimated_minutes INTEGER,
    created_by TEXT REFERENCES users(id),
    sort_order INTEGER NOT NULL DEFAULT 0,
    created_at TEXT NOT NULL DEFAULT (CURRENT_TIMESTAMP),
    updated_at TEXT NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);
CREATE INDEX IF NOT EXISTS idx_missions_venture ON missions(venture_id);
CREATE INDEX IF NOT EXISTS idx_missions_status ON missions(status);

-- Evidence
CREATE TABLE IF NOT EXISTS evidences (
    id TEXT PRIMARY KEY,
    venture_id TEXT NOT NULL REFERENCES ventures(id) ON DELETE CASCADE,
    mission_id TEXT NOT NULL REFERENCES missions(id) ON DELETE CASCADE,
    uploader_id TEXT NOT NULL REFERENCES users(id),
    evidence_type TEXT NOT NULL CHECK(evidence_type IN ('image','text','link')),
    storage_url TEXT,
    text_content TEXT,
    thumbnail_url TEXT,
    file_size_bytes INTEGER,
    mime_type TEXT,
    submitted_at TEXT NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);
CREATE INDEX IF NOT EXISTS idx_evidences_mission ON evidences(mission_id);
CREATE INDEX IF NOT EXISTS idx_evidences_venture ON evidences(venture_id);

-- Evidence reviews
CREATE TABLE IF NOT EXISTS evidence_reviews (
    id TEXT PRIMARY KEY,
    evidence_id TEXT NOT NULL REFERENCES evidences(id) ON DELETE CASCADE,
    reviewer_type TEXT NOT NULL CHECK(reviewer_type IN ('ai','human_override')),
    verdict TEXT NOT NULL DEFAULT 'pending' CHECK(verdict IN ('valid','weak','invalid','suspicious','pending')),
    score REAL,
    rationale TEXT NOT NULL,
    next_action TEXT,
    overridden_by TEXT REFERENCES users(id),
    overridden_at TEXT,
    processing_time_ms INTEGER,
    ai_raw_output TEXT,
    created_at TEXT NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);
CREATE INDEX IF NOT EXISTS idx_reviews_evidence ON evidence_reviews(evidence_id);

-- Scores
CREATE TABLE IF NOT EXISTS scores (
    id TEXT PRIMARY KEY,
    venture_id TEXT NOT NULL REFERENCES ventures(id) ON DELETE CASCADE,
    clarity_score REAL NOT NULL DEFAULT 0,
    focus_score REAL NOT NULL DEFAULT 0,
    economics_score REAL NOT NULL DEFAULT 0,
    execution_score REAL NOT NULL DEFAULT 0,
    evidence_score REAL NOT NULL DEFAULT 0,
    market_response_score REAL NOT NULL DEFAULT 0,
    total_score REAL NOT NULL DEFAULT 0,
    is_final BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TEXT NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);
CREATE INDEX IF NOT EXISTS idx_scores_venture ON scores(venture_id);

-- Decisions
CREATE TABLE IF NOT EXISTS decisions (
    id TEXT PRIMARY KEY,
    venture_id TEXT NOT NULL REFERENCES ventures(id) ON DELETE CASCADE,
    decision TEXT NOT NULL CHECK(decision IN ('continue','repeat','pivot','stop')),
    rationale TEXT NOT NULL,
    score_snapshot TEXT NOT NULL, -- JSON
    triggered_by TEXT NOT NULL DEFAULT 'system',
    created_at TEXT NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);
CREATE INDEX IF NOT EXISTS idx_decisions_venture ON decisions(venture_id);

-- Notifications
CREATE TABLE IF NOT EXISTS notifications (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    venture_id TEXT REFERENCES ventures(id) ON DELETE CASCADE,
    type TEXT NOT NULL,
    title TEXT NOT NULL,
    body TEXT,
    is_read BOOLEAN NOT NULL DEFAULT FALSE,
    metadata TEXT, -- JSON
    created_at TEXT NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);
CREATE INDEX IF NOT EXISTS idx_notifications_user ON notifications(user_id);

-- Audit logs
CREATE TABLE IF NOT EXISTS audit_logs (
    id TEXT PRIMARY KEY,
    venture_id TEXT REFERENCES ventures(id),
    user_id TEXT REFERENCES users(id),
    action TEXT NOT NULL,
    entity_type TEXT,
    entity_id TEXT,
    old_value TEXT, -- JSON
    new_value TEXT, -- JSON
    ip_address TEXT,
    created_at TEXT NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);
CREATE INDEX IF NOT EXISTS idx_audit_venture ON audit_logs(venture_id);
CREATE INDEX IF NOT EXISTS idx_audit_created ON audit_logs(created_at);

-- Mentor
CREATE TABLE IF NOT EXISTS mentor_assignments (
    id TEXT PRIMARY KEY,
    mentor_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    founder_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    venture_id TEXT NOT NULL REFERENCES ventures(id) ON DELETE CASCADE,
    status TEXT NOT NULL DEFAULT 'active' CHECK(status IN ('active','archived')),
    created_at TEXT NOT NULL DEFAULT (CURRENT_TIMESTAMP),
    UNIQUE(mentor_id, venture_id)
);

CREATE TABLE IF NOT EXISTS mentor_comments (
    id TEXT PRIMARY KEY,
    mentor_id TEXT NOT NULL REFERENCES users(id),
    venture_id TEXT NOT NULL REFERENCES ventures(id) ON DELETE CASCADE,
    mission_id TEXT REFERENCES missions(id),
    evidence_id TEXT REFERENCES evidences(id),
    content TEXT NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TEXT NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);
