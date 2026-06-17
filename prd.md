# Product Requirements Document (PRD)
## Produk: BukaDulu
## Versi: 1.0
## Tanggal: 2026-06-02
## Status: Draft for MVP Approval

## 1. Ringkasan Produk
BukaDulu adalah aplikasi eksekusi bisnis F&B tahap 0-ke-1 yang membantu calon founder berhenti dari fase "kepikiran" dan bergerak ke fase "bukti pasar pertama". Produk ini tidak diposisikan sebagai generator business plan atau restaurant management system, melainkan sebagai execution system yang memecah ide menjadi eksperimen lapangan kecil, mengukur bukti, lalu memutuskan apakah ide dilanjutkan, dipivot, atau dimatikan.

Masalah yang ingin diselesaikan sangat spesifik: banyak orang ingin memulai usaha F&B, tetapi berhenti pada fase bingung, takut salah, terlalu banyak ide, dan tidak punya sistem yang memaksa mereka bergerak. Di pasar sudah ada aplikasi validator ide berbasis AI dan software manajemen restoran, tetapi celah terbesar berada pada fase sebelum bisnis benar-benar lahir: fase validasi lapangan yang murah, cepat, dan disiplin.[cite:7][cite:11][cite:20]

## 2. Latar Belakang
### 2.1 Problem Statement
Calon pebisnis F&B pemula sering mengalami kombinasi masalah berikut:
- Tidak tahu langkah pertama yang konkret.
- Terjebak overthinking, riset berlebihan, dan gonta-ganti ide.
- Langsung berpikir besar: branding, logo, menu banyak, sewa tempat, padahal belum ada bukti demand.
- Tidak mampu menilai apakah ide mereka layak secara margin, operasional, dan minat pasar.
- Tidak punya mentor atau sistem yang memaksa akuntabilitas harian.

### 2.2 Opportunity
Aplikasi validator ide yang ada cenderung menghasilkan laporan, skor, dan analisis AI, sedangkan software restoran cenderung berfokus pada POS, inventory, scheduling, dan operasi bisnis yang sudah berjalan.[cite:7][cite:11][cite:20] BukaDulu mengambil posisi di tengah: mengubah niat menjadi aksi, lalu aksi menjadi bukti.

## 3. Visi Produk
Menjadi sistem paling efektif untuk membantu calon pebisnis F&B mendapatkan penjualan valid pertama dengan risiko serendah mungkin.

## 4. Misi Produk
Dalam 14 hari, membantu pengguna:
1. Menyempitkan ide menjadi 1 model jualan yang jelas.
2. Menentukan 1-3 SKU awal yang layak diuji.
3. Menghitung realitas dasar usaha: HPP, harga, margin, modal awal, dan kapasitas.
4. Menjalankan eksperimen lapangan kecil.
5. Mengumpulkan bukti permintaan pasar.
6. Mengambil keputusan: lanjut, pivot, atau stop.

## 5. Positioning
### 5.1 Positioning Statement
Untuk calon pebisnis F&B yang ingin mulai tapi mandek di kepala, BukaDulu adalah aplikasi eksekusi validasi awal yang memaksa pengguna mengumpulkan bukti pasar nyata, bukan sekadar membuat rencana. Tidak seperti idea validator biasa atau software restoran, BukaDulu fokus pada hari-hari paling krusial sebelum bisnis pertama lahir.[cite:7][cite:11][cite:20]

### 5.2 Value Proposition
- Mengurangi kebingungan awal.
- Mengurangi risiko salah langkah dan overspending.
- Memaksa disiplin eksekusi.
- Membantu pengguna membedakan semangat palsu dari bukti nyata.
- Membuat keputusan bisnis awal menjadi objektif.

## 6. Tujuan Produk
### 6.1 Business Goals
- Membangun kategori baru: pre-launch execution system untuk F&B.
- Mencapai product-market fit awal pada segmen calon pebisnis F&B rumahan.
- Membuka monetisasi via subscription sprint, template operasional, dan layanan lanjutan.

### 6.2 User Goals
- Tahu langkah pertama yang harus dilakukan.
- Tahu produk apa yang harus dijual dulu.
- Tahu apakah ide layak diteruskan.
- Mendapat penjualan pertama atau bukti penolakan pasar dengan cepat.

### 6.3 Success Criteria
- Pengguna menyelesaikan sprint 14 hari.
- Pengguna berhasil menjalankan minimal 1 eksperimen lapangan.
- Pengguna mengunggah minimal 1 bukti validasi nyata.
- Pengguna mengambil keputusan bisnis yang jelas pada akhir sprint.

## 7. Segmen Pengguna
### 7.1 Primary Persona: Aspiring F&B Founder
Karakteristik:
- Usia 20-40 tahun.
- Ingin memulai usaha makanan/minuman rumahan.
- Modal terbatas.
- Belum pernah menjalankan bisnis secara serius.
- Sangat semangat, tetapi rendah struktur.

Pain points:
- Bingung mulai dari mana.
- Ingin menu banyak sejak awal.
- Takut rugi dan malu gagal.
- Tidak paham margin dan operasional.

### 7.2 Secondary Persona: Side Hustler
Karakteristik:
- Sudah bekerja full time.
- Ingin menguji ide bisnis tanpa berhenti kerja.
- Waktu terbatas.

Pain points:
- Sulit menentukan eksperimen yang realistis.
- Tidak punya sistem monitoring harian.

### 7.3 Tertiary Persona: Mentor/Coach
Karakteristik:
- Pendamping UMKM, mentor bisnis, atau inkubator.
- Membutuhkan kerangka validasi yang rapi dan repeatable.

Pain points:
- Sulit memonitor banyak mentee.
- Sulit membedakan mentee yang progress vs sekadar wacana.

## 8. Jobs To Be Done
- Ketika memiliki ide bisnis F&B tapi bingung mulai, pengguna ingin aplikasi yang langsung mengubah ide menjadi langkah konkret supaya mereka bisa bergerak hari ini.
- Ketika ragu apakah ide mereka layak, pengguna ingin melihat realitas margin, biaya, dan respons pasar supaya tidak membuang waktu.
- Ketika cenderung menunda, pengguna ingin sistem yang memaksa akuntabilitas supaya ide tidak berhenti di kepala.

## 9. Prinsip Produk
- Evidence over opinions.
- Progress harus dikunci oleh bukti, bukan niat.
- Fokus pada tindakan murah, cepat, dan nyata.
- Satu hero product lebih baik daripada banyak menu.
- Aplikasi harus berani mengatakan ide buruk adalah ide buruk.

## 10. Ruang Lingkup MVP
### 10.1 Outcome MVP
Dalam 14 hari, user bisa bergerak dari ide mentah ke salah satu outcome berikut:
- Penjualan pertama tercapai.
- Terdapat bukti permintaan kuat.
- Terdapat bukti bahwa ide tidak layak, sehingga dihentikan lebih awal.

### 10.2 Fitur Utama MVP
#### A. Idea Intake & Idea Compressor
Fungsi:
- Menangkap ide mentah user dalam bahasa bebas.
- Mengubahnya menjadi konsep bisnis satu kalimat.
- Mengidentifikasi target customer, use case, dan konteks jualan.

Output:
- Problem statement.
- Customer hypothesis.
- Value proposition.
- Daftar asumsi utama.

#### B. Menu Focus Engine
Fungsi:
- Memaksa user memilih maksimum 1-3 SKU awal.
- Menilai kompleksitas operasional tiap SKU.
- Menyarankan hero product.

Output:
- Rekomendasi hero SKU.
- Daftar menu yang ditunda.
- Alasan pemangkasan menu.

#### C. Reality Ledger
Fungsi:
- Menghitung HPP kasar.
- Mengestimasi harga jual, margin kotor, modal awal minimum, kapasitas produksi, dan titik impas mikro.
- Menandai risiko margin tipis atau beban operasional berlebih.

Output:
- Cost sheet sederhana.
- Harga tes.
- Batas minimum order untuk break even.
- Alarm risiko.

#### D. Field Missions
Fungsi:
- Menghasilkan tugas harian konkret berbasis tahap user.
- Tugas dapat berupa polling, pre-order, sampling, titip jual, wawancara calon pembeli, atau uji harga.
- Tiap misi punya definisi selesai yang jelas.

Output:
- Daily mission list.
- Deadline.
- Evidence requirement.

#### E. Evidence Upload & Review
Fungsi:
- User mengunggah bukti: foto produk, screenshot chat, daftar pesanan, hasil polling, catatan feedback.
- Sistem menilai apakah bukti valid, lemah, ambigu, atau gagal.

Output:
- Evidence status.
- Feedback tegas.
- Keputusan lanjut, ulang, pivot, atau stop.

#### F. Launch Readiness Score
Fungsi:
- Menghasilkan skor berbasis eksekusi, bukan opini.
- Skor berasal dari kombinasi: kejelasan customer, fokus menu, margin, jumlah bukti, kualitas bukti, dan respons pasar.

Output:
- Skor kesiapan uji jual.
- Risk flags.
- Rekomendasi keputusan.

#### G. Founder Courtroom (MVP-lite)
Fungsi:
- Menjalankan review adversarial sederhana dengan tiga peran: calon pembeli, operator dapur, dan reviewer bisnis.
- Menguji kelemahan asumsi user.

Output:
- Objection list.
- Weak assumption list.
- Required action before next step.

### 10.3 Fitur di Luar MVP
- POS.
- Inventory management lanjutan.
- Payroll.
- Multi-store management.
- Marketplace supplier.
- Integrasi pembayaran.
- Brand kit/logomaker.
- Akuntansi penuh.

## 11. User Journey
### 11.1 End-to-End Flow
1. User onboarding.
2. User memasukkan ide mentah.
3. Sistem merumuskan konsep bisnis awal.
4. User memilih customer segment awal.
5. Sistem memangkas menu menjadi 1-3 SKU.
6. User mengisi data bahan dan biaya kasar.
7. Sistem menghitung margin dan risiko awal.
8. Sistem membuat misi lapangan.
9. User menjalankan misi dan unggah bukti.
10. Sistem menilai bukti.
11. Jika lolos, user lanjut ke sprint berikutnya; jika gagal, user diminta pivot atau stop.
12. Di akhir sprint, user menerima keputusan akhir.

### 11.2 Moment of Truth
Momen kunci produk bukan saat user selesai onboarding, melainkan saat user pertama kali menyadari bahwa mereka harus mengunggah bukti nyata untuk bisa lanjut. Inilah inti diferensiasi produk.

## 12. User Stories
### 12.1 Founder
- Sebagai calon pebisnis F&B, saya ingin memasukkan ide mentah dalam bahasa biasa agar aplikasi membantu saya menyusunnya menjadi konsep yang jelas.
- Sebagai pengguna pemula, saya ingin tahu SKU mana yang harus diuji lebih dulu agar saya tidak membuang tenaga pada menu yang salah.
- Sebagai orang dengan modal terbatas, saya ingin tahu apakah margin saya masuk akal agar saya tidak menjual produk rugi.
- Sebagai orang yang mudah menunda, saya ingin diberi tugas harian yang konkret agar saya benar-benar bergerak.
- Sebagai pengguna, saya ingin aplikasi menilai bukti saya secara tegas agar saya tahu apakah harus lanjut atau berhenti.

### 12.2 Mentor
- Sebagai mentor, saya ingin melihat progres mentee berdasarkan bukti agar pendampingan lebih objektif.

## 13. Kebutuhan UX
### 13.1 UX Principles
- Satu layar, satu keputusan utama.
- Bahasa sederhana dan keras kepala, bukan jargon bisnis generik.
- Progress hanya terlihat jika ada bukti.
- Visual progress harus terasa seperti permainan taktis, bukan form isian panjang.

### 13.2 Key UX Decisions
- Onboarding maksimal 7 menit.
- Tiap tahap punya checklist jelas.
- Tugas harian harus dapat diselesaikan dalam 15-60 menit.
- Feedback harus tegas: "cukup", "kurang kuat", "ulang", "pivot", atau "stop".
- Dokumen dan angka ditampilkan sesederhana mungkin.

## 14. Fitur Detail dan Acceptance Criteria
### 14.1 Idea Compressor
Acceptance criteria:
- User dapat mengirim deskripsi ide bebas minimal 20 karakter.
- Sistem menghasilkan konsep satu kalimat, customer segment, masalah utama, dan daftar asumsi.
- User dapat mengedit hasil sebelum mengunci tahap.

### 14.2 Menu Focus Engine
Acceptance criteria:
- User dapat memasukkan beberapa ide menu.
- Sistem memberi skor kompleksitas dan rekomendasi hero SKU.
- User tidak dapat memilih lebih dari 3 SKU untuk MVP.

### 14.3 Reality Ledger
Acceptance criteria:
- User dapat mengisi bahan, porsi, biaya kemasan, ongkir opsional, dan biaya tenaga kerja estimasi.
- Sistem menghitung HPP, harga jual minimum, dan margin kotor.
- Sistem menandai status merah jika margin di bawah ambang yang ditetapkan.

### 14.4 Field Missions
Acceptance criteria:
- Sistem menghasilkan minimal 1 misi prioritas setiap hari.
- Tiap misi memiliki tujuan, estimasi durasi, deadline, dan definisi bukti selesai.
- User dapat menandai misi selesai hanya jika bukti diunggah.

### 14.5 Evidence Review
Acceptance criteria:
- User dapat mengunggah gambar, teks, dan link bukti.
- Sistem mengklasifikasikan bukti menjadi valid/weak/invalid.
- Sistem memberikan feedback dan next action.

### 14.6 Launch Readiness Score
Acceptance criteria:
- Skor dihitung ulang setiap ada bukti baru.
- Skor menampilkan komponen penyusunnya.
- Sistem memberikan keputusan akhir pada hari ke-14 atau saat user memenuhi ambang tertentu.

## 15. Metric Framework
### 15.1 North Star Metric
Persentase pengguna yang mencapai "first validated market evidence" dalam 14 hari.

### 15.2 Primary Metrics
- Activation rate: user yang menyelesaikan idea compression dalam 24 jam.
- Mission completion rate.
- Evidence submission rate.
- Evidence validity rate.
- 14-day decision completion rate.
- First sale rate.

### 15.3 Secondary Metrics
- Average time to first mission.
- Average days to first valid evidence.
- Menu reduction rate.
- Pivot rate.
- Stop-early rate.
- Retention day 7 dan day 14.

### 15.4 Guardrail Metrics
- Task abandonment rate.
- False positive recommendation rate.
- User frustration score.
- Evidence review disagreement rate.

## 16. Decision Engine Logic
### 16.1 Stage Gates
Setiap tahap memiliki gate:
- Gate 1: Ide cukup jelas.
- Gate 2: SKU sudah fokus.
- Gate 3: Margin tidak fatal.
- Gate 4: Misi lapangan selesai.
- Gate 5: Bukti memadai.
- Gate 6: Keputusan akhir.

### 16.2 Outcome Types
- Continue: evidence cukup kuat.
- Repeat: eksperimen perlu diulang.
- Pivot: hipotesis perlu diubah.
- Stop: ide tidak layak pada bentuk sekarang.

## 17. Risiko Produk
### 17.1 Product Risks
- User merasa aplikasi terlalu keras dan berhenti.
- AI memberi saran terlalu generik.
- Perhitungan HPP terlalu kasar dan menyesatkan.
- Bukti pengguna mudah dimanipulasi.
- Field missions terlalu berat untuk pemula.

### 17.2 Mitigasi
- Gunakan tone tegas tapi tetap membantu.
- Simpan transparansi alasan rekomendasi.
- Tampilkan disclaimer bahwa angka awal adalah estimasi.
- Gunakan bukti ganda bila perlu.
- Sesuaikan misi dengan level komitmen user.

## 18. Asumsi dan Dependensi
### 18.1 Asumsi
- User memiliki smartphone.
- User bersedia melakukan eksperimen lapangan kecil.
- User dapat menyediakan data bahan baku kasar.
- User nyaman mengunggah bukti sederhana.

### 18.2 Dependensi
- LLM untuk ide decomposition dan feedback.
- Rule engine untuk gate dan scoring.
- OCR/image classification opsional untuk evaluasi bukti.
- Notification system untuk task reminder.

## 19. Strategi Go-To-Market Awal
Target awal:
- Calon bisnis rumahan.
- Komunitas UMKM pemula.
- Konten edukasi startup kecil dan realistis.
- Mentor bisnis atau inkubator mikro sebagai channel distribusi.

Akuisisi awal:
- Konten pendek: "Jangan buka usaha dulu sebelum lolos 14 hari ini".
- Landing page dengan audit ide gratis.
- Program cohort kecil dengan mentor.

## 20. Pricing Hypothesis
- Free tier: idea compression + 3 misi awal.
- Paid sprint: 14-day validation program.
- Pro/mentor tier: dashboard pantau beberapa founder.

## 21. Roadmap
### Phase 1: MVP
- Idea intake.
- Menu focus.
- Reality ledger.
- Field missions.
- Evidence review.
- Launch score.

### Phase 2
- Mentor dashboard.
- Template SOP produksi kecil.
- Pre-order link generator.
- Market feedback clustering.

### Phase 3
- Supplier recommendation.
- Cost benchmarking by region.
- Cohort mode.
- Integration with lightweight bookkeeping.

## 22. Open Questions
- Seberapa agresif tone aplikasi harus dibuat?
- Apakah evidence review otomatis cukup akurat untuk fase awal?
- Apakah scoring cukup dipercaya user tanpa mentor manusia?
- Ambang margin minimum harus generik atau disesuaikan per kategori?
- Apakah MVP perlu fokus satu kategori dulu, misalnya minuman, rice bowl, atau snack?

## 23. Keputusan Produk yang Direkomendasikan
- Fokus ke F&B rumahan dulu, jangan semua bisnis.
- Jangan tambahkan fitur branding atau business plan generator pada MVP.
- Jadikan bukti sebagai mata uang progress.
- Jadikan keputusan "stop" sebagai fitur, bukan kegagalan.
- Optimalkan 14 hari pertama sebagai proposisi nilai utama.

## 24. Lampiran: Daftar Screen Inti MVP
- Welcome / onboarding.
- Idea capture.
- Concept summary.
- Customer hypothesis.
- Menu selection.
- Cost input.
- Margin dashboard.
- Daily mission board.
- Evidence upload.
- Evidence review.
- Readiness score.
- Final decision screen.

