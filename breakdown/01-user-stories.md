# User Stories — BukaDulu MVP
## Versi 1.0 | Target: Flutter Web+Mobile × Go Backend

> Setiap story mengacu ke FR-xxx dari SRS. Format: *"Sebagai [role], saya ingin [action] agar [benefit]"*

---

## Epic 1: Authentication & Account (P0)

### US-001 Registrasi
| | |
|---|---|
| **Story** | Sebagai calon founder, saya ingin mendaftar dengan email atau nomor telepon agar bisa mulai menggunakan aplikasi. |
| **FR** | FR-001, FR-004 |
| **Acceptance Criteria** | |
| | 1. Form registrasi menerima email atau nomor telepon + password |
| | 2. Password minimal 8 karakter |
| | 3. Validasi format email dan nomor telepon |
| | 4. Role otomatis "founder" saat registrasi |
| | 5. Email/telepon unik — duplicate ditolak dengan pesan jelas |
| | 6. Account langsung aktif (tanpa verifikasi di MVP) |
| | 7. Error handling untuk koneksi terputus |

### US-002 Login & Logout
| | |
|---|---|
| **Story** | Sebagai founder, saya ingin login dan logout dengan aman agar data saya terlindungi. |
| **FR** | FR-002, FR-003 |
| **Acceptance Criteria** | |
| | 1. Login dengan email/telepon + password |
| | 2. Token-based session (JWT) |
| | 3. Logout menghapus session token |
| | 4. Session expired setelah 24 jam |
| | 5. Auto-redirect ke login jika token invalid/expired |

---

## Epic 2: Venture Workspace (P0)

### US-010 Membuat Venture
| | |
|---|---|
| **Story** | Sebagai founder, saya ingin membuat workspace bisnis baru agar bisa mulai memvalidasi ide saya. |
| **FR** | FR-010, FR-012 |
| **Acceptance Criteria** | |
| | 1. Input: nama usaha, kategori F&B, lokasi |
| | 2. Venture dibuat dengan stage `DRAFT` |
| | 3. User bisa memiliki lebih dari satu venture |
| | 4. Venture langsung bisa diedit |
| | 5. Stage venture tercatat di database |

### US-011 Mengelola Venture
| | |
|---|---|
| **Story** | Sebagai founder, saya ingin mengedit data venture saya agar informasinya tetap akurat. |
| **FR** | FR-011 |
| **Acceptance Criteria** | |
| | 1. Edit nama, kategori, lokasi, status |
| | 2. Hanya owner yang bisa edit |
| | 3. Perubahan tersimpan dengan timestamp |

---

## Epic 3: Idea Intake & Structuring (P0)

### US-020 Capture Ide Mentah
| | |
|---|---|
| **Story** | Sebagai founder, saya ingin memasukkan ide bisnis saya dalam bahasa bebas agar aplikasi membantu menyusunnya. |
| **FR** | FR-020 |
| **Acceptance Criteria** | |
| | 1. Textarea bebas minimal 20 karakter |
| | 2. Placeholder dengan contoh: *"Saya ingin jual nasi goreng homemade, target kantoran sekitar pasar..."* |
| | 3. Tersedia tombol "Proses Ide" |
| | 4. Loading state selama AI processing |

### US-021 Generate Structured Concept
| | |
|---|---|
| **Story** | Sebagai founder, saya ingin ide mentah saya diubah jadi konsep bisnis terstruktur agar saya punya arah jelas. |
| **FR** | FR-021 |
| **Acceptance Criteria** | |
| | 1. AI menghasilkan: one-line concept, target customer, value proposition, key assumptions, early risks |
| | 2. Output ditampilkan dalam format structured card per elemen |
| | 3. Setiap elemen bisa diedit manual (US-022) |
| | 4. Loading state wajib selama AI processing |
| | 5. Jika AI gagal, tampilkan pesan error + retry button |
| | 6. Raw input dan output AI disimpan untuk debugging (SRS 9.1) |

### US-022 Edit & Konfirmasi Konsep
| | |
|---|---|
| **Story** | Sebagai founder, saya ingin mengedit hasil struktur AI sebelum mengunci agar sesuai dengan visi saya. |
| **FR** | FR-022, FR-023, FR-024 |
| **Acceptance Criteria** | |
| | 1. Semua field hasil AI bisa diedit bebas |
| | 2. Tombol "Konfirmasi" mengunci stage ke `IDEA_DEFINED` |
| | 3. Setelah dikunci, versi tersimpan sebagai baseline |
| | 4. Versi lama tetap bisa dilihat |
| | 5. Tidak bisa kembali ke tahap sebelumnya tanpa reset |

---

## Epic 4: Customer Hypothesis (P0)

### US-030 Define Customer Segment
| | |
|---|---|
| **Story** | Sebagai founder, saya ingin menentukan siapa target pelanggan saya agar produk sesuai dengan kebutuhan mereka. |
| **FR** | FR-030, FR-031, FR-032 |
| **Acceptance Criteria** | |
| | 1. Input: usia, konteks beli, budget, lokasi, momen konsumsi |
| | 2. Sistem menawarkan segment suggestion dari AI |
| | 3. Sistem menandai jika deskripsi terlalu umum (misal: "semua orang") |
| | 4. Konfirmasi mengunci customer segment |
| | 5. Stage maju ke `CUSTOMER_DEFINED` |

---

## Epic 5: Menu Focus Engine (P0)

### US-040 Tambah Kandidat Menu
| | |
|---|---|
| **Story** | Sebagai founder, saya ingin menambahkan beberapa ide menu agar saya bisa pilih yang paling layak diuji. |
| **FR** | FR-040 |
| **Acceptance Criteria** | |
| | 1. Form: nama menu, deskripsi singkat, perkiraan bahan utama |
| | 2. Bisa tambah banyak menu |
| | 3. Daftar menu ditampilkan sebagai cards |
| | 4. Bisa hapus menu kandidat |

### US-041 Skor & Rekomendasi Hero SKU
| | |
|---|---|
| **Story** | Sebagai founder, saya ingin sistem menilai kompleksitas tiap menu dan merekomendasikan hero SKU. |
| **FR** | FR-041, FR-042 |
| **Acceptance Criteria** | |
| | 1. Sistem memberi skor kompleksitas per menu (bahan, prep time, spoilage risk, process variation) |
| | 2. Sistem merekomendasikan 1 hero SKU |
| | 3. Sistem membatasi maksimal 3 SKU aktif |
| | 4. Sistem menunjukkan alasan rekomendasi |

### US-042 Pilih & Kunci SKU
| | |
|---|---|
| **Story** | Sebagai founder, saya ingin memilih SKU final yang akan saya uji di MVP. |
| **FR** | FR-043, FR-044 |
| **Acceptance Criteria** | |
| | 1. Wajib pilih minimal 1 SKU |
| | 2. Maksimal 3 SKU — sistem blokir jika lebih |
| | 3. Menu yang tidak dipilih masuk daftar "ditunda" |
| | 4. Konfirmasi mengunci pilihan SKU |
| | 5. Stage maju ke `SKU_FOCUSED` |

---

## Epic 6: Cost & Margin Engine (P0)

### US-050 Input Biaya Bahan
| | |
|---|---|
| **Story** | Sebagai founder, saya ingin memasukkan bahan dan biaya agar tahu HPP produk saya. |
| **FR** | FR-050, FR-051, FR-052 |
| **Acceptance Criteria** | |
| | 1. Per SKU bisa input daftar bahan: nama, unit, kuantitas, harga |
| | 2. Input biaya kemasan per porsi |
| | 3. Input estimasi tenaga kerja dan overhead |
| | 4. Validasi: harga dan kuantitas harus angka positif |

### US-051 Hitung HPP & Margin
| | |
|---|---|
| **Story** | Sebagai founder, saya ingin sistem menghitung HPP, margin, dan break-even agar saya tahu usaha ini realistis atau tidak. |
| **FR** | FR-053, FR-054, FR-055, FR-056 |
| **Acceptance Criteria** | |
| | 1. Sistem menghitung HPP per porsi otomatis |
| | 2. Menampilkan harga jual minimum + saran harga tes |
| | 3. Klasifikasi margin: 🟢 Sehat (>40%) / 🟡 Tipis (20-40%) / 🔴 Berbahaya (<20%) |
| | 4. Estimasi break-even minimum (unit per hari/bulan) |
| | 5. Perhitungan realtime — berubah saat input berubah |
| | 6. Tombol "Konfirmasi" mengunci data biaya |
| | 7. Stage maju ke `COST_EVALUATED` |

---

## Epic 7: Mission Orchestrator (P0)

### US-060 Generate Misi Harian
| | |
|---|---|
| **Story** | Sebagai founder, saya ingin sistem membuat misi lapangan harian agar saya tahu persis apa yang harus dilakukan. |
| **FR** | FR-060, FR-061 |
| **Acceptance Criteria** | |
| | 1. Sistem generate minimal 1 misi per hari berdasarkan stage dan risk profile |
| | 2. Tiap misi punya: tujuan, estimasi waktu, prioritas, deadline, definisi bukti |
| | 3. Misi ditampilkan di mission board |
| | 4. Bisa regenerate misi (dengan batasan) |

### US-061 Terima & Jalankan Misi
| | |
|---|---|
| **Story** | Sebagai founder, saya ingin menerima misi dan menandainya sedang dikerjakan. |
| **FR** | FR-062, FR-064 |
| **Acceptance Criteria** | |
| | 1. Tombol "Terima Misi" untuk mulai |
| | 2. Bisa tunda misi (masuk daftar pending) |
| | 3. Founder atau mentor bisa buat misi manual |

### US-062 Selesaikan Misi dengan Bukti
| | |
|---|---|
| **Story** | Sebagai founder, saya tidak bisa menandai misi selesai tanpa mengunggah bukti. |
| **FR** | FR-063 |
| **Acceptance Criteria** | |
| | 1. Tombol "Selesai" disabled sampai ada evidence yang diupload |
| | 2. Flow: upload evidence → review → mark complete |
| | 3. Stage maju ke `MISSION_ACTIVE` |

---

## Epic 8: Evidence Management (P0)

### US-070 Upload Evidence
| | |
|---|---|
| **Story** | Sebagai founder, saya ingin mengunggah bukti berupa foto, teks, atau tautan agar misi saya dinilai. |
| **FR** | FR-070, FR-071, FR-072 |
| **Acceptance Criteria** | |
| | 1. Upload: image (jpeg/png, max 5MB), text, atau link |
| | 2. Multiple evidence per mission |
| | 3. Auto-attach ke mission aktif |
| | 4. Progress upload (bar/percentage) |
| | 5. Jika upload gagal, data form tidak hilang |
| | 6. Tipe evidence dan timestamp tersimpan |

### US-071 Lihat Evidence
| | |
|---|---|
| **Story** | Sebagai founder, saya ingin melihat semua evidence yang pernah saya upload. |
| **FR** | FR-073 |
| **Acceptance Criteria** | |
| | 1. Galeri evidence per venture |
| | 2. Filter by mission, tipe, status review |
| | 3. Preview image inline |

---

## Epic 9: Evidence Review Engine (P0)

### US-080 Auto Review Evidence
| | |
|---|---|
| **Story** | Sebagai founder, saya ingin evidence saya langsung dinilai oleh sistem agar saya tahu apakah bukti saya cukup kuat. |
| **FR** | FR-080, FR-081, FR-082 |
| **Acceptance Criteria** | |
| | 1. AI menilai evidence dengan status: ✅ Valid / ⚠️ Weak / ❌ Invalid / 🤔 Suspicious |
| | 2. Setiap verdict disertai alasan jelas dalam bahasa Indonesia |
| | 3. Sistem merekomendasikan next action: continue / repeat / pivot |
| | 4. Review async — tidak ngeblock user |
| | 5. 90% review selesai dalam 60 detik (NFR-003) |
| | 6. Status review (pending/processing/done) |

### US-081 Human Override
| | |
|---|---|
| **Story** | Sebagai admin, saya ingin bisa override hasil review otomatis jika ada ketidaksesuaian. |
| **FR** | FR-083, FR-084 |
| **Acceptance Criteria** | |
| | 1. Admin bisa mengubah verdict |
| | 2. Setiap override tercatat di audit trail |
| | 3. Override memicu recalculation score |

---

## Epic 10: Scoring & Decision Engine (P0)

### US-090 Lihat Score
| | |
|---|---|
| **Story** | Sebagai founder, saya ingin melihat readiness score saya agar tahu seberapa siap saya uji jual. |
| **FR** | FR-090, FR-091 |
| **Acceptance Criteria** | |
| | 1. Total score ditampilkan dengan 6 komponen: clarity, focus, economics, execution, evidence, market response |
| | 2. Score diupdate tiap ada bukti baru (FR-092) |
| | 3. Breakdown komponen bisa dilihat per item |
| | 4. Visual progress bar per komponen |

### US-091 Dapat Keputusan Akhir
| | |
|---|---|
| **Story** | Sebagai founder, saya ingin sistem memberi keputusan jelas di akhir sprint: lanjut, ulang, pivot, atau stop. |
| **FR** | FR-093, FR-094 |
| **Acceptance Criteria** | |
| | 1. Keputusan hanya keluar jika semua komponen score wajib terisi |
| | 2. Output: continue / repeat / pivot / stop |
| | 3. Setiap keputusan disertai rationale transparan |
| | 4. Menampilkan data pendukung yang melandasi keputusan |
| | 5. Stage maju ke CONTINUE / REPEAT / PIVOT / STOP |

---

## Epic 11: Founder Courtroom (MVP-lite)

### US-100 Jalankan Courtroom Review
| | |
|---|---|
| **Story** | Sebagai founder, saya ingin ide saya di-review dari 3 sudut pandang: pembeli, operator dapur, dan reviewer bisnis. |
| **FR** | FR-100, FR-101, FR-102 |
| **Acceptance Criteria** | |
| | 1. Sistem generate review dari 3 perspektif (mult-turn AI prompt) |
| | 2. Output: objection list + weakest assumptions |
| | 3. Tersedia action items berdasarkan objections |
| | 4. Courtroom bisa dijalankan kapan saja di tahap post-idea |

---

## Epic 12: Notifications (P0)

### US-110 Dapat Notifikasi
| | |
|---|---|
| **Story** | Sebagai founder, saya ingin mendapat pengingat misi dan notifikasi review agar tidak kehilangan momentum. |
| **FR** | FR-110, FR-111 |
| **Acceptance Criteria** | |
| | 1. Notifikasi untuk: mission deadline, review selesai, score update |
| | 2. Jika user pasif >48 jam, kirim reactivation prompt |
| | 3. Notifikasi in-app (push di MVP lanjutan) |

---

## Epic 13: Mentor Dashboard (P1)

### US-120 Lihat Mentee Progress
| | |
|---|---|
| **Story** | Sebagai mentor, saya ingin melihat progress semua founder yang saya bina. |
| **FR** | FR-120, FR-121 |
| **Acceptance Criteria** | |
| | 1. List mentee dengan stage, score, mission status |
| | 2. Filter by stage, status, recency |
| | 3. Bisa drill-down ke detail venture |

### US-121 Komentar Mentor
| | |
|---|---|
| **Story** | Sebagai mentor, saya ingin memberi komentar pada mission dan evidence. |
| **FR** | FR-122 |
| **Acceptance Criteria** | |
| | 1. Komentar per mission atau per evidence |
| | 2. Founder mendapat notifikasi komentar baru |
| | 3. Komentar tidak bisa diedit (cukup soft delete) |

---

## Prioritasi Sprint

| Sprint | Stories |
|---|---|
| **Sprint 1** | US-001, US-002, US-010, US-011, US-020, US-021, US-022 |
| **Sprint 2** | US-030, US-040, US-041, US-042, US-050, US-051 |
| **Sprint 3** | US-060, US-061, US-062, US-070, US-071, US-080, US-081 |
| **Sprint 4** | US-090, US-091, US-100, US-110, US-120, US-121 |
