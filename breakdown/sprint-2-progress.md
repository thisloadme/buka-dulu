# Sprint 2 Progress — BukaDulu MVP
## Status: ✅ COMPLETED
## Tanggal: 2026-06-08

---

## Ringkasan

Sprint 2 menyelesaikan **Customer Segment, Menu Focus Engine, dan Cost & Margin Engine** (US-030, US-040, US-041, US-042, US-050, US-051).

---

## Backend — Endpoint Baru

| Endpoint | Method | Status |
|---|---|---|
| `/api/v1/ventures/{id}/customer` | POST | ✅ Create customer segment |
| `/api/v1/ventures/{id}/customer` | GET | ✅ Get customer segment |
| `/api/v1/ventures/{id}/customer/confirm` | POST | ✅ Lock & transition stage |
| `/api/v1/ventures/{id}/menus` | POST | ✅ Add menu candidate |
| `/api/v1/ventures/{id}/menus` | GET | ✅ List menus |
| `/api/v1/ventures/{id}/menus/{menuId}` | PUT | ✅ Update status (active/max 3) |
| `/api/v1/ventures/{id}/menus/{menuId}` | DELETE | ✅ Delete menu |
| `/api/v1/ventures/{id}/menus/focus` | POST | ✅ Lock SKU selection |
| `/api/v1/ventures/{id}/ingredients` | POST | ✅ Add ingredient |
| `/api/v1/ventures/{id}/ingredients` | GET | ✅ List ingredients by menu |
| `/api/v1/ventures/{id}/ingredients/{id}` | DELETE | ✅ Delete ingredient |
| `/api/v1/ventures/{id}/packaging` | POST | ✅ Add packaging cost |
| `/api/v1/ventures/{id}/packaging` | GET | ✅ List packaging costs |
| `/api/v1/ventures/{id}/cost/calculate/{menuId}` | POST | ✅ HPP + margin + break-even |
| `/api/v1/ventures/{id}/cost/summary/{menuId}` | GET | ✅ Get cost summary |
| `/api/v1/ventures/{id}/cost/summaries` | GET | ✅ All cost summaries |
| `/api/v1/ventures/{id}/cost/confirm` | POST | ✅ Lock cost & transition |

### Bugs Fixed
- `UNIQUE constraint failed: users.phone` → Pass nil instead of empty string to DB
- `converting NULL to string is unsupported` → Use `sql.NullString` + `scanUser()` helper
- Login failing → Both bugs above caused scan errors

### Stage Transitions Verified
```
draft → idea_defined → customer_defined → sku_focused → cost_evaluated ✅
```

---

## Frontend — Pages Baru

| Page | File | Status |
|---|---|---|
| Customer Segment | `pages/customer/customer_page.dart` | ✅ 6 field form + confirm |
| Menu Focus | `pages/menu/menu_page.dart` | ✅ Add/activate/delete + max 3 + focus |
| Cost & Margin | `pages/cost/cost_page.dart` | ✅ Ingredients input + packaging + calculate + margin status |

`flutter analyze`: 0 errors, 1 warning (deprecated `value` → `initialValue`), 4 infos

---

## Project Files Added/Modified

### Backend (new)
- `internal/domain/customer.go` — CustomerSegment struct
- `internal/domain/menu.go` — Menu struct  
- `internal/domain/cost.go` — Ingredient, PackagingCost, CostSummary structs
- `internal/repository/customer.go` — CRUD
- `internal/repository/menu.go` — CRUD + CountActive
- `internal/repository/ingredient.go` — CRUD + CostSummary + Packaging
- `internal/service/customer.go` — Business logic + "too general" detection
- `internal/service/menu.go` — Max 3 SKU enforcement + Focus
- `internal/service/cost.go` — HPP calc, margin analysis, break-even
- `internal/handler/customer.go` — HTTP handlers
- `internal/handler/menu.go` — HTTP handlers
- `internal/handler/cost.go` — HTTP handlers

### Backend (modified)
- `internal/handler/router.go` — All Sprint 2 routes registered
- `internal/repository/user.go` — `scanUser()` helper, nil handling for phone/email
- `cmd/server/main.go` — All Sprint 2 services wired

### Frontend (new)
- `pages/customer/customer_page.dart`
- `pages/menu/menu_page.dart`
- `pages/cost/cost_page.dart`
- `routing/router.dart` — New Sprint 2 routes

### Frontend (modified)
- `data/datasources/api.dart` — All Sprint 2 API methods

---

## Cara Test Full Flow (Web)

```bash
# Backend (terminal 1)
cd backend && ./server

# Frontend (terminal 2)
cd frontend && flutter run -d chrome
```

Flow: Register → Login → Dashboard → Ide Baru → 
  Capture Ide → Proses → Confirm → 
  Customer Segmen → Confirm → 
  Menu (add 1-3, activate) → Focus → 
  Cost (add ingredients, calculate HPP) → Confirm

---

## Catatan untuk Sprint 3

Misi harian (US-060-062) + Evidence upload + Review engine.
