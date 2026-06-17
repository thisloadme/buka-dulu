# Arsitektur Sistem — BukaDulu MVP
## Versi 1.0 | Modular Monolith × Go Backend × Flutter Frontend

---

## 1. Prinsip Arsitektur

| Prinsip | Penjelasan |
|---|---|
| **Modular Monolith** | Satu service monolithic dengan domain boundary jelas. Microservices hanya setelah PMF. |
| **Domain-Driven** | Setiap domain punya own handler, service, repository, model. |
| **API-First** | Frontend komunikasi lewat REST API. Contract ditentukan sebelum implementasi. |
| **Stateless Backend** | Semua state di database. Backend bisa di-scale horizontal. |
| **Async Review** | Evidence review jalan async — user tidak diblok. |

---

## 2. High-Level Architecture

```
┌─────────────────────────────────────────────────┐
│                  Flutter Client                  │
│              (Web + Mobile)                      │
│         GoRouter + Riverpod + Repository         │
└──────────────────┬──────────────────────────────┘
                   │ HTTPS / JSON
                   ▼
┌─────────────────────────────────────────────────┐
│           Go API Server (chi router)             │
│                                                   │
│  ┌────────────┐  ┌────────────┐  ┌────────────┐ │
│  │ Middleware  │  │  Handler   │  │   Service  │ │
│  │ (Auth/CORS) │──▶    Layer   │──▶   Layer    │ │
│  └────────────┘  └────────────┘  └──────┬─────┘ │
│                                         │       │
│  ┌────────────┐  ┌────────────┐  ┌──────▼─────┐ │
│  │  External  │  │    LLM     │  │ Repository  │ │
│  │  Services  │◀─▶  Service   │  │   Layer     │ │
│  └────────────┘  └────────────┘  └──────┬─────┘ │
│                                          │       │
└──────────────────────────────────────────┼───────┘
                                           │
                    ┌──────────────────────┼──────┐
                    │                      │      │
                    ▼                      ▼      │
           ┌────────────────┐    ┌────────────┐   │
           │   SQLite/PG    │    │  Object    │   │
           │   (Database)   │    │  Storage   │   │
           └────────────────┘    │ (evidence) │   │
                                 └────────────┘   │
```

---

## 3. Domain Modules

Setiap domain adalah **paket independen** dalam `internal/`:

| Domain | Tanggung Jawab | Handler | Service | Repository |
|---|---|---|---|---|
| **auth** | Register, login, JWT, session | ✅ | ✅ | ✅ |
| **venture** | CRUD workspace, stage tracking | ✅ | ✅ | ✅ |
| **idea** | Capture, AI structuring, versioning | ✅ | ✅ | ✅ |
| **customer** | Segment definition, validation | ✅ | ✅ | ✅ |
| **menu** | SKU candidates, complexity scoring | ✅ | ✅ | ✅ |
| **cost** | HPP calculation, margin analysis | ✅ | ✅ | ✅ |
| **mission** | Generate, assign, track missions | ✅ | ✅ | ✅ |
| **evidence** | Upload, storage, metadata | ✅ | ✅ | ✅ |
| **review** | AI evidence review, scoring | ✅ | ✅ | ✅ |
| **scoring** | Readiness score, final decision | ✅ | ✅ | ✅ |
| **notification** | In-app notif, reminders | — | ✅ | ✅ |
| **mentor** | Mentee list, progress, comments | ✅ | ✅ | ✅ |
| **llm** | AI client wrapper (OpenAI/Anthropic) | — | ✅ | — |

### Dependency Graph (service layer)

```
handler → service → repository
              ↕
            llm service → external AI
```

Service layer boleh panggil service lain:
- `idea.service` → `llm.service`
- `menu.service` → `llm.service` (for scoring)
- `review.service` → `llm.service` (for evidence review)
- `mission.service` → `llm.service` (for mission generation)
- `scoring.service` → `venture.service`, `evidence.service`, `menu.service`

---

## 4. Tech Stack Detail

### Backend (Go)

| Lapisan | Pilihan | Alasan |
|---|---|---|
| **Router** | `chi` (go-chi/chi/v5) | Ringan, stdlib-compatible, middleware composable |
| **DB** | `sqlx` (jmoiron/sqlx) | Type-safe, raw SQL tanpa ORM overhead |
| **Migration** | `golang-migrate` | URL-based, support SQLite & PG |
| **Auth** | JWT via `golang-jwt/jwt/v5` | Stateless session, gampang di-Flutter |
| **Validation** | `go-playground/validator` | Struct tags validation |
| **AI Client** | Standard `net/http` + JSON | Langsung panggil OpenAI/Anthropic API |
| **Config** | `envconfig` + `.env` | DB URL, API keys, port |
| **Logging** | `slog` (stdlib) | Structured logging, zero dependency |

### Frontend (Flutter)

| Lapisan | Pilihan | Alasan |
|---|---|---|
| **State** | `Riverpod` (+ `flutter_riverpod`) | Type-safe, testable, no boilerplate |
| **Router** | `go_router` | Declarative routing, deep linking |
| **HTTP** | `dio` + `retrofit` (optional) | Interceptor, retry, multipart upload |
| **Storage** | `flutter_secure_storage` | JWT token storage |
| **Image** | `image_picker` | Upload evidence |
| **Env** | `flutter_dotenv` | API base URL |

### Database

| Layer | Dev | Production |
|---|---|---|
| **Database** | SQLite (file-based) | PostgreSQL 16 |
| **Media** | Local filesystem | S3-compatible (MinIO → AWS S3) |

---

## 5. AI Service Architecture

```
┌──────────────┐     ┌─────────────────┐     ┌───────────────┐
│  Service     │────▶│  LLM Service    │────▶│  OpenAI /     │
│  (request)   │     │  (internal)     │     │  Anthropic    │
└──────────────┘     │                 │     └───────────────┘
                     │  - Prompt mgmt  │
                     │  - Retry logic  │
                     │  - Schema       │
                     │    validation   │
                     │  - Raw logging  │
                     └────────┬────────┘
                              │
                     ┌────────▼────────┐
                     │  Response cache │
                     │  (optional)     │
                     └─────────────────┘
```

**AI digunakan untuk:**
1. Idea structuring (raw → structured concept)
2. Menu complexity scoring
3. Evidence review & verdict
4. Mission generation
5. Founder courtroom (3 perspective adversarial review)

**Safety:**
- Output AI selalu dinormalisasi ke schema terstruktur
- Raw input & output disimpan untuk debugging
- System prompt domain-specific per task
- Tidak boleh memberikan saran ilegal atau jaminan kesuksesan

---

## 6. Data Flow: Sprint 14 Hari

```
Day 1-2:                               Day 3-5:
┌──────────┐   ┌──────────┐          ┌──────────┐
│ Register │──▶│ Capture  │          │  Cost    │
│ Login    │   │ Idea     │──▶ AI ──▶│  Engine  │
│ Venture  │   │ Struct   │          │  Margin  │
└──────────┘   └──────────┘          └──────────┘
                    │                      │
                    ▼                      ▼
             ┌──────────────┐       ┌──────────────┐
             │ Customer Seg │       │ Menu Focus   │
             │ Confirmation │       │ Hero SKU     │
             └──────────────┘       └──────────────┘

Day 6-13:                              Day 14:
┌──────────┐   ┌──────────┐          ┌──────────┐
│ Missions │──▶│ Evidence │──▶ AI ──▶│  Score   │
│ Board    │   │ Upload   │   Review │  &       │──▶ Decision
│ Daily    │   │          │          │  Final   │
└──────────┘   └──────────┘          │  Gate    │
                                     └──────────┘
```

---

## 7. Error Handling Strategy

| Layer | Approach |
|---|---|
| **Handler** | Catch service errors → return structured JSON `{"error": "message", "code": "ERR_XXX"}` |
| **Service** | Return `(result, error)` — tidak pernah panic |
| **Repository** | Wrap DB errors ke domain errors (not found, conflict, internal) |
| **LLM** | Retry 3x exponential backoff → fallback ke "review gagal, coba lagi" |
| **Upload** | Jika upload media gagal, metadata tetap disimpan sebagai draft |

---

## 8. Security

| Area | Implementation |
|---|---|
| **Auth** | JWT with HMAC-SHA256, 24h expiry |
| **Password** | bcrypt (cost 12) |
| **API** | HTTPS only, CORS restricted ke origin frontend |
| **Media** | Signed URL untuk akses evidence (pre-signed, 1h expiry) |
| **Authorization** | Per-role & per-resource: founder hanya bisa akses venture sendiri |
| **Audit** | Setiap perubahan status, override, dan akses sensitif tercatat |

---

## 9. Observability (MVP)

| Area | Tool |
|---|---|
| **Logging** | `slog` → stdout → journald / CloudWatch |
| **Metrics** | Prometheus via `promhttp` (optional di MVP) |
| **Tracing** | Skip di MVP — cukup structured logging |
| **Health** | `GET /health` — DB ping, LLM status |
