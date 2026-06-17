# Software Requirements Specification (SRS)
## Sistem: BukaDulu
## Versi: 1.0
## Tanggal: 2026-06-02
## Status: Draft for Engineering Design

## 1. Pendahuluan
### 1.1 Tujuan Dokumen
Dokumen ini mendefinisikan kebutuhan perangkat lunak secara lengkap untuk sistem BukaDulu, sebuah aplikasi validasi eksekusi bisnis F&B tahap awal. Dokumen ini menjadi acuan bagi engineering, product, design, QA, dan AI/ML team dalam membangun MVP yang konsisten.

### 1.2 Ruang Lingkup Sistem
Sistem mendukung pengguna dari tahap ide mentah hingga keputusan lanjut, pivot, atau stop melalui serangkaian stage-gate berbasis evidence. Sistem menyediakan modul intake ide, pemfokusan menu, simulasi ekonomi dasar, generator misi, pengunggahan bukti, penilaian bukti, dan penentuan readiness score.

### 1.3 Definisi dan Istilah
- User: calon founder F&B yang menggunakan aplikasi.
- Mentor: pengguna pendamping yang memonitor founder.
- SKU: menu atau produk individual.
- Evidence: bukti yang diunggah pengguna untuk membuktikan aktivitas lapangan.
- Mission: tugas validasi terstruktur.
- Gate: syarat minimum agar user dapat lanjut ke tahap berikutnya.
- Readiness Score: skor kesiapan berbasis bukti dan parameter bisnis awal.
- Founder Courtroom: review adversarial terhadap asumsi pengguna.

### 1.4 Referensi Konseptual
Kategori aplikasi validator ide sudah memiliki banyak produk yang berfokus pada analisis dan laporan, sedangkan software restoran mayoritas fokus pada operasi bisnis yang sudah berjalan.[cite:7][cite:11][cite:20] BukaDulu didesain untuk menempati lapisan sebelum keduanya: lapisan eksekusi pra-peluncuran.

## 2. Deskripsi Umum
### 2.1 Perspektif Produk
BukaDulu adalah web/mobile-first application berbasis akun yang mengorkestrasi kombinasi rule engine, workflow state machine, dan layanan AI untuk membantu validasi bisnis awal.

### 2.2 Fungsi Tingkat Tinggi
Sistem harus mampu:
- Menerima input ide bebas.
- Menstrukturkan ide menjadi hipotesis bisnis.
- Membatasi fokus menu.
- Menghitung ekonomi dasar usaha.
- Membuat misi validasi harian.
- Menerima dan menilai evidence.
- Menghitung score dan keputusan akhir.
- Menyediakan audit trail progres pengguna.

### 2.3 Kelas Pengguna
#### A. Founder
Hak utama:
- Membuat workspace usaha.
- Mengedit profil ide.
- Menjalankan misi.
- Mengunggah evidence.
- Melihat hasil evaluasi.

#### B. Mentor
Hak utama:
- Melihat progress beberapa founder.
- Memberi komentar manual.
- Melihat evidence dan score.

#### C. Admin/Internal Ops
Hak utama:
- Mengelola template mission.
- Mengelola scoring parameters.
- Review evidence yang di-flag.
- Mengelola kategori dan prompt sistem.

### 2.4 Lingkungan Operasi
- Client: browser modern mobile dan desktop.
- Backend: REST API atau GraphQL service.
- Database: relational database untuk transactional data.
- Object storage: media evidence.
- AI service: LLM inference, classification, summarization.
- Queue/worker: asynchronous processing untuk scoring dan evidence review.

### 2.5 Batasan Sistem
- Sistem bukan POS.
- Sistem bukan aplikasi akuntansi penuh.
- Sistem tidak menjamin keberhasilan bisnis.
- Hasil AI bersifat rekomendasi, bukan keputusan hukum atau keuangan resmi.

### 2.6 Asumsi
- User mengakses internet stabil untuk unggah bukti.
- User mengerti bahasa Indonesia dasar.
- Data biaya yang dimasukkan user cukup mendekati realitas.

## 3. Arsitektur Konseptual
### 3.1 Komponen Utama
1. Authentication & User Management Service.
2. Venture Workspace Service.
3. Idea Structuring Service.
4. Menu Focus Engine.
5. Cost & Margin Engine.
6. Mission Orchestrator.
7. Evidence Management Service.
8. Evidence Review Engine.
9. Scoring & Decision Engine.
10. Notification Service.
11. Mentor Dashboard Service.
12. Audit Log & Analytics Service.

### 3.2 Pola Arsitektur
Arsitektur direkomendasikan menggunakan modular monolith pada MVP agar iterasi cepat, dengan pemisahan domain yang bersih untuk memungkinkan ekstraksi ke microservices di masa depan.

### 3.3 State Machine Inti
State venture:
- DRAFT
- IDEA_DEFINED
- CUSTOMER_DEFINED
- SKU_FOCUSED
- COST_EVALUATED
- MISSION_ACTIVE
- EVIDENCE_SUBMITTED
- EVIDENCE_REVIEWED
- READY_TO_DECIDE
- CONTINUE
- REPEAT
- PIVOT
- STOP

## 4. Kebutuhan Fungsional
### 4.1 Authentication and Account Management
#### FR-001 Registration
Sistem harus menyediakan registrasi menggunakan email atau nomor telepon.

#### FR-002 Login
Sistem harus menyediakan login aman menggunakan password atau OTP.

#### FR-003 Session Management
Sistem harus menjaga sesi pengguna dan menyediakan logout.

#### FR-004 Role Assignment
Sistem harus mendukung role founder, mentor, dan admin.

### 4.2 Venture Workspace
#### FR-010 Create Venture
Founder harus dapat membuat minimal satu venture/workspace bisnis.

#### FR-011 Edit Venture
Founder harus dapat mengedit nama venture, kategori usaha, lokasi, dan status.

#### FR-012 Venture Stage
Sistem harus menyimpan status stage venture dan histori perpindahannya.

### 4.3 Idea Intake and Structuring
#### FR-020 Capture Idea
Founder harus dapat memasukkan ide bebas dalam bentuk teks panjang.

#### FR-021 Generate Structured Concept
Sistem harus memproses ide mentah dan menghasilkan:
- one-line concept,
- target customer,
- value proposition,
- key assumptions,
- early risks.

#### FR-022 Manual Edit
Founder harus dapat mengedit hasil struktur ide sebelum dikunci.

#### FR-023 Lock Stage
Setelah dikonfirmasi, hasil konsep harus dikunci sebagai baseline versi 1.

#### FR-024 Versioning
Sistem harus menyimpan perubahan konsep per versi.

### 4.4 Customer Hypothesis
#### FR-030 Define Segment
Founder harus dapat memilih atau membuat customer segment target.

#### FR-031 Segment Attributes
Sistem harus menyimpan atribut customer segment seperti usia, konteks beli, budget, lokasi, dan momen konsumsi.

#### FR-032 Hypothesis Check
Sistem harus menandai jika definisi customer terlalu umum.

### 4.5 Menu Focus Engine
#### FR-040 Add Menu Candidates
Founder harus dapat menambahkan beberapa kandidat menu.

#### FR-041 Complexity Scoring
Sistem harus memberi skor kompleksitas per menu berdasarkan bahan, waktu prep, risiko spoilage, dan variasi proses.

#### FR-042 Hero SKU Recommendation
Sistem harus merekomendasikan satu hero SKU dan maksimal tiga SKU MVP.

#### FR-043 Menu Limit Enforcement
Sistem harus mencegah lebih dari tiga SKU aktif pada tahap MVP.

#### FR-044 Deferred Menu List
Sistem harus menyimpan daftar menu yang ditunda.

### 4.6 Cost and Margin Engine
#### FR-050 Ingredient Input
Founder harus dapat memasukkan daftar bahan, unit, kuantitas, dan harga.

#### FR-051 Packaging Input
Founder harus dapat memasukkan biaya kemasan.

#### FR-052 Labor and Overhead Input
Founder harus dapat memasukkan estimasi biaya tenaga kerja dan overhead.

#### FR-053 HPP Calculation
Sistem harus menghitung HPP per porsi.

#### FR-054 Selling Price Suggestion
Sistem harus memberi harga tes yang disarankan berdasarkan target margin.

#### FR-055 Margin Classification
Sistem harus mengklasifikasikan margin ke level sehat, tipis, atau berbahaya.

#### FR-056 Micro Break-even
Sistem harus menampilkan estimasi break-even minimum.

### 4.7 Mission Orchestrator
#### FR-060 Generate Missions
Sistem harus menghasilkan daftar mission berdasarkan stage venture dan risk profile.

#### FR-061 Mission Metadata
Setiap mission harus memiliki tujuan, estimasi waktu, prioritas, deadline, dan evidence requirement.

#### FR-062 Mission Acceptance
Founder harus dapat menerima atau menunda mission.

#### FR-063 Mission Completion Guard
Mission tidak dapat ditandai selesai tanpa evidence yang sesuai.

#### FR-064 Manual Mission
Founder atau mentor harus dapat membuat mission manual.

### 4.8 Evidence Management
#### FR-070 Evidence Submission
Founder harus dapat mengunggah evidence berupa gambar, teks, tautan, atau kombinasi.

#### FR-071 Evidence Metadata
Sistem harus menyimpan tipe evidence, waktu unggah, mission terkait, dan catatan user.

#### FR-072 Multiple Evidence per Mission
Satu mission harus dapat memiliki lebih dari satu evidence.

#### FR-073 Evidence Retrieval
Founder, mentor, dan admin harus dapat melihat evidence sesuai hak akses.

### 4.9 Evidence Review Engine
#### FR-080 Automated Review
Sistem harus menilai evidence dan memberikan status valid, weak, invalid, atau suspicious.

#### FR-081 Review Explanation
Sistem harus memberikan alasan evaluasi dalam bahasa sederhana.

#### FR-082 Next Action Recommendation
Sistem harus menghasilkan tindakan lanjut: continue, repeat, pivot, atau stop.

#### FR-083 Human Override
Admin atau mentor dengan izin harus dapat mengubah hasil review otomatis.

#### FR-084 Review Audit Trail
Setiap review dan override harus tercatat.

### 4.10 Scoring and Decision Engine
#### FR-090 Score Calculation
Sistem harus menghitung readiness score berdasarkan komponen terdefinisi.

#### FR-091 Score Factors
Komponen minimum skor:
- clarity score,
- focus score,
- economics score,
- execution score,
- evidence score,
- market response score.

#### FR-092 Threshold Management
Admin harus dapat mengubah threshold score dan rule gate.

#### FR-093 Final Decision
Sistem harus menghasilkan keputusan continue, repeat, pivot, atau stop saat syarat tercapai.

#### FR-094 Decision Rationale
Sistem harus menampilkan alasan keputusan secara transparan.

### 4.11 Founder Courtroom
#### FR-100 Multi-Role Review
Sistem harus dapat menghasilkan review dari minimal tiga perspektif: customer, operator, dan business reviewer.

#### FR-101 Objection List
Sistem harus menampilkan daftar keberatan utama.

#### FR-102 Assumption Attack
Sistem harus mengidentifikasi asumsi paling lemah.

### 4.12 Notifications
#### FR-110 Reminder
Sistem harus mengirim pengingat mission, deadline, dan review result.

#### FR-111 Escalation
Jika user pasif selama periode tertentu, sistem harus mengirim prompt aktivasi ulang.

### 4.13 Mentor Dashboard
#### FR-120 Mentee List
Mentor harus dapat melihat daftar founder yang dibina.

#### FR-121 Progress Summary
Mentor harus dapat melihat stage, score, mission status, dan recent evidence.

#### FR-122 Commenting
Mentor harus dapat memberi komentar pada mission dan evidence.

### 4.14 Analytics & Audit
#### FR-130 Event Tracking
Sistem harus mencatat event inti untuk analitik produk.

#### FR-131 Audit Logs
Sistem harus menyimpan audit log perubahan data kritikal.

## 5. Kebutuhan Antarmuka Eksternal
### 5.1 User Interface Requirements
- UI harus mobile-first.
- UI harus mendukung Bahasa Indonesia pada MVP.
- Terdapat dashboard utama venture.
- Terdapat mission board dengan status visual jelas.
- Terdapat halaman khusus score dan decision rationale.

### 5.2 API Requirements
- Semua resource inti harus tersedia melalui API versi terkelola, misalnya `/api/v1/...`.
- API harus menggunakan JSON.
- API harus mendukung pagination pada list besar.
- API harus menggunakan token-based authentication.

### 5.3 Storage Interface
- Media evidence harus disimpan pada object storage dengan URL aman dan terbatas.
- Metadata disimpan pada relational database.

### 5.4 AI Service Interface
Sistem AI minimal memiliki endpoint atau service contract untuk:
- concept structuring,
- menu scoring,
- evidence review,
- courtroom objections,
- recommendation generation.

## 6. Model Data Konseptual
### 6.1 Entity List
- users
- roles
- ventures
- venture_versions
- customer_segments
- menus
- menu_scores
- ingredients
- recipe_costs
- missions
- mission_templates
- evidences
- evidence_reviews
- scores
- decisions
- comments
- notifications
- audit_logs
- analytics_events

### 6.2 Entity Relationship Overview
- Satu user dapat memiliki banyak ventures.
- Satu venture memiliki banyak versions.
- Satu venture memiliki banyak menus, missions, evidences, dan scores.
- Satu mission dapat memiliki banyak evidences.
- Satu evidence memiliki satu atau lebih review records.
- Satu venture memiliki banyak decision snapshots.

### 6.3 Minimal Field Definitions
#### users
- id
- role_id
- full_name
- email
- phone
- password_hash
- status
- created_at
- updated_at

#### ventures
- id
- owner_user_id
- name
- category
- region
- stage
- current_version
- created_at
- updated_at

#### missions
- id
- venture_id
- title
- description
- mission_type
- priority
- due_at
- status
- evidence_required
- created_by
- created_at

#### evidences
- id
- venture_id
- mission_id
- uploader_user_id
- evidence_type
- storage_url
- text_content
- submitted_at

#### evidence_reviews
- id
- evidence_id
- reviewer_type
- verdict
- score
- rationale
- created_at

#### scores
- id
- venture_id
- clarity_score
- focus_score
- economics_score
- execution_score
- evidence_score
- market_response_score
- total_score
- created_at

## 7. Business Rules
### 7.1 Gate Rules
- Venture tidak boleh masuk stage `SKU_FOCUSED` sebelum konsep dan segment customer dikunci.
- Venture tidak boleh masuk stage `MISSION_ACTIVE` sebelum cost evaluation selesai.
- Mission tidak boleh berstatus completed tanpa evidence.
- Keputusan final tidak boleh dihasilkan bila komponen skor wajib belum tersedia.

### 7.2 Score Rules
- Total score harus disimpan bersama breakdown komponen.
- Threshold keputusan harus dapat diubah oleh admin.
- Review override harus memicu recalculation score.

### 7.3 Role Rules
- Founder hanya dapat mengakses venture miliknya sendiri kecuali dibagikan ke mentor.
- Mentor tidak boleh mengubah data biaya inti tanpa izin founder.
- Admin dapat melihat semua data untuk kebutuhan operasional dan QA.

## 8. Non-Functional Requirements
### 8.1 Performance
#### NFR-001
95 persen request API baca harus merespons di bawah 500 ms pada beban normal.

#### NFR-002
95 persen request tulis non-media harus merespons di bawah 800 ms pada beban normal.

#### NFR-003
Evidence review asynchronous awal harus selesai dalam maksimal 60 detik untuk 90 persen kasus.

### 8.2 Availability
#### NFR-010
Sistem harus menargetkan availability 99.5 persen per bulan untuk MVP.

### 8.3 Scalability
#### NFR-020
Sistem harus mampu menangani minimal 10.000 venture aktif tanpa redesign arsitektur total.

### 8.4 Security
#### NFR-030
Password harus di-hash dengan algoritma kuat.

#### NFR-031
Semua trafik harus melalui HTTPS.

#### NFR-032
Akses ke object storage evidence harus dibatasi dengan signed URL atau mekanisme setara.

#### NFR-033
Sistem harus menerapkan authorization per role dan per resource.

#### NFR-034
Audit log harus tersedia untuk aktivitas sensitif.

### 8.5 Privacy
#### NFR-040
Sistem harus memberi tahu user bahwa data evidence dapat diproses oleh AI.

#### NFR-041
Sistem harus menyediakan mekanisme hapus akun dan data sesuai kebijakan retention.

### 8.6 Reliability
#### NFR-050
Kalkulasi score harus idempotent dan reproducible untuk input yang sama.

#### NFR-051
Kegagalan service AI tidak boleh menyebabkan kehilangan data input user.

### 8.7 Maintainability
#### NFR-060
Kode harus dipisahkan per domain module.

#### NFR-061
Konfigurasi scoring dan mission template harus dapat diubah tanpa deploy ulang penuh jika memungkinkan.

### 8.8 Observability
#### NFR-070
Sistem harus memiliki structured logging, metrics, dan error tracing.

#### NFR-071
Setiap job asynchronous harus memiliki status processing yang dapat ditelusuri.

## 9. Kebutuhan AI/Rule Engine
### 9.1 General AI Requirements
- Model harus diarahkan oleh system prompt/domain rules yang konsisten.
- Output AI harus dinormalisasi ke schema terstruktur.
- Sistem harus menyimpan raw input dan output AI untuk debugging internal.

### 9.2 Explainability
- Setiap rekomendasi AI harus memiliki explanation field.
- Sistem harus memisahkan fakta input user dari inferensi AI.

### 9.3 Safety
- AI tidak boleh memberikan saran yang mendorong tindakan ilegal.
- AI tidak boleh menyatakan keberhasilan bisnis sebagai kepastian.
- AI tidak boleh memalsukan bukti atau menganjurkan manipulasi bukti.

### 9.4 Rule Engine Requirements
- Rule engine harus dapat menggabungkan parameter deterministic dan hasil probabilistic dari AI.
- Rule engine harus menjadi lapisan final untuk keputusan gate kritikal.

## 10. Validasi dan Error Handling
### 10.1 Input Validation
- Semua field wajib harus divalidasi di client dan server.
- Format file evidence harus divalidasi.
- Ukuran unggahan harus dibatasi.

### 10.2 Error Handling
- Sistem harus mengembalikan pesan error yang jelas dan dapat ditindaklanjuti.
- Kegagalan AI harus menghasilkan status retry atau fallback manual review.
- Jika upload media gagal, metadata draft tidak boleh hilang.

## 11. Logging, Analytics, dan Event Tracking
### 11.1 Core Product Events
- signup_completed
- venture_created
- idea_structured
- customer_defined
- menu_selected
- cost_completed
- mission_generated
- mission_started
- evidence_submitted
- evidence_reviewed
- score_updated
- final_decision_generated

### 11.2 Analytics Requirements
- Event harus menyertakan venture_id, user_id, timestamp, dan metadata kunci.
- Analytics harus mendukung funnel 14 hari.
- Sistem harus mendukung segmentasi berdasarkan kategori produk dan wilayah.

## 12. QA dan Acceptance Test
### 12.1 System Acceptance Criteria
Sistem diterima untuk MVP bila:
- Founder dapat menyelesaikan flow inti tanpa intervensi admin.
- Evidence dapat diunggah dan direview dengan stabil.
- Score dan decision dapat dihitung ulang tanpa inkonsistensi.
- Mentor dapat memonitor minimal 10 founder tanpa masalah performa berarti.

### 12.2 Test Categories
- Unit test untuk perhitungan cost dan score.
- Integration test untuk workflow mission ke evidence review.
- API test untuk auth, venture, mission, evidence, score.
- End-to-end test untuk flow 14 hari simulatif.
- Security test untuk auth dan media access.

## 13. Deployment dan Operasional
### 13.1 Environment
- local
- staging
- production

### 13.2 CI/CD Minimum
- automated tests on pull request
- migration check
- linting
- deploy to staging before production

### 13.3 Backup
- Database backup harian.
- Retention backup sesuai kebijakan operasional.

## 14. Future Extension Points
- Multi-language support.
- Category-specific rule packs.
- Region-based ingredient benchmark.
- Integration with payment/pre-order tools.
- Mentor cohort analytics.

## 15. Prioritas Implementasi
### P0
- Auth dasar.
- Venture workspace.
- Idea structuring.
- Menu focus.
- Cost engine.
- Mission orchestration.
- Evidence upload.
- Evidence review.
- Score & decision engine.

### P1
- Mentor dashboard.
- Founder Courtroom lebih kaya.
- Manual admin override tools.
- Analytics dashboard internal.

### P2
- Benchmark data eksternal.
- Recommendation personalization lanjutan.
- Lightweight collaboration.

## 16. Open Technical Questions
- Apakah evidence review dilakukan synchronous untuk teks dan asynchronous untuk gambar?
- Apakah scoring dihitung fully event-driven atau request-driven?
- Apakah MVP perlu OCR sejak awal atau cukup metadata + text explanation?
- Apakah mentor comments harus realtime?
- Seberapa granular versioning venture perlu disimpan?

## 17. Lampiran: Use Case Inti
### UC-01 Founder memulai venture
Actor: Founder  
Precondition: akun aktif  
Main flow:
1. Founder membuat venture.
2. Founder memasukkan ide.
3. Sistem menyusun konsep.
4. Founder mengonfirmasi konsep.

### UC-02 Founder menyelesaikan mission
Actor: Founder  
Precondition: venture berada di stage mission active  
Main flow:
1. Founder memilih mission.
2. Founder menjalankan aksi lapangan.
3. Founder mengunggah evidence.
4. Sistem mereview evidence.
5. Sistem mengubah status mission.

### UC-03 Sistem menghasilkan keputusan
Actor: System  
Precondition: data score minimum tersedia  
Main flow:
1. Sistem mengagregasi komponen score.
2. Sistem menerapkan threshold dan gate rules.
3. Sistem menghasilkan keputusan dan rationale.
4. Founder melihat hasil dan next actions.

