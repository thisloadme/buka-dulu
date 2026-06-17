# BukaDulu — Design System

## Brand Identity

BukaDulu adalah execution system untuk calon pebisnis F&B. Visual identity mencerminkan karakter produk: **tegas, fungsional, tidak membuang waktu**. Palet warna terinspirasi dari maskot belalah .sigil dan api semangat eksekusi.

### Color Palette

```
Orange (Brand Primary)     #ea580c   — aksi, CTA, urgensi, hero
Orange Hover               #c2410c   — interaksi aktif
Amber (Accent)             #f59e0b   — highlight, badge, warning ringan
Cream (Bg Ringan)          #fff7ed   — kartu, section divider subtle

Dark (Heading/Title)       #1c1917   — teks utama, judul
Body                       #57534e   — teks paragraf, caption
Label                      #292524   — label form, teks kecil
Border                     #e7e5e4   — batas kartu, divider
Border Light               #f5f5f4   — batas subtle
Bg Warm                    #fafaf9   — background section, landing

White                      #ffffff   — kartu, kontainer form
Black                      #000000   — hanya untuk overlay / modal backdrop
```

```
Usage constraints:
  - Orange hanya untuk: CTA primer, ikon aktif, stage badge aktif, highlight data
  - Dark (#1c1917) untuk heading dan teks utama
  - Body (#57534e) untuk teks biasa, caption, subtitle
  - Jangan gunakan gradient di mana pun. Semua background solid.
  - Jangan gunakan emoji di UI. Ikon dengan SVG.
```

### Typography

**Font Family:** Inter (sans-serif) — system-ui, -apple-system, Segoe UI, Roboto sebagai fallback.

```
Display Hero     — clamp(2.25rem, 5.5vw, 3.75rem)  weight 300, letter-spacing -0.05em
Display Large    — clamp(2rem, 4vw, 3rem)          weight 300, letter-spacing -0.04em
Body Large       — 1.125rem                         weight 300, line-height 1.55
Body Text        — 1rem                              weight 300, line-height 1.6
Caption          — 0.875rem                          weight 400, color var(--body)
Micro            — 0.75rem                           weight 500, uppercase, letter-spacing 0.08em
                  color var(--brand-orange)
```

### Spacing Grid

Base unit: 4px. All spacing derived from multiples of 4.

```
Section padding:      80px–100px vertical (landing), 24px (auth page)
Card padding:         28px–32px
Container max-width:  1120px
Auth container:       420px max-width
Button padding:       10px 22px (default) / 14px 32px (lg)
Gap grid:             20px–24px (card grid), 32px (step cards)
Border radius:        6px (buttons), 8px–10px (cards), 12px (auth container), 100px (badge/pill)
```

### Shadows

```
Card shadow:       rgba(80,50,20,0.12) 0px 50px 80px -40px,
                   rgba(0,0,0,0.06) 0px 20px 40px -20px
Auth box shadow:   rgba(80,50,20,0.08) 0px 20px 40px -25px
Hover elevation:   rgba(80,50,20,0.10) 0px 20px 40px -20px
```

---

## Design Principles

### 1. Satu Layar, Satu Keputusan Utama
Setiap halaman memiliki satu primary action yang dominan. Tidak boleh ada dua tombol dengan visual weight yang sama dalam satu layar.

### 2. Progress Dikunci oleh Bukti
UI tidak boleh menampilkan progress yang tidak didukung data. Progress bar hanya bergerak saat ada evidence. Area "selesai" hanya muncul jika gate terlewati.

### 3. Bahasa Tegas, Bukan Jargon
Tidak ada kata "optimalisasi", "sinergi", "leverage". Gunakan bahasa sehari-hari: "Coba tanya 10 orang", "Hitung modal dulu", "Bukti belum cukup, ulangi".

### 4. Fokus pada Aksi, Bukan Bacaan
Hindari paragraf panjang di dalam alur. Setiap teks non-heading maksimal 2 baris. Detail ada di tooltip atau expandable section.

### 5. Hitam-Putih dengan Aksen Orange
Tidak ada gradient. Tidak ada emoji. Semua ikon berupa SVG. Warna orange hanya untuk elemen yang membutuhkan perhatian segera.

---

## Component Design

### Button System

| Variant | Background | Text | Border | Use Case |
|---------|-----------|------|--------|----------|
| Primary | `var(--brand-orange)` | White | None | CTA utama, konfirmasi, lanjut |
| Ghost | Transparent | `var(--brand-orange)` | `#fed7aa` | Aksi sekunder, edit, batal |
| White | White | `var(--brand-orange)` | None | CTA di atas bg gelap (section pricing/final-cta) |
| Outline Light | Transparent | White | `rgba(255,255,255,0.25)` | Aksi sekunder di atas bg gelap |
| Disabled | `var(--border)` | `#a8a29e` | None | Tombol yang belum bisa diakses |

Semua button memiliki:
- `border-radius: 6px` (default) atau `8px` (btn-lg)
- `transition: all 0.2s`
- Icon SVG 20x20 inline, gap 8px dari teks
- `font-family: inherit`

### Card System

| Card Type | Background | Border | Use Case |
|-----------|-----------|--------|----------|
| Default | White | `var(--border)` | Kartu fitur, venture card, pricing |
| Cream | `var(--brand-cream)` | `#fed7aa` | Tahap selesai, step preview |
| Dark | `rgba(255,255,255,0.04)` | `rgba(255,255,255,0.08)` | Kartu di atas bg gelap |
| Featured | `rgba(255,255,255,0.07)` | `var(--brand-orange)` | Pricing card terpilih |

Hover state: `translateY(-2px)` dengan border-color `#fed7aa`.

### Badge & Tag

| Variant | Background | Text Color | Size |
|---------|-----------|-----------|------|
| Stage badge (default) | `rgba(234,88,12,0.08)` | `var(--brand-orange)` | 0.8125rem, pill |
| Feature tag | `rgba(234,88,12,0.08)` + border `rgba(234,88,12,0.12)` | `var(--brand-orange)` | 0.6875rem, uppercase |
| Hero badge | Orange solid | White | Pill, absolute positioned |

### Form Elements

```
Input:
  - border: 1px solid var(--border)
  - border-radius: 8px
  - padding: 12px 16px
  - focus: border-color var(--brand-orange)
  - placeholder: #a8a29e

Label:
  - font-size: 0.875rem
  - font-weight: 500
  - color: var(--label)
  - margin-bottom: 6px

Error state:
  - bg: #fef2f2
  - border: #fecaca
  - text: #dc2626
  - border-radius: 8px
  - padding: 10px 14px
```

### Navigation

```
Navbar:
  - position: sticky, top: 0, z-index: 100
  - background: rgba(255,255,255,0.85) + backdrop-filter blur(14px)
  - border-bottom: 1px solid var(--border-light)
  - height: 64px
  
Nav links:
  - font-size: 0.875rem
  - color: var(--label)
  - hover: color var(--brand-orange)

Mobile menu:
  - hamburger toggle (3 span bars)
  - dropdown: full-width, bg rgba(255,255,255,0.98)
```

---

## Screen Architecture

### Screen Tree

```
/ (Splash/Landing)
├── /login
├── /register
├── /download              — post-auth landing (coming soon untuk APK)
├── /dashboard             — daftar venture (auth required)
│   ├── /venture/new
│   ├── /venture/:id
│   │   ├── /idea          — capture + result
│   │   ├── /customer      — segment definisi
│   │   ├── /menu          — fokus SKU
│   │   ├── /cost          — biaya & margin
│   │   ├── /missions      — board misi
│   │   │   └── /missions/:mid
│   │   ├── /evidence      — upload & list
│   │   │   └── /evidence/:eid
│   │   ├── /score         — readiness score
│   │   ├── /decision      — keputusan akhir
│   │   └── /courtroom     — founder courtroom
├── /mentor/dashboard      — (role mentor)
│   └── /mentor/mentee/:vid
└── /settings
```

### Layout Structure

```
Public pages (Login, Register):
  [full viewport center]
    auth-container (420px)
      logo + header
      form
      footer link

Dashboard (authenticated):
  [Sticky Navbar] — logo + nav + user menu
  [Main Content]
    container (1120px max)
      content area
  [Footer] — minimal, dark bg

Venture flow:
  [App Shell] — back button + title + stage indicator
  [Content] — single-column, max 720px
  [Bottom CTA] — primary action button, sticky on mobile
```

---

## Screen-by-Screen Design Spec

### 1. Landing Page

```
┌──────────────────────────────────────────────┐
│  Navbar: [BukaDulu logo]  Cara Kerja  Fitur  │
│          Harga  FAQ  Masuk  [Coba Gratis]    │
├──────────────────────────────────────────────┤
│                                              │
│  Hero Section (bg-warm)                      │
│  ┌────────────────────────────────────┐      │
│  │  badge "Riset dan overthinking..."  │      │
│  │  [Punya ide jualan makanan?         │      │
│  │   Buktikan dulu dalam 14 hari.]     │      │
│  │  Subtitle: "Bukan business plan..." │      │
│  │  [Cek Ide Gratis] [Lihat Cara Kerja]│      │
│  │  14 hari    1-3 SKU    100% bukti   │      │
│  └────────────────────────────────────┘      │
│                                              │
│  Problem Section (dark #1c1917)              │
│  6 problem cards with icons                   │
│                                              │
│  How It Works (bg-warm)                      │
│  5 step circles (orange solid)               │
│                                              │
│  Features (white)                            │
│  6 feature cards with tags + icons            │
│                                              │
│  Testimonial (bg-warm)                        │
│  Quote card with avatar                       │
│                                              │
│  Pricing (dark #1c1917)                      │
│  2 pricing cards, featured dengan orange     │
│  border                                       │
│                                              │
│  Final CTA (orange solid)                    │
│  "Berhenti memikirkan. Mulai buktikan."      │
│  [Cek Ide Saya Gratis] (white button)        │
│                                              │
│  Footer (dark)                               │
└──────────────────────────────────────────────┘
```

### 2. Auth Pages (Login / Register)

```
┌──────────────────────────────┐
│         BukaDulu              │  — logo (40x40) + brand name
│                              │
│  [Masuk] atau [Daftar]       │  — heading
│  Subtitle                    │
│                              │
│  Email atau Nomor HP         │
│  ┌────────────────────────┐  │
│  │                        │  │  — input, focus orange border
│  └────────────────────────┘  │
│                              │
│  Password                    │
│  ┌────────────────────────┐  │
│  │                        │  │  — obscured
│  └────────────────────────┘  │
│                              │
│  ⚠️ Error message inline    │  — merah bg merah muda
│                              │
│  ┌────────────────────────┐  │
│  │       MASUK             │  │  — orange solid, full width
│  └────────────────────────┘  │
│                              │
│  Belum punya akun? Daftar   │  — link orange
└──────────────────────────────┘
```

### 3. Dashboard — Venture List

```
┌──────────────────────────────────────┐
│  ☰ BukaDulu                    Logout│
├──────────────────────────────────────┤
│                                      │
│  [+ Ide Baru]                        │  — FAB / prominent button
│                                      │
│  ┌──────────────────────────────┐    │
│  │  Warung Nasi AYGE            │    │
│  │  🟢 Tahap: Misi              │    │  — venture card
│  │  Skor: 72  ████░░░░  3/5     │    │
│  │  [Lanjutkan]                 │    │
│  └──────────────────────────────┘    │
│                                      │
│  ┌──────────────────────────────┐    │
│  │  Es Cincau Hitz              │    │
│  │  🟡 Tahap: Biaya             │    │
│  │  Skor: -                     │    │
│  │  [Lanjutkan]                 │    │
│  └──────────────────────────────┘    │
│                                      │
│  Empty state:                        │
│  "Belum ada ide yang divalidasi."    │
│  [Mulai Validasi]                    │
└──────────────────────────────────────┘
```

### 4. Venture Idea — Capture & Result

**Capture screen:**
```
┌──────────────────────────────────────┐
│  ← Ide Baru                          │
├──────────────────────────────────────┤
│                                      │
│  Ceritakan ide bisnismu              │
│                                      │
│  ┌────────────────────────────────┐  │
│  │                                │  │
│  │  Textarea besar (min 20 char)  │  │
│  │                                │  │
│  └────────────────────────────────┘  │
│                                      │
│  Contoh: "Saya mau jual nasi         │
│  goreng homemade..."                 │
│                                      │
│  Nama Usaha: [______________]        │
│  Kategori:  [▼ Makanan Berat]        │
│  Lokasi:    [______________]         │
│                                      │
│  ┌────────────────────────────────┐  │
│  │        PROSES IDE              │  │
│  └────────────────────────────────┘  │
└──────────────────────────────────────┘
```

**Result screen — structured concept cards:**
```
┌──────────────────────────────────────┐
│  ← Konsep Ide                        │
├──────────────────────────────────────┤
│  ✅ Ide berhasil diproses!            │
│                                      │
│  ┌─ Konsep 1 Kalimat ──────────────┐ │
│  │  Nasi Goreng AYGE siap saji...  │ │  — editable
│  └─────────────────────────────────┘ │
│                                      │
│  ┌─ Target Customer ───────────────┐ │
│  │  Karyawan kantor 25-40th...      │ │  — editable
│  └─────────────────────────────────┘ │
│                                      │
│  ┌─ Key Assumptions ───────────────┐ │
│  │  • Kantoran butuh makan siang   │ │  — editable
│  │  • Harga 15rb sesuai            │ │
│  └─────────────────────────────────┘ │
│                                      │
│  [KONFIRMASI IDE] — orange solid     │
└──────────────────────────────────────┘

Loading state:
  Skeleton card with pulse animation
  "AI sedang menyusun konsep..."
  
Error state:
  "Gagal memproses. Coba lagi."
  [Coba Lagi] — ghost button
```

### 5. Customer Segment

```
┌──────────────────────────────────────┐
│  ← Target Pelanggan                  │
├──────────────────────────────────────┤
│                                      │
│  Siapa yang akan beli?               │
│                                      │
│  Usia:          [25 - 40 tahun]      │
│  Konteks Beli:  [Makan siang...]     │
│  Budget:        [Rp 15.000-20.000]  │
│  Lokasi:        [Area Pasar Baru]    │
│  Momen:         [Senin-Jumat...]     │
│                                      │
│  ⚠️ Deskripsi agak umum.             │  — warning amber
│     Coba tambah lebih spesifik.      │
│                                      │
│  ┌────────────────────────────────┐  │
│  │        KONFIRMASI              │  │
│  └────────────────────────────────┘  │
└──────────────────────────────────────┘
```

### 6. Menu Focus

```
┌──────────────────────────────────────┐
│  ← Pilih Menu                        │
├──────────────────────────────────────┤
│                                      │
│  Tambah menu yang akan kamu uji ↓    │
│                                      │
│  [+ Tambah Menu]                     │
│                                      │
│  ┌──────────────────────────────┐    │
│  │ ⭐ Nasi Goreng AYGE [HERO]   │    │  — rekomendasi sistem
│  │ 🟢 Kompleksitas: 3/10        │    │
│  │ Bahan: 8  Prep: 15m          │    │
│  │ [Aktif] [Hapus]              │    │
│  └──────────────────────────────┘    │
│                                      │
│  ┌──────────────────────────────┐    │
│  │ 🍗 Ayam Geprek               │    │
│  │ 🟡 Kompleksitas: 6/10        │    │
│  │ [Aktif] [Hapus]              │    │
│  └──────────────────────────────┘    │
│                                      │
│  3 SKU — maksimum                    │  — counter
│                                      │
│  Menu ditunda:                       │
│  • Nasi Goreng Seafood               │
│  • Mie Ayam                          │
│                                      │
│  ┌────────────────────────────────┐  │
│  │        KUNCI SKU               │  │
│  └────────────────────────────────┘  │
└──────────────────────────────────────┘
```

### 7. Cost & Margin

```
┌──────────────────────────────────────┐
│  ← Biaya & Margin                    │
├──────────────────────────────────────┤
│                                      │
│  ── Nasi Goreng AYGE ──              │  — tab/accordion per SKU
│                                      │
│  Bahan:                              │
│  ┌──────────┬──────┬─────┬──────┐   │
│  │ Nasi     │ 200  │ gr  │ 200  │   │  — inline editable row
│  │ Telur    │ 1    │ pcs │ 2000 │   │
│  │ Bumbu    │ 1    │ ptr │ 500  │   │
│  │ [+ Baris]                      │   │
│  └──────────┴──────┴─────┴──────┘   │
│                                      │
│  Kemasan:  [Rp 1.500]               │
│  Tenaga:   [Rp 1.000]               │
│  Overhead: [Rp 500]                 │
│                                      │
│  ────────────────────────────        │
│  HPP:          Rp 3.850              │  — highlighted data
│  Harga Tes:    Rp 15.000             │
│  Margin:       🟢 74% (Sehat)        │  — 🟢🟡🔴 status
│  Break-even:   25 porsi/bulan        │
│  ────────────────────────────        │
│                                      │
│  ⚠️ Margin di bawah 20%.            │  — warning jika merah
│     Risiko tinggi.                   │
│                                      │
│  ┌────────────────────────────────┐  │
│  │      KONFIRMASI BIAYA          │  │
│  └────────────────────────────────┘  │
└──────────────────────────────────────┘
```

**Margin status indicator:**
```
🟢 Sehat       > 40%   — green: #22c55e
🟡 Tipis        20-40% — amber: #f59e0b
🔴 Berbahaya   < 20%   — red: #ef4444
```

### 8. Mission Board

```
┌──────────────────────────────────────┐
│  ← Misi Harian                       │
├──────────────────────────────────────┤
│                                      │
│  📅 Hari ke-3 dari 14                │
│  ████░░░░░░░░  22%                   │  — 14-day progress bar
│                                      │
│  Today's Priority:                   │
│  ┌──────────────────────────────┐    │
│  │ 🔴 [High] Polling            │    │  — priority badge
│  │ Tanya 10 orang di sekitar    │    │
│  │ kantormu. ⏱ 30 menit         │    │
│  │ 📅 Deadline: Besok            │    │
│  │ [Terima] [Tunda]             │    │
│  └──────────────────────────────┘    │
│                                      │
│  ── Sedang Dikerjakan ──             │
│  ☑️ Sampling produk                  │
│    (upload evidence) [Upload]        │
│                                      │
│  ── Selesai ──                       │
│  ✅ Observasi lokasi  Valid ✓        │
│                                      │
│  [+ Misi Manual]                     │
│                                      │
│  Semua misi selesai!                 │  — empty state
│  "Sistem mengecek evidence..."        │
└──────────────────────────────────────┘
```

### 9. Evidence Upload

```
┌──────────────────────────────────────┐
│  ← Upload Evidence                   │
├──────────────────────────────────────┤
│                                      │
│  Misi: Polling                       │
│  "Tanya 10 orang..."                 │
│                                      │
│  📷 [Ambil Foto]                     │  — 3 upload methods
│  📝 [Catatan]                        │
│  🔗 [Link/URL]                       │
│                                      │
│  ┌──────────────────────────────┐    │
│  │                              │    │  — image preview
│  │     [Preview gambar]         │    │
│  │                              │    │
│  └──────────────────────────────┘    │
│                                      │
│  Catatan: [____________________]     │
│                                      │
│  ┌────────────────────────────────┐  │
│  │       KIRIM EVIDENCE           │  │
│  └────────────────────────────────┘  │
│                                      │
│  ── Sudah Diupload ──                │
│  ✅ Screenshot chat  Valid ✓ [Lihat] │
│  ✅ Catatan survey   ⏳ Review...    │
└──────────────────────────────────────┘
```

### 10. Score Dashboard

```
┌──────────────────────────────────────┐
│  ← Skor Kesiapan                     │
├──────────────────────────────────────┤
│                                      │
│          ╭─────────╮                 │
│          │   72    │                 │  — large circle gauge
│          │  /100   │                 │
│          ╰─────────╯                 │
│      READY TO DECIDE                 │  — stage badge
│                                      │
│  ── Komponen Skor ──                 │
│                                      │
│  Clarity        ████████░░  82      │  — progress bar per
│  Focus          ████████░░  80      │    komponen
│  Economics      █████████░  90      │
│  Execution      ██████░░░░  60      │
│  Evidence       █████░░░░░  50      │
│  Market Resp.   ████░░░░░░  45      │
│                                      │
│  ⚠️ Kelemahan: Evidence skor        │  — weakness warning
│     masih rendah. Upload bukti       │
│     tambahan.                        │
│                                      │
│  ┌────────────────────────────────┐  │
│  │     HASILKAN KEPUTUSAN         │  │  — disabled jika belum
│  └────────────────────────────────┘  │    semua komponen siap
│                                      │
│  ── Riwayat Skor ──                  │
│  📈 0 → 45 → 60 → 72                │  — line chart mini
└──────────────────────────────────────┘
```

### 11. Final Decision

**CONTINUE outcome:**
```
┌──────────────────────────────────────┐
│  ← Keputusan Akhir                   │
├──────────────────────────────────────┤
│                                      │
│          ╭─────────╮                 │
│          │    ✅   │                 │  — large icon (SVG)
│          │ CONTINUE│                 │
│          ╰─────────╯                 │
│                                      │
│  Ide kamu layak untuk lanjut         │
│  ke fase uji jual!                   │
│                                      │
│  ── Alasan ──                        │
│  • Margin 74% (sehat)               │
│  • 4 evidence valid                  │
│  • Responden positif                 │
│  • Customer jelas                    │
│                                      │
│  ── Langkah Selanjutnya ──           │
│  1. Produksi 10 porsi                │
│  2. Pre-order ke 5 orang             │
│  3. Catat feedback                   │
│                                      │
│  [Lihat Detail Skor]                 │
│  [Mulai Venture Baru]                │
│                                      │
│  "Butuh mentor?" [Cari Mentor]       │
└──────────────────────────────────────┘
```

**STOP outcome:**
```
┌──────────────────────────────────────┐
│          ╭─────────╮                 │
│          │    🛑   │                 │
│          │   STOP  │                 │
│          ╰─────────╯                 │
│                                      │
│  Berdasarkan evidence yang ada,      │
│  ide ini tidak layak lanjut dalam    │
│  bentuk sekarang.                    │
│                                      │
│  Ini bukan kegagalan — ini bukti     │
│  kamu berani tahu sebelum rugi.      │
│                                      │
│  Pelajaran:                          │
│  • Margin terlalu tipis              │
│  • Tidak ada permintaan pasar        │
│                                      │
│  [Coba Ide Baru]                     │
│  [Unduh Laporan]                     │
└──────────────────────────────────────┘
```

### 12. Founder Courtroom

```
┌──────────────────────────────────────┐
│  ← Founder Courtroom                 │
├──────────────────────────────────────┤
│                                      │
│  Sistem akan menguji idemu dari      │
│  3 sudut pandang:                    │
│                                      │
│  👤 Calon Pembeli                    │  — 3 perspective cards
│  🔪 Operator Dapur                   │
│  📋 Reviewer Bisnis                  │
│                                      │
│  ┌──────────────────────────────┐    │
│  │ [Mulai Courtroom Review]     │    │
│  └──────────────────────────────┘    │
│                                      │
│  Result (setelah AI process):        │
│                                      │
│  ┌─ Calon Pembeli ───────────────┐   │
│  │ ❓ "Harganya murah, tapi apa   │   │  — objection cards
│  │    saya yakin kenyang?"        │   │
│  └───────────────────────────────┘   │
│                                      │
│  ┌─ Operator Dapur ──────────────┐   │
│  │ ❓ "Ayam geprek butuh minyak   │   │
│  │    banyak, margin tipis"      │   │
│  └───────────────────────────────┘   │
│                                      │
│  ┌─ Reviewer Bisnis ─────────────┐   │
│  │ ❓ "Target hanya kantoran,     │   │
│  │    weekend gimana?"           │   │
│  └───────────────────────────────┘   │
│                                      │
│  Asumsi terlemah:                    │
│  • Harga 15rb cukup untuk margin     │
│  • Kantoran mau beli tiap hari       │
│                                      │
│  Action items: [Validasi asumsi...]  │
└──────────────────────────────────────┘
```

---

## Venture State Machine — Visual Representation

Setiap venture menampilkan stage badge dan progress indicator yang mencerminkan state machine:

```
                    DRAFT
                      │ capture & confirm idea
                      ▼
              ┌───────────────┐
              │  IDEA_DEFINED │ ◀── (reset)
              └───────┬───────┘
                      │ define customer
                      ▼
           ┌───────────────────┐
           │ CUSTOMER_DEFINED  │
           └───────┬───────────┘
                    │ select & lock SKU
                    ▼
             ┌──────────────┐
             │ SKU_FOCUSED  │
             └──────┬───────┘
                    │ confirm cost
                    ▼
          ┌──────────────────┐
          │ COST_EVALUATED   │
          └──────┬───────────┘
                  │ accept first mission
                  ▼
           ┌──────────────┐
      ┌───│ MISSION_ACTIVE│◀──────────────┐
      │   └──────┬───────┘                │
      │          │ upload evidence        │ more missions
      │          ▼                        │
      │  ┌──────────────────┐             │
      │  │EVIDENCE_SUBMITTED│             │
      │  └──────┬───────────┘             │
      │         │ AI review complete      │
      │         ▼                        │
      │  ┌──────────────────┐             │
      │  │EVIDENCE_REVIEWED │─────────────┘
      │  └──────┬───────────┘
      │         │ all missions done
      │         ▼
      │  ┌──────────────────┐
      │  │ READY_TO_DECIDE  │
      │  └──────┬───────────┘
      │         │ generate decision
      │         ▼
      │  ┌─────────┐ ┌──────┐ ┌─────┐ ┌────┐
      │  │CONTINUE │ │REPEAT│ │PIVOT│ │STOP│
      │  └─────────┘ └──────┘ └─────┘ └────┘
      │       │          │       │       │
      └───────┴──────────┘───────┴───────┘
                 End states — read-only
```

**Stage badge color mapping:**
```
DRAFT              gray    — #78716c
IDEA_DEFINED       orange  — #ea580c
CUSTOMER_DEFINED   orange  — #ea580c
SKU_FOCUSED        orange  — #ea580c
COST_EVALUATED     orange  — #ea580c
MISSION_ACTIVE     blue    — #3b82f6
EVIDENCE_SUBMITTED amber   — #f59e0b
EVIDENCE_REVIEWED  green   — #22c55e
READY_TO_DECIDE    purple  — #8b5cf6
CONTINUE           green   — #22c55e
REPEAT             amber   — #f59e0b
PIVOT              amber   — #f59e0b
STOP               red     — #ef4444
```

---

## Data Visualization Components

### Score Gauge (Circle Progress)
```
Diameter: 160px
Stroke width: 8px
Track color: #e7e5e4
Fill color: #ea580c (or dynamic based on score tier)
Text: large (2.5rem) bold, "/100" small caption below
```

### Score Breakdown Bar
```
Label (left) | Bar (center) | Value (right)
Height: 24px per bar
Bar bg: #e7e5e4
Bar fill: #ea580c, rounded 4px
Animasi: width transition 0.3s ease
```

### 14-Day Progress Bar
```
Full width container
Height: 8px, rounded 4px
Track: var(--border-light)
Fill: var(--brand-orange)
Label: "Hari ke-X dari 14" di atas
```

### Mission Priority Badge
```
High:   🔴 #ef4444 — bg rgba(239,68,68,0.1)
Medium: 🟡 #f59e0b — bg rgba(245,158,11,0.1)
Low:    🟢 #22c55e — bg rgba(34,197,94,0.1)
```

---

## Responsive Breakpoints

| Breakpoint | Target | Layout |
|-----------|--------|--------|
| < 768px | Mobile | Single column, hamburger nav, stacked cards |
| 768px+ | Tablet/Desktop | Multi-column grid, horizontal nav, sidebar optional |
| Container max | 1120px | Desktop content width |

Mobile adjustments:
- Auth container: full width (padding 24px)
- 3-column grid → 1 column (hero stats, app preview, pricing, features)
- Nav links: hidden behind hamburger
- Venture cards: stacked full width
- Tables (cost): scroll horizontal atau stacked rows

---

## Empty States

| Screen | Illustration | Message | CTA |
|--------|-------------|---------|-----|
| Dashboard (no venture) | SVG ilustrasi ide kosong | "Belum ada ide yang divalidasi. Mulai dari sini." | [Mulai Validasi] |
| Mission board (all done) | SVG checklist | "Semua misi selesai! Sistem mengecek evidence..." | — |
| Evidence (none yet) | SVG upload icon | "Belum ada bukti yang diupload untuk misi ini." | [Upload Evidence] |
| Mentor dashboard (no mentee) | SVG user icon | "Belum ada mentee yang ditugaskan." | — |

---

## Loading States

| Pattern | Implementation | Usage |
|---------|---------------|-------|
| Skeleton card | Gray (#e7e5e4) rounded 8px dengan pulse animation | Card lists, venture detail |
| Button spinner | Inline loading spinner, button disabled + dim 0.6 | Form submit, AI process |
| Full page skeleton | 3-4 skeleton cards stacked | Dashboard, score page |
| Inline shimmer | Text-width gray bar with gradient shimmer | AI result loading |

---

## Error States

| Error | Display | Action |
|-------|---------|--------|
| Network failure | Inline red box: "Koneksi terputus. Coba lagi." | [Coba Lagi] button |
| Validation | Red text below field: "Password minimal 8 karakter" | Inline, per field |
| Auth failed | Inline red box: "Email atau password salah" | Form tetap terisi |
| AI failed | Inline orange box: "Gagal memproses. Coba lagi." | [Coba Lagi] button |
| 404 | Full page: "Halaman tidak ditemukan" | [Kembali ke Dashboard] |
| 500 | Full page: "Terjadi kesalahan. Tim kami sudah diberi tahu." | [Coba Lagi] |

---

## Motion & Transition

| Element | Animation | Duration | Easing |
|---------|-----------|----------|--------|
| Page enter | fadeIn + translateY(20px) | 0.6s | ease |
| Card hover | translateY(-2px) | 0.3s | ease |
| Button hover | translateY(-1px) | 0.2s | ease |
| Skor gauge | width from 0 to target | 0.5s | ease-out |
| Form focus | border-color transition | 0.2s | ease |
| Modal overlay | bg opacity 0→0.5 | 0.2s | ease |

---

## Iconography

Semua ikon menggunakan SVG inline (bukan emoji, bukan font icon).

**SVG icon set minimum:**
- Arrow left (back navigation)
- Download (upload evidence)
- Camera (photo evidence)
- Link (link evidence)
- Check circle (valid)
- Alert triangle (warning)
- X circle (invalid)
- Star (hero badge)
- Plus (add item)
- Trash (delete)
- Menu hamburger (mobile nav)
- User (profile/avatar)
- Shield (courtroom)
- Chart bar (score)
- Clock (mission deadline)
- Flag (decision outcome)

Stroke width: 1.5–2px, strokeLinecap: round, strokeLinejoin: round.
Ukuran: 20x20 (inline), 40x40 (icon card), 64x64 (hero section).
Warna: mengikuti `currentColor` atau `stroke="var(--brand-orange)"` untuk ikon yang perlu aksen.

---

## File Delivery

Semua output desain dikirim sebagai single-page HTML artifact yang dapat berdiri sendiri, bukan file gambar. HTML berisi semua komponen, token, dan screen mockup yang dapat di-render di browser tanpa dependencies eksternal. CSS custom properties untuk semua design token sehingga perubahan warna/typography cukup di satu tempat.
