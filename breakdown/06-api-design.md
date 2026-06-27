# API Design — BukaDulu MVP
## Versi 1.0 | REST × JSON × Token Auth

---

## 1. Convention

| Aturan | Detail |
|---|---|
| **Base URL** | Dev: `http://localhost:8080/api/v1` / Prod: `https://api.bukadulu.com/api/v1` |
| **Format** | JSON (Content-Type: `application/json`) |
| **Auth** | Bearer token di header `Authorization: Bearer <jwt>` |
| **Pagination** | `?page=1&limit=20` → Response: `{ data: [...], meta: { page, limit, total, total_pages } }` |
| **Errors** | `{ "error": { "code": "ERR_001", "message": "Deskripsi error", "details": {} } }` |
| **Dates** | ISO 8601 (`2026-06-08T10:30:00Z`) |
| **IDs** | UUID v4 |

---

## 2. Auth Endpoints

### POST /auth/register
```
Request:
{
    "full_name": "Budi Santoso",
    "email": "budi@email.com",        // required (OTP verification)
    "password": "min8char",
    "role": "founder"                  // default: founder
}

Response 201:
{
    "user": { "id", "full_name", "email", "role", "status": "pending", "created_at" },
    "token": "",                       // empty — user must verify OTP first
    "expires_at": "2026-06-09T10:30:00Z"  // OTP expiry
}

Errors:
- 409: Email already registered
- 422: Validation failed
```

### POST /auth/verify-otp
```
Request:
{
    "email": "budi@email.com",
    "otp": "123456"
}

Response 200:
{
    "user": { "id", "full_name", "email", "role", "status": "active" },
    "token": "eyJhbGci...",
    "expires_at": "2026-06-09T10:30:00Z"
}

Errors:
- 401: Invalid OTP
- 401: OTP expired
- 409: Account already verified
```

### POST /auth/resend-otp
```
Request:
{
    "email": "budi@email.com"
}

Response 200:
{
    "message": "OTP sent to your email"
}

Errors:
- 404: User not found
- 409: Account already verified
```

### POST /auth/login
```
Request:
{
    "email_or_phone": "budi@email.com",
    "password": "..."
}

Response 200:
{
    "user": { "id", "full_name", "email", "role" },
    "token": "eyJhbGci...",
    "expires_at": "2026-06-09T10:30:00Z"
}

Errors:
- 401: Invalid credentials
- 403: Account suspended
```

### POST /auth/refresh
```
Header: Authorization: Bearer <token>

Response 200:
{
    "token": "eyJ...",
    "expires_at": "..."
}
```

### POST /auth/logout
```
Header: Authorization: Bearer <token>

Response 204: No Content
```

---

## 3. Venture Endpoints

### POST /ventures
```
Request:
{
    "name": "Warung Nasi Ayam Geprek",
    "category": "makanan_berat",     // optional
    "region": "Jakarta Pusat"         // optional
}

Response 201:
{
    "id": "uuid",
    "name": "...",
    "category": "...",
    "region": "...",
    "stage": "draft",
    "current_version": 1,
    "created_at": "..."
}
```

### GET /ventures
```
Query: ?page=1&limit=20&stage=draft

Response 200:
{
    "data": [
        {
            "id", "name", "category", "region",
            "stage", "current_version",
            "latest_score": { "total_score", ... },
            "mission_stats": { "completed", "total" },
            "created_at", "updated_at"
        }
    ],
    "meta": { "page": 1, "limit": 20, "total": 5, "total_pages": 1 }
}
```

### GET /ventures/:id
```
Response 200:
{
    "id", "name", "category", "region",
    "stage", "current_version",
    "owner_user_id",
    "created_at", "updated_at",
    // embedded if available:
    "idea": { ... },
    "customer_segment": { ... },
    "menus": [...],
    "latest_score": { ... }
}
```

### PUT /ventures/:id
```
Request:
{
    "name": "Nama Baru",
    "category": "minuman",
    "region": "Bandung"
}

Response 200: updated venture
```

---

## 4. Idea Endpoints

### POST /ventures/:id/idea
```
Request:
{
    "raw_input": "Saya mau jual nasi goreng homemade... minimal 20 char"
}

Response 201:
{
    "id": "uuid",
    "raw_input": "...",
    "status": "pending"  // pending | processing | done | failed
}
```

### POST /ventures/:id/idea/process
```
Trigger AI structuring. Async — return immediately, poll status.

Response 202:
{
    "id": "uuid",
    "status": "processing",
    "estimated_seconds": 15
}
```

### GET /ventures/:id/idea
```
Response 200:
{
    "id", "venture_id", "raw_input",
    "one_line_concept": "...",
    "target_customer": "...",
    "value_proposition": "...",
    "key_assumptions": "...",
    "early_risks": "...",
    "version": 1,
    "is_locked": false,
    "status": "done",       // pending | processing | done | failed
    "created_at", "updated_at"
}

Response 404: No idea yet
```

### PUT /ventures/:id/idea
```
Update structured concept fields before locking.

Request:
{
    "one_line_concept": "Edited concept...",
    "target_customer": "..."
}

Response 200: updated idea
```

### POST /ventures/:id/idea/confirm
```
Lock stage → IDEA_DEFINED. Creates version snapshot.

Response 200:
{
    "idea": { ... },
    "venture": { "stage": "idea_defined", "current_version": 1 }
}
```

---

## 5. Customer Segment Endpoints

### POST /ventures/:id/customer-segments
```
Request:
{
    "name": "Karyawan Kantoran",
    "age_range": "25-40",
    "buy_context": "Makan siang di kantor",
    "budget_range": "15000-20000",
    "location": "Area Pasar Baru Jakarta",
    "consumption_moment": "Senin-Jumat jam 11.30-13.00"
}

Response 201: customer_segment + validation warning if too general
```

### GET /ventures/:id/customer-segments
```
Response 200: [customer_segments...]
```

### PUT /ventures/:id/customer-segments/:sid
```
Request: partial update fields
Response 200: updated segment
```

### POST /ventures/:id/customer-segments/:sid/confirm
```
Lock customer → stage CUSTOMER_DEFINED.
Response 200: updated venture stage
```

---

## 6. Menu Endpoints

### POST /ventures/:id/menus
```
Request:
{
    "name": "Nasi Goreng AYGE",
    "description": "Nasi goreng dengan topping ayam geprek"
}

Response 201: menu (complexity_score = null until scored)
```

### GET /ventures/:id/menus
```
Query: ?status=active

Response 200: [menu...]
```

### PUT /ventures/:id/menus/:mid
```
Request:
{
    "status": "active",       // candidate | active | deferred | dropped
    "is_hero": true
}
Response 200
```

### DELETE /ventures/:id/menus/:mid
```
Response 204
```

### POST /ventures/:id/menus/score
```
Trigger AI complexity scoring for all candidate menus.

Response 202:
{
    "status": "processing",
    "estimated_seconds": 10
}
```

### POST /ventures/:id/menus/focus
```
Lock menu selection → stage SKU_FOCUSED.
All non-selected menus → 'deferred'.

Request: empty (uses current selection)
Response 200: venture with stage 'sku_focused'
Errors:
- 422: No active menu selected
- 422: More than 3 active menus
```

---

## 7. Cost/Ingredient Endpoints

### POST /ventures/:id/ingredients
```
Request:
{
    "menu_id": "uuid",
    "name": "Nasi",
    "unit": "gram",
    "quantity": 200,
    "unit_price": 1
}
Response 201
```

### GET /ventures/:id/ingredients
```
Query: ?menu_id=uuid
Response 200: [ingredient...]
```

### PUT /ventures/:id/ingredients/:iid
```
Response 200: updated
```

### DELETE /ventures/:id/ingredients/:iid
```
Response 204
```

### POST /ventures/:id/packaging
```
Request:
{
    "menu_id": "uuid",
    "name": "Box Makan",
    "unit_price": 1500,
    "quantity_per_pcs": 1
}
Response 201
```

### GET /ventures/:id/cost/summary
```
Response 200:
[
    {
        "menu_id": "uuid",
        "menu_name": "Nasi Goreng AYGE",
        "hpp_per_porsi": 3850,
        "suggested_price": 15000,
        "target_margin": 74.3,
        "margin_status": "sehat",
        "break_even_unit": 25,
        "labor_per_unit": 1000,
        "overhead_per_unit": 500
    }
]
```

### POST /ventures/:id/cost/confirm
```
Lock cost data → stage COST_EVALUATED.
Also save packaging and labor data first.

Response 200: venture with stage 'cost_evaluated'
```

---

## 8. Mission Endpoints

### POST /ventures/:id/missions/generate
```
Trigger AI mission generation based on current stage + risk profile.

Response 202:
{
    "status": "processing",
    "estimated_seconds": 10
}
```

### GET /ventures/:id/missions
```
Query: ?status=pending,accepted&type=polling&page=1&limit=20

Response 200:
{
    "data": [
        {
            "id", "title", "description", "mission_type",
            "priority", "status", "due_at",
            "evidence_required", "estimated_minutes",
            "created_by", "sort_order",
            "evidence_count": 2,
            "latest_evidence_verdict": "valid"
        }
    ],
    "meta": { ... }
}
```

### POST /ventures/:id/missions
```
Manual mission creation (founder or mentor).

Request:
{
    "title": "...",
    "description": "...",
    "mission_type": "observation",
    "priority": "medium",
    "due_at": "2026-06-10T17:00:00Z",
    "evidence_required": "...",
    "estimated_minutes": 30
}
Response 201
```

### POST /ventures/:id/missions/:mid/accept
```
Response 200: mission status → 'accepted'
Also auto-transition venture stage to MISSION_ACTIVE if first mission accepted.
```

### POST /ventures/:id/missions/:mid/complete
```
Validate: must have at least 1 evidence uploaded.
Then mark status → 'completed'.

Response 200
Errors:
- 422: No evidence uploaded
```

### PUT /ventures/:id/missions/:mid
```
Update mission (status, priority, etc.)
Response 200
```

---

## 9. Evidence Endpoints

### POST /ventures/:id/evidence
```
Multipart form:
- mission_id (required): UUID
- evidence_type (required): image | text | link
- file (optional): image file (max 5MB, jpeg/png)
- text_content (optional): text evidence
- link_url (optional): URL evidence

Response 201:
{
    "id", "mission_id", "evidence_type",
    "storage_url", "thumbnail_url",
    "file_size_bytes", "mime_type",
    "submitted_at",
    "review_status": "pending"
}

Errors:
- 413: File too large
- 415: Unsupported file type
- 422: No file/text/link provided
```

### GET /ventures/:id/evidence
```
Query: ?mission_id=uuid&type=image&page=1&limit=20

Response 200: { data: [evidence...], meta: {...} }
```

### GET /ventures/:id/evidence/:eid
```
Response 200:
{
    "id", "mission_id", "venture_id",
    "evidence_type", "storage_url", "thumbnail_url",
    "text_content", "file_size_bytes", "mime_type",
    "submitted_at",
    "review": { verdict, rationale, next_action }
}
```

---

## 10. Review Endpoints

### POST /ventures/:id/evidence/:eid/review
```
Trigger AI review for a specific evidence.

Response 202:
{
    "status": "processing",
    "estimated_seconds": 30
}
```

### POST /ventures/:id/evidence/:eid/review/override
```
Admin/mentor override.

Request:
{
    "verdict": "valid",        // valid | weak | invalid | suspicious
    "rationale": "..."
}

Response 200:
{
    "id", "evidence_id", "verdict",
    "rationale", "reviewer_type": "human_override",
    "overridden_by": "admin-uuid"
}
```

---

## 11. Score Endpoints

### GET /ventures/:id/score
```
Response 200:
{
    "id", "venture_id",
    "clarity_score": 82,
    "focus_score": 80,
    "economics_score": 90,
    "execution_score": 60,
    "evidence_score": 50,
    "market_response_score": 45,
    "total_score": 72,
    "is_final": false,
    "created_at": "..."
}

404: No score yet
```

### POST /ventures/:id/score/calculate
```
Trigger score calculation (updates if evidence changed).

Response 200: score object (as above)
```

### GET /ventures/:id/score/history
```
Response 200: [score...] ordered by created_at desc
```

---

## 12. Decision Endpoints

### POST /ventures/:id/decision/generate
```
Generate final decision based on score components.

Response 200:
{
    "id", "venture_id",
    "decision": "continue",      // continue | repeat | pivot | stop
    "rationale": "...",
    "score_snapshot": { ... },   // copy of score at decision time
    "created_at": "..."
}

Errors:
- 422: Score components incomplete
- 409: Decision already generated
```

### GET /ventures/:id/decision
```
Response 200: { decision object }
```

---

## 13. Notification Endpoints

### GET /notifications
```
Query: ?is_read=false&type=mission_new&page=1&limit=20

Response 200:
{
    "data": [
        {
            "id", "user_id", "venture_id",
            "type", "title", "body",
            "is_read": false,
            "metadata": { "mission_id": "uuid" },
            "created_at": "..."
        }
    ],
    "meta": { ... }
}
```

### PUT /notifications/:id/read
```
Response 200: { is_read: true }
```

### PUT /notifications/read-all
```
Mark all as read for current user.
Response 200: { affected: 5 }
```

---

## 14. Mentor Endpoints

### GET /mentor/mentees
```
Response 200:
{
    "data": [
        {
            "founder": { "id", "full_name", "avatar_url" },
            "venture": { "id", "name", "stage" },
            "latest_score": { "total_score", ... },
            "mission_stats": { "completed": 3, "total": 5 },
            "last_active_at": "..."
        }
    ]
}
```

### GET /mentor/mentees/:venture_id
```
Response 200: full venture detail + founder info
```

### POST /mentor/mentees/:venture_id/comments
```
Request:
{
    "mission_id": "uuid",     // optional
    "evidence_id": "uuid",    // optional
    "content": "Coba perbaiki kualitas fotonya"
}
Response 201
```

### GET /mentor/mentees/:venture_id/comments
```
Response 200: [comments...]
```

---

## 15. Health & Admin

### GET /health
```
Response 200:
{
    "status": "ok",
    "version": "1.0.0",
    "db": "connected",
    "llm": "available",
    "uptime_seconds": 12345
}
```

---

## 16. Standard Error Codes

| Code | HTTP | Meaning |
|---|---|---|
| `ERR_AUTH_001` | 401 | Invalid token |
| `ERR_AUTH_002` | 401 | Token expired |
| `ERR_AUTH_003` | 403 | Insufficient role |
| `ERR_AUTH_004` | 401 | Invalid credentials |
| `ERR_VAL_001` | 422 | Validation failed |
| `ERR_VAL_002` | 422 | Duplicate entry |
| `ERR_VAL_003` | 422 | Business rule violation (e.g. max 3 SKU) |
| `ERR_NF_001` | 404 | Resource not found |
| `ERR_INT_001` | 500 | Internal server error |
| `ERR_LLM_001` | 503 | AI service unavailable |
| `ERR_UPL_001` | 413 | File too large |
| `ERR_UPL_002` | 415 | Unsupported file type |
| `ERR_GATE_001` | 422 | Stage gate not satisfied |
