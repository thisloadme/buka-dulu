# Sprint 1 Progress — BukaDulu MVP
## Status: ✅ COMPLETED
## Tanggal: 2026-06-08
## Durasi: 1 sesi

---

## Ringkasan

Sprint 1 berhasil menyelesaikan **backend + frontend core** untuk US-001, US-002, US-010, US-011, US-020, US-021, US-022.

---

## Backend (Go)

### Struktur
```
backend/
├── cmd/server/main.go              # Entry point
├── internal/
│   ├── config/config.go            # Env loading (PORT, DB_URL, JWT_SECRET, LLM_API_KEY)
│   ├── config/db.go                # SQLite init + migration runner
│   ├── domain/
│   │   ├── user.go                 # User, RegisterRequest, LoginRequest, AuthResponse
│   │   ├── venture.go              # Venture, VentureStage enum (13 stages)
│   │   ├── idea.go                 # Idea, StructuredConcept, UpdateIdeaRequest
│   │   └── errors.go               # Domain errors + AppError codes
│   ├── engine/stage.go             # State machine: allowed transitions, gate check, stage order
│   ├── repository/
│   │   ├── user.go                 # Create, FindByEmail, FindByPhone, FindByID, UpdateLastLogin
│   │   ├── venture.go              # CRUD + UpdateStage
│   │   └── idea.go                 # Create, FindByVenture, Update
│   ├── service/
│   │   ├── auth.go                 # Register (bcrypt + JWT), Login, ValidateToken
│   │   ├── venture.go              # Create, List, Get, Update, TransitionStage
│   │   ├── idea.go                 # Capture, Process (AI), Update, Confirm
│   │   └── llm.go                  # OpenAI client + mock mode (dev fallback)
│   ├── handler/
│   │   ├── auth.go                 # POST /auth/register, /auth/login
│   │   ├── venture.go              # CRUD ventures
│   │   ├── idea.go                 # Capture, Process, Get, Update, Confirm
│   │   ├── helpers.go              # writeJSON utility
│   │   └── router.go               # Chi router setup, middleware, route mounting
│   └── middleware/
│       ├── auth.go                 # JWT verification + CORS + Logger
│       └── recovery.go             # Panic recovery
├── migrations/001_init.up.sql      # Full schema (15+ tables)
├── server                          # Compiled binary
├── .env                            # Configuration
└── data/bukadulu.db                # SQLite database (gitignored)
```

### Test Result (end-to-end)

| Flow | Status |
|---|---|
| `POST /auth/register` | ✅ |
| `POST /auth/login` | ✅ (JWT token returned) |
| `POST /ventures` | ✅ |
| `GET /ventures` | ✅ |
| `GET /ventures/:id` | ✅ |
| `POST /ventures/:id/idea` (capture) | ✅ |
| `GET /ventures/:id/idea` | ✅ |
| `POST /ventures/:id/idea/process` | ✅ (mock AI: return structured concept) |
| `POST /ventures/:id/idea/confirm` | ✅ (stage transition: draft → idea_defined) |
| `GET /health` | ✅ |

### Key Technical Decisions

| Issue | Solution |
|---|---|
| SQLite datetime → Go time.Time scan error | Changed domain timestamps to `string` (RFC3339), set in Go code, not DB |
| NULL last_login_at scan error | Used `sql.NullString` in repository, mapped to `*string` in domain |
| No LLM API key for dev | Added mock mode — when `LLM_API_KEY` empty, returns realistic dummy concept |
| Token parsing in frontend curl test | Used Python `requests` library instead of shell curl |

### Dependencies

```
chi/v5, sqlx (not used directly — using database/sql), go-sqlite3, 
golang-jwt/v5, google/uuid, golang.org/x/crypto/bcrypt
```

---

## Frontend (Flutter)

### Struktur
```
frontend/
├── lib/
│   ├── main.dart                   # Entry: dotenv + ProviderScope
│   ├── app.dart                    # MaterialApp.router
│   ├── config/
│   │   ├── api_config.dart         # Dio provider (base URL from .env)
│   │   └── theme.dart              # Material 3 theme, Inter font
│   ├── domain/models/
│   │   ├── user.dart               # User.fromJson
│   │   ├── auth_response.dart      # AuthResponse.fromJson
│   │   ├── venture.dart            # Venture.fromJson
│   │   └── idea.dart               # Idea.fromJson
│   ├── data/datasources/api.dart   # AuthApi: register, login, venture, idea CRUD
│   ├── presentation/
│   │   ├── providers/
│   │   │   ├── token_provider.dart   # StateProvider<String?>
│   │   │   ├── auth_provider.dart    # AuthNotifier (StateNotifier)
│   │   │   ├── venture_provider.dart # ventureListProvider, ventureDetailProvider
│   │   │   └── idea_provider.dart    # ideaProvider (FutureProvider.family)
│   │   └── pages/
│   │       ├── auth/
│   │       │   ├── login_page.dart       # Login form with validation
│   │       │   └── register_page.dart    # Register form with validation
│   │       ├── dashboard/
│   │       │   └── dashboard_page.dart   # Venture list + empty state + stage badges
│   │       ├── venture/
│   │       │   └── venture_create_page.dart  # Create venture form
│   │       └── idea/
│   │           ├── idea_capture_page.dart    # Textarea for raw idea
│   │           └── idea_result_page.dart     # AI concept display + confirm
│   └── routing/router.dart        # GoRouter: 6 routes
├── .env                           # API_BASE_URL=http://localhost:8080/api/v1
└── pubspec.yaml                   # Dependencies: riverpod, go_router, dio, etc.
```

### Analysis Result
- `flutter analyze`: 0 errors, 0 warnings, 4 info (unnecessary underscores — style)
- Tested with `flutter build` → passes

---

## Cara Running

### Backend
```bash
cd backend
# Set LLM_API_KEY di .env untuk AI real, atau kosongkan untuk mock
./server
# → http://localhost:8080
```

### Frontend (web)
```bash
cd frontend
flutter run -d chrome
# → http://localhost:3000 (flutter default)
```

---

## Yang Belum / Sprint 2

| Item | Sprint |
|---|---|
| Customer segment page (US-030) | Sprint 2 |
| Menu focus engine (US-040-042) | Sprint 2 |
| Cost & margin engine (US-050-051) | Sprint 2 |
| AI integration real (OpenAI/Anthropic) | Sprint 2+ |
| Mission board (US-060-062) | Sprint 3 |
| Evidence upload & review (US-070-080) | Sprint 3 |
| Scoring & decision (US-090-091) | Sprint 4 |
| Mentor dashboard (US-120-121) | Sprint 4 |
| Flutter web deployment | Sprint 4+ |
| PostgreSQL production | Setelah MVP validated |
