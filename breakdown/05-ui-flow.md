# UI Flow & Screen Spec — BukaDulu MVP
## Versi 1.0 | Flutter Web+Mobile

---

## 1. Navigation Tree

```
/ (Splash)
├── /login
├── /register
├── /dashboard (daftar venture) — 404 redirect jika belum login
│   ├── /venture/new
│   ├── /venture/:id/idea
│   │   └── /venture/:id/idea/result
│   ├── /venture/:id/customer
│   ├── /venture/:id/menu
│   ├── /venture/:id/cost
│   ├── /venture/:id/missions
│   │   └── /venture/:id/missions/:mission_id
│   ├── /venture/:id/evidence
│   │   └── /venture/:id/evidence/:evidence_id
│   ├── /venture/:id/score
│   ├── /venture/:id/decision
│   └── /venture/:id/courtroom
├── /mentor/dashboard (hanya role mentor)
│   └── /mentor/mentee/:venture_id
└── /settings
```

GoRouter config:
```dart
GoRouter(
  initialLocation: '/',
  redirect: (context, state) {
    final isLoggedIn = ...;
    final isOnAuthPage = state.matchedLocation.startsWith('/login') || ...;
    if (!isLoggedIn && !isOnAuthPage) return '/login';
    if (isLoggedIn && isOnAuthPage) return '/dashboard';
    return null;
  },
  routes: [
    GoRoute(path: '/', ...),
    GoRoute(path: '/login', ...),
    GoRoute(path: '/register', ...),
    ShellRoute(
      builder: (context, state, child) => AppShell(child: child),
      routes: [
        GoRoute(path: '/dashboard', ...),
        GoRoute(path: '/venture/:id', ...),
        // nested routes...
      ],
    ),
  ],
);
```

---

## 2. Screen-by-Screen Spec

### Screen 1: Login

```
┌─────────────────────────┐
│          BukaDulu        │  (logo + tagline)
│                          │
│  Email atau Nomor HP     │  [input field]
│  Password                │  [input field, obscured]
│                          │
│  [  MASUK  ]            │  (primary button, full width)
│                          │
│  Belum punya akun?       │
│  Daftar                  │  (text link)
│                          │
│  ─── atau ───            │
│                          │
│  [  Google  ]            │  (secondary — skip di MVP awal)
└─────────────────────────┘

States: loading (disabled button, spinner), error (inline message)
```

### Screen 2: Register

```
┌─────────────────────────┐
│         Daftar           │
│                          │
│  Nama Lengkap            │  [input]
│  Email                   │  [input]
│  Nomor Telepon           │  [input]
│  Password                │  [input, obscured]
│  Konfirmasi Password     │  [input, obscured]
│                          │
│  [  DAFTAR  ]           │
│                          │
│  Sudah punya akun?       │
│  Masuk                   │
└─────────────────────────┘

Validation: inline per field, form-level error summary
```

### Screen 3: Dashboard (Daftar Venture)

```
┌─────────────────────────┐
│  ☰           BukaDulu   │  (AppBar)
│                          │
│  [ + Ide Baru ]         │  (FAB atau button prominent)
│                          │
│  ┌──────────────────┐   │
│  │ Warung Nasi Aya  │   │  (venture card)
│  │ 🟢 Tahap: Misi    │   │
│  │ Skor: 72          │   │
│  │ Progress: ████░░  │   │
│  │ 3 misi selesai    │   │
│  │ [Lanjutkan]       │   │
│  └──────────────────┘   │
│                          │
│  ┌──────────────────┐   │
│  │ Es Cincau Hitz    │   │
│  │ 🟡 Tahap: Biaya   │   │
│  │ Skor: -           │   │
│  │ [Lanjutkan]       │   │
│  └──────────────────┘   │
│                          │
│  Empty state:            │
│  "Belum ada ide yang     │
│   divalidasi. Mulai      │
│   dari sini ↓"           │
│  [ Mulai Validasi ]     │
└─────────────────────────┘

Components: VentureCard (image/icon, name, stage badge, score, progress bar)
Empty state: ilustrasi + CTA "Mulai Validasi"
```

### Screen 4: Capture Ide (New Venture)

```
┌─────────────────────────┐
│  ←         Ide Baru     │
│                          │
│  Ceritakan ide bisnismu  │
│                          │
│  ┌────────────────────┐  │
│  │                     │  │
│  │  (textarea besar)   │  │
│  │                     │  │
│  │                     │  │
│  │  Min 20 karakter    │  │
│  └────────────────────┘  │
│                          │
│  Contoh: "Saya mau       │
│   jual nasi goreng       │
│   dengan topping ayam    │
│   geprek, target         │
│   kantoran..."           │
│                          │
│  Nama Usaha: [         ] │
│  Kategori:  [▼ Pilih   ] │
│  Lokasi:    [         ] │
│                          │
│  [  PROSES IDE  ]       │  (disabled jika < 20 char)
└─────────────────────────┘

Loading state: skeleton + "AI sedang menyusun konsep..."
Error state: "Gagal memproses, coba lagi" + retry button
```

### Screen 5: Hasil Struktur Ide

```
┌─────────────────────────┐
│  ←     Konsep Ide       │
│                          │
│  ✅ Ide berhasil         │
│     diproses!            │
│                          │
│  ┌─ Konsep 1 Kalimat ──┐│
│  │ Nasi Goreng AYGE    ││  (editable)
│  │ siap saji untuk     ││
│  │ kantoran di sekitar  ││
│  │ pasar, harga 15rb   ││
│  └─────────────────────┘│
│                          │
│  ┌─ Target Customer ───┐│
│  │ Karyawan kantor 25- ││  (editable)
│  │ 40th, budget makan  ││
│  │ siang 15-20rb, di   ││
│  │ area pasar          ││
│  └─────────────────────┘│
│                          │
│  ┌─ Key Assumptions ───┐│
│  │ • Kantoran butuh     ││  (editable)
│  │   makan siang cepat  ││
│  │ • Harga 15rb sesuai  ││
│  │ •...                 ││
│  └─────────────────────┘│
│                          │
│  [Edit] setiap section   │  (inline edit mode)
│                          │
│  [  KONFIRMASI IDE  ]   │  (lock stage)
└─────────────────────────┘
```

### Screen 6: Customer Segment

```
┌─────────────────────────┐
│  ←    Target Pelanggan  │
│                          │
│  Siapa yang akan beli?   │
│                          │
│  ┌─ Usia ──────────────┐│
│  │ 25 - 40 tahun       ││
│  └─────────────────────┘│
│  ┌─ Konteks Beli ──────┐│
│  │ Makan siang di       ││
│  │ kantor / dibungkus   ││
│  └─────────────────────┘│
│  ┌─ Budget ────────────┐│
│  │ Rp 15.000 - 20.000  ││
│  └─────────────────────┘│
│  ┌─ Lokasi ────────────┐│
│  │ Area perkantoran     ││
│  │ Pasar Baru           ││
│  └─────────────────────┘│
│  ┌─ Momen Konsumsi ────┐│
│  │ Senin-Jumat jam      ││
│  │ 11.30-13.00          ││
│  └─────────────────────┘│
│                          │
│  ⚠️ Deskripsi agak umum. │(warning jika terlalu general)
│  Coba tambah lebih       │
│  spesifik.               │
│                          │
│  [  KONFIRMASI  ]       │
└─────────────────────────┘
```

### Screen 7: Menu Focus

```
┌─────────────────────────┐
│  ←    Pilih Menu        │
│                          │
│  Tambah menu yang akan   │
│  kamu uji ↓              │
│                          │
│  [ + Tambah Menu ]      │
│                          │
│  ┌─────────────────────┐ │
│  │ ⭐ Nasi Goreng AYGE  │ │(hero — rekomendasi sistem)
│  │ 🟢 Kompleksitas: 3/10│ │
│  │ Bahan: 8, Prep: 15m  │ │
│  │ [Aktif] [Hapus]      │ │
│  └─────────────────────┘ │
│                          │
│  ┌─────────────────────┐ │
│  │ 🍗 Ayam Geprek       │ │
│  │ 🟡 Kompleksitas: 6/10│ │
│  │ [Aktif] [Hapus]      │ │
│  └─────────────────────┘ │
│                          │
│  ┌─────────────────────┐ │
│  │ Es Teh Manis         │ │
│  │ 🟢 Kompleksitas: 1/10│ │
│  │ [Aktif] [Hapus]      │ │
│  └─────────────────────┘ │
│                          │
│  3 SKU — maksimum        │(counter)
│                          │
│  Menu ditunda:            │
│  • Nasi Goreng Seafood   │
│  • Mie Ayam               │
│                          │
│  [  KUNCI SKU  ]        │  (disabled jika < 1 SKU active)
└─────────────────────────┘
```

### Screen 8: Cost & Margin

```
┌─────────────────────────┐
│  ←    Biaya & Margin    │
│                          │
│  ── Nasi Goreng AYGE ── │  (per SKU tab/accordion)
│                          │
│  Bahan:                  │
│  ┌────────┬────┬────┬──┐ │
│  │ Bahan  | Qty| Sat| Hrg│ │
│  ├────────┼────┼────┼──┤ │
│  │ Nasi   │ 200| gr |200│ │
│  │ Telur  │ 1  | pcs|200│ │
│  │ Bumbu  │ 1  | porsi|5│ │
│  │ Minyak  │ 10 | ml |2 │ │
│  │ ...     │    |    |   │ │
│  │ [+ Baris]            │ │
│  └──────────────────────┘ │
│                          │
│  Kemasan:  [Rp 1.500 ]   │
│  Tenaga:   [Rp 1.000 ]   │
│  Overhead: [Rp   500 ]   │
│                          │
│  ─────────────────────    │
│  HPP:          Rp 3.850   │
│  Harga Tes:    Rp 15.000  │
│  Margin:       🟢 74%     │
│  Break-even:   25 porsi/  │
│                bulan      │
│  ─────────────────────    │
│                          │
│  [  KONFIRMASI BIAYA  ] │
└─────────────────────────┘

Edge case: Margin merah → warning card:
"⚠️ Margin di bawah 20%. Risiko tinggi. Disarankan pivot harga atau efisiensi bahan."
```

### Screen 9: Mission Board

```
┌─────────────────────────┐
│  ←    Misi Harian       │
│                          │
│  📅 Hari ke-3 dari 14   │(progress indicator)
│  ████░░░░░░░░ 22%       │
│                          │
│  Today's Priority:       │
│  ┌────────────────────┐  │
│  │ 🔴 [High] Polling   │  │
│  │ Tanya 10 orang di   │  │
│  │ sekitar kantormu    │  │
│  │ ⏱ 30 menit          │  │
│  │ 📅 Deadline: Besok  │  │
│  │ [Terima] [Tunda]    │  │
│  └────────────────────┘  │
│                          │
│  ┌────────────────────┐  │
│  │ 🟡 [Med] Wawancara  │  │
│  │ Tanya 3 orang...    │  │
│  │ [Mulai]             │  │
│  └────────────────────┘  │
│                          │
│  ── Active ──            │
│  ☑️ Sampling produk      │(sudah diterima, progress)
│    (upload evidence)     │→ button "Upload Evidence"
│                          │
│  ── Completed ──         │
│  ✅ Observasi lokasi     │
│    Evidence: valid ✓     │
│                          │
│  [ + Misi Manual ]      │(opsi mentor/founder)
│                          │
│  Empty state (all done): │
│  "Semua misi selesai!    │
│   Sistem mengecek        │
│   evidence..."           │
└─────────────────────────┘
```

### Screen 10: Upload Evidence

```
┌─────────────────────────┐
│  ←   Upload Evidence    │
│                          │
│  Misi: Polling           │
│  "Tanya 10 orang..."    │
│                          │
│  📷 [Ambil Foto]        │(camera/gallery)
│  📝 [Catatan]           │(text input)
│  🔗 [Link/URL]          │(link input)
│                          │
│  ┌────────────────────┐  │
│  │                    │  │
│  │  [Preview gambar]  │  │
│  │                    │  │
│  └────────────────────┘  │
│                          │
│  Catatan: [             ]│
│  "Isi detail evidence..."│
│                          │
│  [  KIRIM EVIDENCE  ]   │
│                          │
│  ── Already Uploaded ──  │
│  ✅ Screenshot chat      │→ Lihat
│     Valid ✅             │
│  ✅ Catatan survey       │→ Lihat
│     ⏳ Review...         │
└─────────────────────────┘

Upload progress: linear bar selama upload
Error: "Gagal upload. Coba lagi." + retry button
```

### Screen 11: Score Dashboard

```
┌─────────────────────────┐
│  ←   Skor Kesiapan      │
│                          │
│      ╭─────────╮        │
│      │   72    │        │(large circle progress)
│      │ /100    │        │
│      ╰─────────╯        │
│     READY TO DECIDE     │(stage badge)
│                          │
│  ── Breakdown ──         │
│                          │
│  Clarity        ████████ 82│
│  Focus          ████████ 80│
│  Economics      ████████ 90│
│  Execution      ██████░ 60│
│  Evidence       █████░░ 50│
│  Market Resp.   ████░░░ 45│
│                          │
│  ⚠️ Kerang: Evidence     │
│  skor masih rendah.      │
│  Upload bukti tambahan.  │
│                          │
│  [  HASILKAN KEPUTUSAN  ]│(disabled jika belum siap)
│                          │
│  ── Riwayat Skor ──      │(line chart mini)
│  📈 0 ▸ 45 ▸ 60 ▸ 72    │
└─────────────────────────┘
```

### Screen 12: Final Decision

```
┌─────────────────────────┐
│  ←    Keputusan Akhir   │
│                          │
│      ╭─────────╮        │
│      │  ✅     │        │(big icon based on outcome)
│      │ CONTINUE│        │
│      ╰─────────╯        │
│                          │
│  Ide kamu layak untuk    │
│  lanjut ke fase uji      │
│  jual!                   │
│                          │
│  ── Reasoning ──         │
│  • Margin 74% (sehat)   │
│  • 4 evidence valid      │
│  • Responden positif     │
│  • Customer jelas        │
│                          │
│  ── Next Steps ──        │
│  1. Produksi 10 porsi   │
│  2. Pre-order ke 5 org  │
│  3. Catat feedback       │
│                          │
│  [  Lihat Detail Skor  ]│
│  [  Mulai Venture Baru ]│
│                          │
│  ────────────────────    │
│  "Butuh mentor untuk     │
│   fase selanjutnya?"     │
│  [Cari Mentor]           │
└─────────────────────────┘

Variant — STOP screen:
┌─────────────────────────┐
│      ╭─────────╮        │
│      │  🛑     │        │
│      │  STOP   │        │
│      ╰─────────╯        │
│                          │
│  Berdasarkan evidence    │
│  yang ada, ide ini       │
│  tidak layak lanjut      │
│  dalam bentuk sekarang.  │
│                          │
│  Ini bukan kegagalan —   │
│  ini bukti kamu berani   │
│  tahu sebelum rugi. 🔥  │
│                          │
│  Pelajaran:              │
│  • Margin terlalu tipis  │
│  • Tidak ada permintaan  │
│                          │
│  [Coba Ide Baru]         │
│  [Unduh Laporan]         │
└─────────────────────────┘
```

### Screen 13: Founder Courtroom (MVP-lite)

```
┌─────────────────────────┐
│  ←   Founder Courtroom  │
│                          │
│  Sistem akan menguji     │
│  idemu dari 3 sudut      │
│  pandang:                │
│                          │
│  👤 Calon Pembeli        │
│  👨‍🍳 Operator Dapur       │
│  📋 Reviewer Bisnis      │
│                          │
│  [  MULAI REVIEW  ]     │
│                          │
│  ── Hasil Review ──      │(setelah AI selesai)
│                          │
│  🗣 Calon Pembeli:       │
│  "Harga 15rb masih       │
│   mahal buat saya..."    │
│                          │
│  👨‍🍳 Operator Dapur:      │
│  "Bahan gampang, tapi    │
│   butuh 2 orang..."      │
│                          │
│  📋 Reviewer Bisnis:     │
│  "Margin bagus, tapi     │
│   target terlalu luas."  │
│                          │
│  ── Weakest Assumptions   │
│  • Harga 15rb kompetitif │
│  • Kantoran mau nunggu   │
│                          │
│  ── Required Actions ──  │
│  • Validasi harga dgn    │
│    polling               │
│                          │
│  [Tutup]                 │
└─────────────────────────┘
```

---

## 3. Reusable Components

| Component | Location | Behavior |
|---|---|---|
| **StageBadge** | Dashboard, header | Warna sesuai stage: draft(grey), idea(blue), customer(green), sku(orange), cost(yellow), mission(purple), evidence(pink), decision(red→green) |
| **ScoreCircle** | Score, Decision | Animasi circular progress. Warna: <40(red), 40-69(yellow), 70+(green) |
| **MissionCard** | Mission Board | Drag priority, swipe to accept/tunda, expandable |
| **EvidencePreview** | Upload, Review | Thumbnail grid, click to fullscreen, status overlay |
| **CostRow** | Cost Input | Auto-calculate total, swipe to delete |
| **ConfirmBottomSheet** | Any lock stage | "Yakin? Setelah ini tidak bisa edit." + reason |

---

## 4. Responsive Breakpoints

| Screen | Layout | Target |
|---|---|---|
| < 640px | Single column, full width | Mobile |
| 640-1024px | Single column, max-width 480px centered | Tablet |
| > 1024px | Two panel (sidebar + content) | Desktop web |

---

## 5. Empty State & Error State Strategy

Setiap screen harus punya:

| State | Behavior |
|---|---|
| **Loading** | Skeleton shimmer, bukan spinner saja |
| **Empty** | Ilustrasi + copywriting ramah + CTA jelas |
| **Error** | Pesan jelas + retry button + option contact support |
| **Offline** | Banner "Koneksi terputus" + cache data yang bisa |
