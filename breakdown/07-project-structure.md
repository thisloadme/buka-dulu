# Project Structure — BukaDulu MVP
## Versi 1.0 | Go Backend + Flutter Frontend Monorepo

---

## 1. Root Layout

```
buka-dulu/
├── backend/                   # Go API server
├── frontend/                  # Flutter app
├── docs/                      # Dokumentasi tambahan
├── docker-compose.yml         # Dev environment (opsional)
├── Makefile                   # Root commands
├── .gitignore
└── README.md
```

---

## 2. Backend (Go)

```
backend/
├── cmd/
│   └── server/
│       └── main.go                # Entry point: init config, DB, router, start server
│
├── internal/
│   ├── config/
│   │   └── config.go              # Env loading (DB_URL, JWT_SECRET, LLM_API_KEY, PORT)
│   │
│   ├── domain/
│   │   ├── user.go                # User, Role structs
│   │   ├── venture.go             # Venture, VentureStage enum
│   │   ├── idea.go                # Idea struct
│   │   ├── customer.go            # CustomerSegment struct
│   │   ├── menu.go                # Menu struct
│   │   ├── cost.go                # Ingredient, Packaging, CostSummary structs
│   │   ├── mission.go             # Mission struct + MissionType enum
│   │   ├── evidence.go            # Evidence struct + EvidenceType enum
│   │   ├── review.go              # EvidenceReview, Verdict structs
│   │   ├── score.go               # Score, Decision structs
│   │   ├── notification.go        # Notification struct
│   │   └── errors.go              # Domain-specific error types
│   │
│   ├── handler/
│   │   ├── router.go              # Chi router setup, mount all routes
│   │   ├── auth.go                # POST /auth/register, /login, /refresh, /logout
│   │   ├── venture.go             # CRUD ventures
│   │   ├── idea.go                # Capture, process, confirm idea
│   │   ├── customer.go            # Customer segment CRUD
│   │   ├── menu.go                # Menu CRUD, scoring, focus
│   │   ├── cost.go                # Ingredient CRUD, cost calculation, confirm
│   │   ├── mission.go             # Mission CRUD, generate, accept, complete
│   │   ├── evidence.go            # Evidence upload, list, detail
│   │   ├── review.go              # Trigger review, override
│   │   ├── score.go               # Score display, calculate, history
│   │   ├── decision.go            # Generate decision, get decision
│   │   ├── notification.go        # Notification list, mark read
│   │   ├── mentor.go              # Mentee list, comments
│   │   └── health.go              # GET /health
│   │
│   ├── middleware/
│   │   ├── auth.go                # JWT verification, user context injection
│   │   ├── cors.go                # CORS config
│   │   ├── logger.go              # Request logging (slog)
│   │   └── recovery.go            # Panic recovery
│   │
│   ├── service/
│   │   ├── auth.go                # Register logic, password hashing, JWT generation
│   │   ├── venture.go             # Venture business logic, stage transitions
│   │   ├── idea.go                # Idea structuring orchestration → LLM
│   │   ├── customer.go            # Segment validation, "too general" check
│   │   ├── menu.go                # Menu complexity scoring → LLM
│   │   ├── cost.go                # HPP calculation, margin analysis, break-even
│   │   ├── mission.go             # Mission generation → LLM, priority queue
│   │   ├── evidence.go            # Evidence validation, file handling
│   │   ├── review.go              # Evidence review → LLM, verdict logic
│   │   ├── scoring.go             # Score calculation, decision engine
│   │   ├── notification.go        # Notification creation, dispatch
│   │   ├── llm.go                 # AI client: OpenAI/Anthropic API calls
│   │   └── event.go               # Domain event dispatcher (internal)
│   │
│   ├── repository/
│   │   ├── user.go                # User CRUD queries
│   │   ├── venture.go             # Venture CRUD, stage queries
│   │   ├── idea.go                # Idea CRUD, versioning
│   │   ├── customer.go            # Customer segment queries
│   │   ├── menu.go                # Menu CRUD, complexity queries
│   │   ├── cost.go                # Ingredient, packaging, cost summary queries
│   │   ├── mission.go             # Mission CRUD, status queries
│   │   ├── evidence.go            # Evidence CRUD
│   │   ├── review.go              # Review CRUD
│   │   ├── score.go               # Score CRUD
│   │   ├── decision.go            # Decision CRUD
│   │   ├── notification.go        # Notification CRUD
│   │   └── audit.go               # Audit log write
│   │
│   └── engine/
│       ├── stage.go               # State machine: transition validation gate checks
│       ├── decision.go            # Decision matrix: score + economics → outcome
│       └── mission.go             # Mission template selector, priority sorter
│
├── migrations/
│   ├── 001_create_users.up.sql
│   ├── 001_create_users.down.sql
│   ├── 002_create_ventures.up.sql
│   ├── 002_create_ventures.down.sql
│   ├── 003_create_ideas.up.sql
│   ├── 003_create_ideas.down.sql
│   ├── 004_create_customer_segments.up.sql
│   ├── 004_create_customer_segments.down.sql
│   ├── 005_create_menus.up.sql
│   ├── 005_create_menus.down.sql
│   ├── 006_create_ingredients.up.sql
│   ├── 006_create_ingredients.down.sql
│   ├── 007_create_missions.up.sql
│   ├── 007_create_missions.down.sql
│   ├── 008_create_evidences.up.sql
│   ├── 008_create_evidences.down.sql
│   ├── 009_create_scores.up.sql
│   ├── 009_create_scores.down.sql
│   ├── 010_create_notifications.up.sql
│   ├── 010_create_notifications.down.sql
│   ├── 011_create_audit.up.sql
│   ├── 011_create_audit.down.sql
│   ├── 012_create_mentor.up.sql
│   └── 012_create_mentor.down.sql
│
├── storage/                       # Local file storage for dev (gitignored)
│   └── evidence/
│
├── go.mod
├── go.sum
├── .env.example                   # Template for env vars
├── .env                           # Local env (gitignored)
└── Makefile                       # build / run / migrate / test / lint
```

### File Role Guide

| File Pattern | Responsibility |
|---|---|
| `domain/*.go` | Pure structs, enums, constants. Zero dependencies. |
| `handler/*.go` | HTTP layer: parse request, call service, write response. Thin. |
| `service/*.go` | Business logic. Orchestrates domain + repository + LLM. |
| `repository/*.go` | Data access. SQL queries only. Returns domain models. |
| `engine/*.go` | Pure logic functions: state transitions, decision matrix, sorting. Test-heavy. |
| `middleware/*.go` | Request pipeline: auth, CORS, logging, recovery. |
| `config/*.go` | Env parsing into typed struct. |

---

## 3. Frontend (Flutter)

```
frontend/
├── lib/
│   ├── main.dart                      # Entry point: ProviderScope + MaterialApp.router
│   ├── app.dart                       # App widget, theme, configuration
│   │
│   ├── config/
│   │   ├── api_config.dart            # Base URL, timeout, headers
│   │   ├── theme.dart                 # ThemeData, colors, typography
│   │   └── constants.dart             # App-wide constants
│   │
│   ├── domain/
│   │   ├── models/                    # Data classes (dart from json / freezed)
│   │   │   ├── user.dart
│   │   │   ├── venture.dart
│   │   │   ├── idea.dart
│   │   │   ├── customer_segment.dart
│   │   │   ├── menu.dart
│   │   │   ├── ingredient.dart
│   │   │   ├── cost_summary.dart
│   │   │   ├── mission.dart
│   │   │   ├── evidence.dart
│   │   │   ├── review.dart
│   │   │   ├── score.dart
│   │   │   ├── decision.dart
│   │   │   ├── notification.dart
│   │   │   └── paginated_response.dart
│   │   │
│   │   └── enums/
│   │       ├── venture_stage.dart
│   │       ├── menu_status.dart
│   │       ├── mission_status.dart
│   │       ├── verdict.dart
│   │       └── decision_type.dart
│   │
│   ├── data/
│   │   ├── datasources/
│   │   │   ├── auth_api.dart          # HTTP calls for auth
│   │   │   ├── venture_api.dart
│   │   │   ├── idea_api.dart
│   │   │   ├── menu_api.dart
│   │   │   ├── cost_api.dart
│   │   │   ├── mission_api.dart
│   │   │   ├── evidence_api.dart
│   │   │   ├── score_api.dart
│   │   │   ├── mentor_api.dart
│   │   │   └── notification_api.dart
│   │   │
│   │   ├── repositories/
│   │   │   ├── auth_repository.dart       # Wraps datasource + local storage
│   │   │   ├── venture_repository.dart
│   │   │   ├── idea_repository.dart
│   │   │   ├── menu_repository.dart
│   │   │   ├── cost_repository.dart
│   │   │   ├── mission_repository.dart
│   │   │   ├── evidence_repository.dart
│   │   │   ├── score_repository.dart
│   │   │   └── notification_repository.dart
│   │   │
│   │   └── auth/
│   │       └── token_storage.dart         # Flutter secure storage for JWT
│   │
│   ├── presentation/
│   │   ├── providers/                    # Riverpod providers
│   │   │   ├── auth_provider.dart        # Auth state + actions
│   │   │   ├── venture_list_provider.dart
│   │   │   ├── venture_detail_provider.dart
│   │   │   ├── idea_provider.dart
│   │   │   ├── customer_provider.dart
│   │   │   ├── menu_provider.dart
│   │   │   ├── cost_provider.dart
│   │   │   ├── mission_provider.dart
│   │   │   ├── evidence_provider.dart
│   │   │   ├── score_provider.dart
│   │   │   └── notification_provider.dart
│   │   │
│   │   ├── pages/
│   │   │   ├── auth/
│   │   │   │   ├── login_page.dart
│   │   │   │   └── register_page.dart
│   │   │   ├── dashboard/
│   │   │   │   └── dashboard_page.dart
│   │   │   ├── venture/
│   │   │   │   ├── venture_create_page.dart
│   │   │   │   ├── venture_detail_page.dart
│   │   │   │   └── venture_stage_bar.dart
│   │   │   ├── idea/
│   │   │   │   ├── idea_capture_page.dart
│   │   │   │   └── idea_result_page.dart
│   │   │   ├── customer/
│   │   │   │   └── customer_page.dart
│   │   │   ├── menu/
│   │   │   │   └── menu_page.dart
│   │   │   ├── cost/
│   │   │   │   └── cost_page.dart
│   │   │   ├── mission/
│   │   │   │   ├── mission_board_page.dart
│   │   │   │   └── mission_detail_page.dart
│   │   │   ├── evidence/
│   │   │   │   ├── evidence_upload_page.dart
│   │   │   │   └── evidence_detail_page.dart
│   │   │   ├── score/
│   │   │   │   └── score_page.dart
│   │   │   ├── decision/
│   │   │   │   └── decision_page.dart
│   │   │   ├── courtroom/
│   │   │   │   └── courtroom_page.dart
│   │   │   └── mentor/
│   │   │       ├── mentee_list_page.dart
│   │   │       └── mentee_detail_page.dart
│   │   │
│   │   └── widgets/
│   │       ├── common/
│   │       │   ├── loading_skeleton.dart
│   │       │   ├── empty_state.dart
│   │       │   ├── error_state.dart
│   │       │   ├── confirm_bottom_sheet.dart
│   │       │   └── app_shell.dart          # Scaffold with bottom nav / sidebar
│   │       ├── venture_card.dart
│   │       ├── stage_badge.dart
│   │       ├── score_circle.dart
│   │       ├── mission_card.dart
│   │       ├── evidence_preview.dart
│   │       ├── cost_row.dart
│   │       └── progress_indicator_14day.dart
│   │
│   └── routing/
│       ├── router.dart                     # GoRouter config
│       └── route_names.dart                # Named route constants
│
├── assets/
│   ├── images/
│   │   ├── logo.svg
│   │   ├── empty-state.svg
│   │   └── illustrations/
│   └── fonts/                              # (opsional)
│
├── test/
│   ├── unit/
│   │   ├── providers/
│   │   └── services/
│   ├── widget/
│   │   ├── pages/
│   │   └── widgets/
│   └── integration/
│
├── web/                                    # Flutter web config
│   ├── index.html
│   └── manifest.json
│
├── pubspec.yaml
├── analysis_options.yaml
├── .env.example
└── Makefile
```

### Flutter Provider Pattern

```dart
// Typical Riverpod provider pattern:
// 1. Define state notifier
class VentureListNotifier extends StateNotifier<AsyncValue<List<Venture>>> {
  final VentureRepository _repo;
  
  VentureListNotifier(this._repo) : super(const AsyncValue.loading());
  
  Future<void> load() async {
    state = const AsyncValue.loading();
    state = await AsyncValue.guard(() => _repo.getAll());
  }
}

// 2. Expose via provider
final ventureListProvider = StateNotifierProvider<VentureListNotifier, AsyncValue<List<Venture>>>((ref) {
  return VentureListNotifier(ref.read(ventureRepositoryProvider));
});

// 3. Use in page
class DashboardPage extends ConsumerWidget {
  Widget build(BuildContext context, WidgetRef ref) {
    final venturesAsync = ref.watch(ventureListProvider);
    return venturesAsync.when(
      data: (ventures) => ListView.builder(...),
      loading: () => SkeletonLoader(),
      error: (e, _) => ErrorState(e),
    );
  }
}
```

---

## 4. Docker Compose (Dev Environment)

```yaml
version: '3.8'
services:
  backend:
    build: ./backend
    ports:
      - "8080:8080"
    volumes:
      - ./backend:/app
      - ./backend/storage:/app/storage
    env_file: ./backend/.env
    depends_on:
      - db
      - minio
    command: air  # hot reload

  db:
    image: postgres:16-alpine
    environment:
      POSTGRES_USER: bukadulu
      POSTGRES_PASSWORD: devpass
      POSTGRES_DB: bukadulu
    ports:
      - "5432:5432"
    volumes:
      - pgdata:/var/lib/postgresql/data

  minio:
    image: minio/minio
    ports:
      - "9000:9000"
      - "9001:9001"
    environment:
      MINIO_ROOT_USER: minioadmin
      MINIO_ROOT_PASSWORD: minioadmin
    command: server /data --console-address ":9001"
    volumes:
      - miniodata:/data

volumes:
  pgdata:
  miniodata:
```

---

## 5. Makefile (Root)

```makefile
.PHONY: dev-backend dev-frontend migrate-up migrate-down test help

dev-backend:
	cd backend && air

dev-frontend:
	cd frontend && flutter run -d chrome

migrate-up:
	cd backend && go run cmd/migrate/main.go up

migrate-down:
	cd backend && go run cmd/migrate/main.go down

test-backend:
	cd backend && go test ./... -v

test-frontend:
	cd frontend && flutter test

help:
	@grep '^[a-zA-Z0-9_-]*:' Makefile | sed 's/:.*//' | sort

build:
	cd backend && go build -o bin/server cmd/server/main.go
```

---

## 6. Environment Variables

```bash
# Backend (.env)
PORT=8080
DATABASE_URL=sqlite://./data/bukadulu.db  # dev: sqlite, prod: postgres://...
JWT_SECRET=your-secret-key-change-in-prod
JWT_EXPIRY_HOURS=24

LLM_PROVIDER=openai                        # openai | anthropic
LLM_API_KEY=sk-...
LLM_MODEL=gpt-4o-mini                      # cheaper model for MVP

STORAGE_TYPE=local                         # local | s3
STORAGE_PATH=./storage/evidence
# S3_BUCKET=...
# S3_REGION=...
# S3_ACCESS_KEY=...
# S3_SECRET_KEY=...

CORS_ORIGINS=http://localhost:5173,http://localhost:3000,http://localhost:8080

# Frontend (.env)
API_BASE_URL=http://localhost:8080/api/v1
ENVIRONMENT=development
```

---

## 7. Key Design Decisions

| Keputusan | Alasan |
|---|---|
| **Manual SQL (sqlx), bukan ORM** | Kontrol penuh atas query, performa lebih baik, migration jelas |
| **UUID v4 generate di Go** | Kompatibel SQLite & PostgreSQL, tidak tergantung DB |
| **State machine di `engine/stage.go`** | Business rules terpusat, testable tanpa dependency |
| **Slog + structured logging** | Zero external dependency, cukup untuk MVP |
| **Riverpod** | Type-safe, compiles, better than BLoC for product apps |
| **GoRouter** | Standard Flutter routing, web URL support |
| **Air (hot reload) untuk Go** | Live reload tanpa compile manual |
| **sqlite untuk dev** | Zero setup, file-based, same sqlx code works with PG |
