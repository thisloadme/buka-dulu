# Production Stage Checklist — BukaDulu Go-Live
## Versi 1.0 | Target: Production Launch

> Checklist kesiapan go-live production berdasarkan seluruh user story (US-001 s.d. US-121), SRS, dan PRD.
> Setiap item punya status: ✅ Done / ⬜ Pending / 🟡 In Progress / ❌ Blocked / N/A Not Applicable

---

## 1. Infrastructure & Environment

### 1.1 Server & Hosting
| # | Item | Ref | Status | Notes |
|---|---|---|---|---|
| 1.1.1 | Production server provisioned (CPU/RAM/disk sesuai beban) | NFR-010 | ⬜ | Minimal 2 vCPU, 4GB RAM |
| 1.1.2 | Database production instance running (PostgreSQL) | NFR-010 | ⬜ | Managed PostgreSQL (RDS/Cloud SQL) |
| 1.1.3 | Object storage untuk evidence files | FR-070 | ⬜ | S3/MinIO/GCS dengan signed URL |
| 1.1.4 | SMTP/email service configured | FR-001 | ⬜ | SendGrid/Mailgun/SES untuk OTP & notifikasi |
| 1.1.5 | Domain + SSL certificate (HTTPS) | NFR-031 | ⬜ | Let's Encrypt atau managed |
| 1.1.6 | CDN untuk static assets | NFR-001 | ⬜ | Optional untuk MVP |
| 1.1.7 | LLM API provider configured (production key) | FR-021 | ⬜ | OpenAI/Anthropic/Gemini |
| 1.1.8 | Queue/worker infrastructure untuk async review | FR-080 | ⬜ | Redis + worker atau managed queue |

### 1.2 Environment Configuration
| # | Item | Ref | Status | Notes |
|---|---|---|---|---|
| 1.2.1 | `.env.production` with all secrets | SRS §13 | ⬜ | JWT_SECRET, DATABASE_URL, SMTP, LLM_API_KEY, etc |
| 1.2.2 | Secrets management (not in repo) | NFR-030 | ⬜ | Vault / AWS Secrets Manager / env vars |
| 1.2.3 | Config validation on startup | SRS §13 | ⬜ | App must fail fast on missing config |
| 1.2.4 | CORS configuration for production domain | - | ⬜ | Whitelist only production frontend origin |
| 1.2.5 | Rate limiting configured per endpoint | NFR-033 | ⬜ | Auth endpoints: strict; others: moderate |

### 1.3 Database
| # | Item | Ref | Status | Notes |
|---|---|---|---|---|
| 1.3.1 | Production migration run (001_init.up.sql) | SRS §13 | ⬜ | All tables created |
| 1.3.2 | Database backup configured (daily) | SRS §13.3 | ⬜ | Automated backup + retention policy |
| 1.3.3 | Connection pool tuning | NFR-001 | ⬜ | Max connections sesuai instance size |
| 1.3.4 | Database migration rollback tested | SRS §13 | ⬜ | Down migration must work |
| 1.3.5 | Database monitoring active | NFR-070 | ⬜ | Slow query log, connection count, disk usage |
| 1.3.6 | Index review for query performance | NFR-001 | ⬜ | All indexes from migration validated |

---

## 2. Feature Readiness — Per User Story

### 2.1 Epic 1: Authentication & Account (P0)

| # | User Story | Item | Status | Notes |
|---|---|---|---|---|
| US-001 | Registrasi | Form registrasi (email/telepon + password) | ✅ | Done in Sprint 1 |
| US-001 | Registrasi | Password min 8 karakter | ✅ | Validated in handler |
| US-001 | Registrasi | Validasi format email & telepon | ✅ | Server-side validation |
| US-001 | Registrasi | Role otomatis "founder" | ✅ | Default role |
| US-001 | Registrasi | Duplicate email/telepon ditolak | ✅ | ErrDuplicateEntry |
| US-001 | Registrasi | **OTP email verification** | ✅ | **Done — lihat migration 002 + email service** |
| US-001 | Registrasi | Error handling koneksi terputus | 🟡 | Perlu timeout config |
| US-002 | Login & Logout | Login email/telepon + password | ✅ | Done in Sprint 1 |
| US-002 | Login & Logout | JWT token-based session | ✅ | JWT with 24h expiry |
| US-002 | Login & Logout | Logout hapus session | ✅ | Token invalidation |
| US-002 | Login & Logout | Session expired 24 jam | ✅ | Configurable JWT expiry |
| US-002 | Login & Logout | Auto-redirect jika token invalid | ✅ | Middleware check |

### 2.2 Epic 2: Venture Workspace (P0)

| # | User Story | Item | Status | Notes |
|---|---|---|---|---|
| US-010 | Membuat Venture | Input: nama, kategori F&B, lokasi | ✅ | Done in Sprint 1 |
| US-010 | Membuat Venture | Venture dengan stage DRAFT | ✅ | Default stage |
| US-010 | Membuat Venture | Multi-venture support | ✅ | Per user |
| US-010 | Membuat Venture | Stage tercatat di database | ✅ | ventures table |
| US-011 | Mengelola Venture | Edit nama, kategori, lokasi, status | ✅ | Done in Sprint 1 |
| US-011 | Mengelola Venture | Hanya owner yang bisa edit | ✅ | Owner check in service |
| US-011 | Mengelola Venture | Timestamp update | ✅ | updated_at |

### 2.3 Epic 3: Idea Intake & Structuring (P0)

| # | User Story | Item | Status | Notes |
|---|---|---|---|---|
| US-020 | Capture Ide Mentah | Textarea bebas minimal 20 karakter | ✅ | Done in Sprint 1 |
| US-020 | Capture Ide Mentah | Placeholder contoh | ✅ | |
| US-020 | Capture Ide Mentah | Tombol "Proses Ide" | ✅ | |
| US-020 | Capture Ide Mentah | Loading state AI processing | ✅ | |
| US-021 | Generate Structured Concept | AI: one-line concept, target customer, value prop, assumptions, risks | ✅ | Done in Sprint 1 |
| US-021 | Generate Structured Concept | Output structured card per elemen | ✅ | |
| US-021 | Generate Structured Concept | Setiap elemen bisa diedit manual | ✅ | US-022 |
| US-021 | Generate Structured Concept | Loading state | ✅ | |
| US-021 | Generate Structured Concept | Error message + retry jika AI gagal | ✅ | |
| US-021 | Generate Structured Concept | Raw input/output AI disimpan | ✅ | ai_raw_input, ai_raw_output |
| US-022 | Edit & Konfirmasi Konsep | Semua field bisa diedit | ✅ | Done in Sprint 1 |
| US-022 | Edit & Konfirmasi Konsep | Tombol Konfirmasi → stage IDEA_DEFINED | ✅ | |
| US-022 | Edit & Konfirmasi Konsep | Versi tersimpan sebagai baseline | ✅ | |
| US-022 | Edit & Konfirmasi Konsep | Versi lama bisa dilihat | ✅ | venture_versions |
| US-022 | Edit & Konfirmasi Konsep | Tidak bisa kembali tanpa reset | ✅ | Stage gate check |

### 2.4 Epic 4: Customer Hypothesis (P0)

| # | User Story | Item | Status | Notes |
|---|---|---|---|---|
| US-030 | Define Customer Segment | Input: usia, konteks beli, budget, lokasi, momen konsumsi | ✅ | Done in Sprint 2 |
| US-030 | Define Customer Segment | AI segment suggestion | ✅ | |
| US-030 | Define Customer Segment | Flag jika terlalu umum | ✅ | is_too_general |
| US-030 | Define Customer Segment | Konfirmasi → CUSTOMER_DEFINED | ✅ | |

### 2.5 Epic 5: Menu Focus Engine (P0)

| # | User Story | Item | Status | Notes |
|---|---|---|---|---|
| US-040 | Tambah Kandidat Menu | Form: nama, deskripsi, bahan utama | ✅ | Done in Sprint 2 |
| US-040 | Tambah Kandidat Menu | Multi-menu | ✅ | |
| US-040 | Tambah Kandidat Menu | Daftar menu sebagai cards | ✅ | |
| US-040 | Tambah Kandidat Menu | Hapus menu kandidat | ✅ | |
| US-041 | Skor & Rekomendasi Hero SKU | Skor kompleksitas per menu | ✅ | Done in Sprint 2 |
| US-041 | Skor & Rekomendasi Hero SKU | Rekomendasi 1 hero SKU | ✅ | |
| US-041 | Skor & Rekomendasi Hero SKU | Maksimal 3 SKU aktif | ✅ | |
| US-041 | Skor & Rekomendasi Hero SKU | Alasan rekomendasi | ✅ | |
| US-042 | Pilih & Kunci SKU | Minimal 1 SKU wajib | ✅ | Done in Sprint 2 |
| US-042 | Pilih & Kunci SKU | Maks 3 SKU — blokir | ✅ | |
| US-042 | Pilih & Kunci SKU | Menu ditunda masuk daftar deferred | ✅ | |
| US-042 | Pilih & Kunci SKU | Konfirmasi → SKU_FOCUSED | ✅ | |

### 2.6 Epic 6: Cost & Margin Engine (P0)

| # | User Story | Item | Status | Notes |
|---|---|---|---|---|
| US-050 | Input Biaya Bahan | Daftar bahan: nama, unit, kuantitas, harga | ✅ | Done in Sprint 2 |
| US-050 | Input Biaya Bahan | Biaya kemasan per porsi | ✅ | |
| US-050 | Input Biaya Bahan | Estimasi tenaga kerja & overhead | ✅ | |
| US-050 | Input Biaya Bahan | Validasi angka positif | ✅ | |
| US-051 | Hitung HPP & Margin | HPP per porsi otomatis | ✅ | Done in Sprint 2 |
| US-051 | Hitung HPP & Margin | Harga jual minimum + saran harga tes | ✅ | |
| US-051 | Hitung HPP & Margin | Klasifikasi margin 🟢🟡🔴 | ✅ | |
| US-051 | Hitung HPP & Margin | Estimasi break-even unit | ✅ | |
| US-051 | Hitung HPP & Margin | Perhitungan realtime | ✅ | |
| US-051 | Hitung HPP & Margin | Konfirmasi → COST_EVALUATED | ✅ | |

### 2.7 Epic 7: Mission Orchestrator (P0)

| # | User Story | Item | Status | Notes |
|---|---|---|---|---|
| US-060 | Generate Misi Harian | Generate minimal 1 misi/hari | ✅ | Done in Sprint 3 |
| US-060 | Generate Misi Harian | Tiap misi: tujuan, estimasi waktu, prioritas, deadline, definisi bukti | ✅ | |
| US-060 | Generate Misi Harian | Mission board | ✅ | |
| US-060 | Generate Misi Harian | Bisa regenerate misi | ✅ | |
| US-061 | Terima & Jalankan Misi | Tombol "Terima Misi" | ✅ | Done in Sprint 3 |
| US-061 | Terima & Jalankan Misi | Bisa tunda misi | ✅ | |
| US-061 | Terima & Jalankan Misi | Misi manual (founder/mentor) | ✅ | |
| US-062 | Selesaikan Misi dengan Bukti | Tombol "Selesai" disabled tanpa evidence | ✅ | Done in Sprint 3 |
| US-062 | Selesaikan Misi dengan Bukti | Flow: upload → review → complete | ✅ | |
| US-062 | Selesaikan Misi dengan Bukti | Stage → MISSION_ACTIVE | ✅ | |

### 2.8 Epic 8: Evidence Management (P0)

| # | User Story | Item | Status | Notes |
|---|---|---|---|---|
| US-070 | Upload Evidence | Upload: image (jpeg/png, max 5MB), text, link | ✅ | Done in Sprint 3 |
| US-070 | Upload Evidence | Multiple evidence per mission | ✅ | |
| US-070 | Upload Evidence | Auto-attach ke mission | ✅ | |
| US-070 | Upload Evidence | Progress upload | 🟡 | **Mobile: camera/image belum functional** |
| US-070 | Upload Evidence | Jika upload gagal, data form tidak hilang | ✅ | |
| US-070 | Upload Evidence | Tipe evidence & timestamp tersimpan | ✅ | |
| US-071 | Lihat Evidence | Galeri evidence per venture | ✅ | Done in Sprint 3 |
| US-071 | Lihat Evidence | Filter by mission, tipe, status review | ✅ | |
| US-071 | Lihat Evidence | Preview image inline | 🟡 | **Mobile: preview perlu dioptimalkan** |

### 2.9 Epic 9: Evidence Review Engine (P0)

| # | User Story | Item | Status | Notes |
|---|---|---|---|---|
| US-080 | Auto Review Evidence | AI menilai: ✅ Valid / ⚠️ Weak / ❌ Invalid / 🤔 Suspicious | ✅ | Done in Sprint 3 |
| US-080 | Auto Review Evidence | Alasan dalam Bahasa Indonesia | ✅ | |
| US-080 | Auto Review Evidence | Next action: continue/repeat/pivot | ✅ | |
| US-080 | Auto Review Evidence | Review async — tidak ngeblock | ✅ | |
| US-080 | Auto Review Evidence | 90% review selesai dalam 60 detik | 🟡 | **Perlu monitoring latency** |
| US-080 | Auto Review Evidence | Status review (pending/processing/done) | ✅ | |
| US-081 | Human Override | Admin bisa override verdict | ✅ | Done in Sprint 3 |
| US-081 | Human Override | Audit trail override | ✅ | |
| US-081 | Human Override | Override triggers recalculation | ✅ | |

### 2.10 Epic 10: Scoring & Decision Engine (P0)

| # | User Story | Item | Status | Notes |
|---|---|---|---|---|
| US-090 | Lihat Score | 6 komponen score | ✅ | Done in Sprint 4 |
| US-090 | Lihat Score | Score update tiap bukti baru | ✅ | |
| US-090 | Lihat Score | Breakdown komponen | ✅ | |
| US-090 | Lihat Score | Visual progress bar | ✅ | |
| US-091 | Dapat Keputusan Akhir | Keputusan hanya jika semua komponen terisi | ✅ | Done in Sprint 4 |
| US-091 | Dapat Keputusan Akhir | Output: continue/repeat/pivot/stop | ✅ | |
| US-091 | Dapat Keputusan Akhir | Rationale transparan | ✅ | |
| US-091 | Dapat Keputusan Akhir | Data pendukung | ✅ | |
| US-091 | Dapat Keputusan Akhir | Stage → CONTINUE/REPEAT/PIVOT/STOP | ✅ | |

### 2.11 Epic 11: Founder Courtroom (MVP-lite)

| # | User Story | Item | Status | Notes |
|---|---|---|---|---|
| US-100 | Jalankan Courtroom Review | 3 perspektif: pembeli, operator, reviewer | ⬜ | **BELUM IMPLEMENTASI — P1** |
| US-100 | Jalankan Courtroom Review | Objection list + weakest assumptions | ⬜ | |
| US-100 | Jalankan Courtroom Review | Action items | ⬜ | |
| US-100 | Jalankan Courtroom Review | Bisa dijalankan kapan saja post-idea | ⬜ | |

### 2.12 Epic 12: Notifications (P0)

| # | User Story | Item | Status | Notes |
|---|---|---|---|---|
| US-110 | Dapat Notifikasi | Notifikasi: mission deadline, review selesai, score update | 🟡 | **DB table ada, API endpoint belum** |
| US-110 | Dapat Notifikasi | Pasif >48 jam → reactivation prompt | ⬜ | **Belum implementasi** |
| US-110 | Dapat Notifikasi | In-app notification | 🟡 | **Mobile: belum ada UI notifikasi** |

### 2.13 Epic 13: Mentor Dashboard (P1)

| # | User Story | Item | Status | Notes |
|---|---|---|---|---|
| US-120 | Lihat Mentee Progress | List mentee dengan stage, score, mission status | ✅ | Done in Sprint 4 |
| US-120 | Lihat Mentee Progress | Filter by stage, status, recency | ✅ | |
| US-120 | Lihat Mentee Progress | Drill-down ke detail venture | ✅ | |
| US-121 | Komentar Mentor | Komentar per mission/evidence | ✅ | Done in Sprint 4 |
| US-121 | Komentar Mentor | Founder dapat notifikasi | 🟡 | **Notifikasi belum terintegrasi** |
| US-121 | Komentar Mentor | Komentar tidak bisa diedit (soft delete) | ✅ | |

---

## 3. Security & Compliance

| # | Item | Ref | Status | Notes |
|---|---|---|---|---|
| 3.1 | All API endpoints use HTTPS only | NFR-031 | ⬜ | HSTS header recommended |
| 3.2 | Password hashing with bcrypt | NFR-030 | ✅ | Already implemented |
| 3.3 | JWT signing with strong secret | NFR-033 | ✅ | HMAC-SHA256 |
| 3.4 | Authorization per role & per resource | NFR-033 | ✅ | Middleware auth |
| 3.5 | OTP email verification on registration | FR-001 | ✅ | **Done — auth/verify-otp + auth/resend-otp endpoints** |
| 3.6 | Signed URLs for evidence object storage | NFR-032 | ⬜ | **Perlu implementasi** |
| 3.7 | Input sanitization on all endpoints | NFR-033 | 🟡 | Basic done, audit needed |
| 3.8 | File upload size & type validation | FR-070 | ✅ | Max 5MB, jpeg/png |
| 3.9 | Rate limiting on auth endpoints | NFR-033 | ⬜ | Prevent brute force |
| 3.10 | Audit log for sensitive operations | FR-084, FR-131 | ✅ | audit_logs table |
| 3.11 | SQL injection prevention (parameterized queries) | NFR-033 | ✅ | Using pgx parameterized |
| 3.12 | CORS restricted to known origins | - | 🟡 | Need production config |
| 3.13 | Session/token expiry enforcement | FR-003 | ✅ | 24h JWT expiry |
| 3.14 | Data retention & account deletion policy | NFR-041 | ⬜ | **Belum ada mekanisme hapus akun** |
| 3.15 | User consent: data processed by AI | NFR-040 | ⬜ | **Belum ada TOS/privacy notice** |

---

## 4. Performance & Reliability

| # | Item | Ref | Target | Status | Notes |
|---|---|---|---|---|---|
| 4.1 | 95% read API requests < 500ms | NFR-001 | <500ms | ⬜ | Need load testing |
| 4.2 | 95% write API requests < 800ms | NFR-002 | <800ms | ⬜ | Need load testing |
| 4.3 | Evidence review < 60s for 90% cases | NFR-003 | <60s | ⬜ | Monitor AI latency |
| 4.4 | 99.5% availability per month | NFR-010 | 99.5% | ⬜ | Need SLA monitoring |
| 4.5 | Handle 10,000 active ventures | NFR-020 | 10K | ⬜ | Need scale testing |
| 4.6 | Score calculation idempotent & reproducible | NFR-050 | ✅ | Already designed |
| 4.7 | AI failure → no data loss | NFR-051 | ✅ | Raw input saved before AI |
| 4.8 | Database connection pool configured | - | ⬜ | Tune for production load |
| 4.9 | Static assets served via CDN | - | ⬜ | Optional MVP |
| 4.10 | Mobile app performance: cold start <3s | - | <3s | ⬜ | Flutter profile mode test |

---

## 5. Monitoring & Observability

| # | Item | Ref | Status | Notes |
|---|---|---|---|---|
| 5.1 | Structured logging (JSON) | NFR-070 | ✅ | slog in place |
| 5.2 | Application metrics (request rate, latency, errors) | NFR-070 | ⬜ | Prometheus + Grafana |
| 5.3 | Error tracking & alerting | NFR-070 | ⬜ | Sentry / similar |
| 5.4 | Uptime monitoring | NFR-010 | ⬜ | Health endpoint ready: /health |
| 5.5 | Database monitoring (slow queries, connections) | NFR-070 | ⬜ | |
| 5.6 | LLM API latency & error rate monitoring | - | ⬜ | |
| 5.7 | Async job processing status tracking | NFR-071 | ⬜ | Queue monitoring |
| 5.8 | Server resource monitoring (CPU, RAM, disk) | - | ⬜ | |
| 5.9 | Alert on error rate spike >1% | - | ⬜ | |
| 5.10 | Dashboard untuk product metrics (signup, venture, mission, evidence) | FR-130 | ⬜ | Event tracking + analytics |

---

## 6. Deployment & DevOps

| # | Item | Ref | Status | Notes |
|---|---|---|---|---|
| 6.1 | CI/CD pipeline (test → build → deploy) | SRS §13.2 | ⬜ | GitHub Actions / GitLab CI |
| 6.2 | Automated tests run on PR | SRS §13.2 | ⬜ | |
| 6.3 | Migration check in CI | SRS §13.2 | ⬜ | |
| 6.4 | Linting in CI | SRS §13.2 | 🟡 | Flutter: analysis_options.yaml ada |
| 6.5 | Deploy to staging before production | SRS §13.2 | ⬜ | |
| 6.6 | Zero-downtime deployment strategy | - | ⬜ | Blue-green atau rolling |
| 6.7 | Database migration automated in deploy | - | ⬜ | |
| 6.8 | Rollback plan documented | - | ⬜ | |
| 6.9 | Environment separation (local/staging/prod) | SRS §13.1 | 🟡 | Config supports, infra needed |
| 6.10 | Docker containerization | - | ⬜ | Dockerfile + compose |
| 6.11 | Health check endpoint integrated with orchestration | - | ✅ | /health endpoint exists |
| 6.12 | Secrets not exposed in logs/config | - | 🟡 | Password masked in logs |

---

## 7. Data Backup & Disaster Recovery

| # | Item | Ref | Status | Notes |
|---|---|---|---|---|
| 7.1 | Automated daily database backup | SRS §13.3 | ⬜ | |
| 7.2 | Backup retention policy defined | SRS §13.3 | ⬜ | 30 days recommended |
| 7.3 | Backup restoration tested | - | ⬜ | Quarterly test |
| 7.4 | Evidence files backup (object storage) | - | ⬜ | Typically redundant by provider |
| 7.5 | Disaster recovery plan documented | - | ⬜ | RTO/RPO defined |
| 7.6 | Multi-region fallback (if needed) | - | ⬜ | Phase 2 |

---

## 8. Testing & Quality Assurance

| # | Item | Ref | Status | Notes |
|---|---|---|---|---|
| 8.1 | Unit tests for cost calculation | SRS §12.2 | ⬜ | |
| 8.2 | Unit tests for score calculation | SRS §12.2 | ⬜ | |
| 8.3 | Integration tests: mission → evidence → review | SRS §12.2 | ⬜ | |
| 8.4 | API tests for all endpoints | SRS §12.2 | ⬜ | 40+ endpoints |
| 8.5 | E2E test for 14-day flow | SRS §12.2 | ⬜ | |
| 8.6 | Security test: auth & media access | SRS §12.2 | ⬜ | |
| 8.7 | Mobile UI test (Flutter widget test) | - | ⬜ | |
| 8.8 | Load test with expected traffic | - | ⬜ | |
| 8.9 | Edge case: AI timeout/failure | - | ⬜ | |
| 8.10 | Edge case: concurrent evidence upload | - | ⬜ | |
| 8.11 | Founder completes full flow without admin intervention | SRS §12.1 | 🟡 | **Courtroom & notifications missing** |

---

## 9. Mobile App (Flutter) — Play Store Readiness

| # | Item | Ref | Status | Notes |
|---|---|---|---|---|
| 9.1 | App signing configured | - | ⬜ | |
| 9.2 | App icons & splash screen production-ready | - | ⬜ | |
| 9.3 | Build variants: debug vs release | - | ⬜ | |
| 9.4 | API base URL configurable by flavor | - | 🟡 | dotenv configured |
| 9.5 | Google Play Billing (IAP) for subscription | - | ✅ | **Done — `in_app_purchase` + `play_billing`. Disabled in dev via kReleaseMode gate** |
| 9.6 | Google Play Billing disabled in dev/local builds | - | ✅ | **kReleaseMode check + yellow warning banner** |
| 9.7 | Push notification setup (FCM) | US-110 | ⬜ | |
| 9.8 | Minimum OS version configured | - | 🟡 | Check current config |
| 9.9 | App privacy policy (data collection disclosure) | - | ⬜ | Required by Google Play |
| 9.10 | In-app camera for evidence upload | FR-070 | 🟡 | **Mobile: belum functional** |
| 9.11 | Flutter build in release mode tested | - | ⬜ | |
| 9.12 | App size optimized | - | ⬜ | |
| 9.13 | Deep link handling | - | ⬜ | For future push notif |

---

## 10. Legal & Compliance

| # | Item | Ref | Status | Notes |
|---|---|---|---|---|
| 10.1 | Terms of Service (ToS) | - | ⬜ | |
| 10.2 | Privacy Policy | NFR-040 | ⬜ | Including AI data processing |
| 10.3 | Cookie/analytics consent | - | ⬜ | If using analytics |
| 10.4 | Account deletion mechanism | NFR-041 | ⬜ | User can delete own data |
| 10.5 | Data retention policy documented | NFR-041 | ⬜ | |
| 10.6 | AI disclaimer: results are recommendations | SRS §2.5 | ⬜ | Display in UI |
| 10.7 | UU ITE / PDP compliance (Indonesia) | - | ⬜ | Personal data protection |

---

## 11. Operational Readiness

| # | Item | Ref | Status | Notes |
|---|---|---|---|---|
| 11.1 | On-call rotation defined | - | ⬜ | |
| 11.2 | Incident response runbook | - | ⬜ | |
| 11.3 | Known issues / FAQ documented | - | ⬜ | |
| 11.4 | Admin panel for manual override | FR-083 | ⬜ | **Perlu admin UI** |
| 11.5 | Customer support channel | - | ⬜ | Email/support form |
| 11.6 | LLM API fallback if primary provider down | - | ⬜ | |
| 11.7 | Usage/cost monitoring for LLM API | - | ⬜ | Token usage tracking |

---

## 12. Post-Launch Checklist (24-48h After Go-Live)

| # | Item | Status | Notes |
|---|---|---|---|
| 12.1 | Monitor error rate & fix critical bugs | ⬜ | |
| 12.2 | Verify database backup working | ⬜ | |
| 12.3 | Check LLM API latency under real load | ⬜ | |
| 12.4 | Monitor signup flow completion rate | ⬜ | |
| 12.5 | Check evidence upload success rate | ⬜ | |
| 12.6 | Verify email delivery (OTP, notifications) | ⬜ | |
| 12.7 | Review first 100 user journeys | ⬜ | |

---

## Summary

| Category | Total | ✅ Done | 🟡 Partial | ⬜ Pending | ❌ Blocked |
|---|---|---|---|---|---|
| 1. Infrastructure & Environment | 16 | 0 | 1 | 15 | 0 |
| 2. Feature Readiness (User Stories) | 63 | 55 | 4 | 4 | 0 |
| 3. Security & Compliance | 15 | 6 | 2 | 7 | 0 |
| 4. Performance & Reliability | 10 | 2 | 0 | 8 | 0 |
| 5. Monitoring & Observability | 10 | 1 | 0 | 9 | 0 |
| 6. Deployment & DevOps | 12 | 2 | 2 | 8 | 0 |
| 7. Data Backup & DR | 6 | 0 | 0 | 6 | 0 |
| 8. Testing & QA | 11 | 0 | 1 | 10 | 0 |
| 9. Mobile App Readiness | 13 | 2 | 2 | 9 | 0 |
| 10. Legal & Compliance | 7 | 0 | 0 | 7 | 0 |
| 11. Operational Readiness | 7 | 0 | 0 | 7 | 0 |
| 12. Post-Launch | 7 | 0 | 0 | 7 | 0 |
| **TOTAL** | **177** | **67** | **12** | **98** | **0** |

> **Status: 🟡 MVP code complete, but PRODUCTION GAPS exist.** 98 items still pending before safe go-live.
> Priority: signed URLs (#3.6), monitoring (#5), CI/CD (#6), testing (#8).
