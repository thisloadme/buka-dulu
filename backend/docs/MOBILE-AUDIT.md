# BukaDulu Mobile — UI/UX Compliance Audit

**Audit target:** Flutter mobile app (`mobile/`) vs DESIGN.md
**Tanggal:** 2026-06-17
**Status:** 🔴 28 ketidaksesuaian ditemukan (9 critical, 12 major, 7 minor)

---

## Prioritas Critical (harus diperbaiki)

### C-01. ColorScheme seed color salah — hijau bukan orange

| File | Baris |
|------|-------|
| `mobile/lib/config/theme.dart` | 9 |

**Temuan:** `ColorScheme.fromSeed(seedColor: const Color(0xFF1A6B3C))` — seed-nya **hijau**. Semua komponen Material (button, chip, toggle, badge) mewarisi warna hijau, bukan orange.

**Seharusnya:** `seedColor: const Color(0xFFea580c)` — brand primary orange (#ea580c).

**Referensi DESIGN.md:** Color Palette → Orange #ea580c.

---

### C-02. Gradient di Splash Screen

| File | Baris |
|------|-------|
| `mobile/lib/presentation/pages/splash_screen.dart` | 65-72, 87-89 |

**Temuan:**
- Background splash menggunakan `LinearGradient` dark #1c1917 → #292524
- Logo container menggunakan `LinearGradient` orange #ea580c → #f59e0b

**Seharusnya:** Semua background solid, tanpa gradient. Splash bg bisa solid `#1c1917`, logo solid `#ea580c`.

**Referensi DESIGN.md:** Prinsip #5 — "Tidak ada gradient. Semua background solid."

---

### C-03. Material Icons di seluruh UI — bukan SVG

| File | Baris | Icon yang dipakai |
|------|-------|-------------------|
| Semua page | seluruh | `Icons.*` |

**Temuan:** Seluruh aplikasi menggunakan Material Design icons (Flutter `Icons.*`), bukan SVG inline.

**Seharusnya:** Semua ikon berupa **SVG inline** 20x20 / 40x40 / 64x64 dengan stroke `1.5-2px`, `strokeLinecap: round`, `strokeLinejoin: round`. Tidak ada font icons.

**Referensi DESIGN.md:** Iconography → "Semua ikon menggunakan SVG inline (bukan emoji, bukan font icon)."

---

### C-04. Emoji di UI — bukan SVG

| File | Baris | Emoji |
|------|-------|-------|
| `mobile/lib/presentation/pages/cost/cost_page.dart` | 201 | 🟢 Sehat, 🟡 Tipis, 🔴 Berbahaya |
| `mobile/lib/presentation/pages/cost/cost_page.dart` | 207 | unit/bulan |
| `mobile/lib/presentation/pages/score/score_page.dart` | 118-120 | 🚀, 🔄, 🔀, 🛑 |
| `mobile/lib/presentation/pages/mission/mission_board_page.dart` | 135 | ⏱ |

**Seharusnya:** Semua status indicator pakai SVG icon + teks. Contoh: ✅ valid → SVG check-circle, 🟢 Sehat → SVG circle fill green + teks "Sehat".

**Referensi DESIGN.md:** Prinsip #5 — "Tidak ada emoji. Ikon dengan SVG."

---

### C-05. Heading/body font weight tidak sesuai

| File | Baris | Weight |
|------|-------|--------|
| `mobile/lib/presentation/pages/auth/login_page.dart` | 52 | `FontWeight.bold` (700) |
| `mobile/lib/presentation/pages/auth/register_page.dart` | 56 | `FontWeight.bold` (700) |
| `mobile/lib/presentation/pages/dashboard/dashboard_page.dart` | 142 | `FontWeight.w600` |
| Banyak page | — | `FontWeight.w600` di title |

**Seharusnya:**
- Display Hero / Display Large → **weight 300**
- Body Large / Body Text → **weight 300**
- Hanya Micro (0.75rem uppercase) yang weight **500**
- Judul section boleh weight **400**

**Referensi DESIGN.md:** Typography → semua weight 300 kecuali Micro yang weight 500.

---

### C-06. Stage colors di dashboard tidak cocok DESIGN.md

| File | Baris |
|------|-------|
| `mobile/lib/presentation/pages/dashboard/dashboard_page.dart` | 11-25 |

**Temuan (mapping salah vs seharusnya):**

| Stage | Flutter (salah) | DESIGN.md (benar) |
|-------|-----------------|-------------------|
| `draft` | `Colors.grey` | #78716c (gray) ✅ |
| `idea_defined` | `Colors.blue` 🔴 | #ea580c (orange) |
| `customer_defined` | `Colors.green` 🔴 | #ea580c (orange) |
| `sku_focused` | `Colors.orange` | #ea580c (orange) ✅ |
| `cost_evaluated` | `Colors.amber` ⚠️ | #ea580c (orange) — deket tp beda |
| `mission_active` | `Colors.purple` 🔴 | #3b82f6 (blue) |
| `evidence_submitted` | `Colors.pink` 🔴 | #f59e0b (amber) |
| `evidence_reviewed` | `Colors.teal` 🔴 | #22c55e (green) |
| `ready_to_decide` | `Colors.indigo` 🔴 | #8b5cf6 (purple) |
| `continue` | `Colors.green` ✅ | #22c55e (green) |
| `repeat` | `Colors.orange` ✅ | #f59e0b (amber) |
| `pivot` | `Colors.amber` ✅ | #f59e0b (amber) |
| `stop` | `Colors.red` ✅ | #ef4444 (red) |

---

### C-07. Score breakdown bar warna random — harus orange solid

| File | Baris |
|------|-------|
| `mobile/lib/presentation/pages/score/score_page.dart` | 80-85 |

**Temuan:** Score bar per komponen pakai warna berbeda (blue, orange, green, purple, teal, pink).

**Seharusnya:** Semua bar fill pakai `#ea580c` (orange) — konsisten dengan brand. Bar bg `#e7e5e4`.

**Referensi DESIGN.md:** Data Visualization → Score Breakdown Bar → "Bar fill: #ea580c"

---

### C-08. Body text pakai `Colors.grey[600]` — bukan #57534e

| File | Baris | Warna |
|------|-------|-------|
| `login_page.dart` | 54 | `Colors.grey[600]` (~#757575) |
| `idea_capture_page.dart` | 60 | `Colors.grey[600]` |
| `idea_result_page.dart` | 177 | `Colors.grey[600]` |
| `dashboard_page.dart` | 78 | `Colors.grey[600]` |
| `dashboard_page.dart` | 145 | `Colors.grey[600]` |
| Banyak page | — | `Colors.grey[600]` |

**Seharusnya:** `Color(0xFF57534e)` — brand body color. Grey Material tidak cocok dengan palet warm tone BukaDulu.

**Referensi DESIGN.md:** Color Palette → Body #57534e.

---

### C-09. Button primary tidak orange — default Material dari seed hijau

| File | Baris |
|------|-------|
| `mobile/lib/config/theme.dart` | 29-35 |

**Temuan:** `ElevatedButton.styleFrom` mewarisi colorScheme — karena seed hijau, button default jadi hijau.

**Seharusnya:** Primary button explicit `backgroundColor: Color(0xFFea580c)`, `foregroundColor: Colors.white`, `borderRadius: 6px`.

**Referensi DESIGN.md:** Button System → Primary: bg #ea580c, text White, border-radius 6px.

---

## Prioritas Major (perlu diperbaiki)

### M-01. Tidak ada Ghost/White/Outline-Light button variants

| File | Status |
|------|--------|
| `mobile/lib/config/theme.dart` | Missing |

**Temuan:** Hanya `ElevatedButton` yang dikonfigurasi di theme. Tidak ada:
- `OutlinedButton` untuk Ghost variant (transparent bg, orange text, #fed7aa border)
- `White button` untuk CTA di atas dark bg (white bg, orange text)
- `OutlineLight` untuk dark bg (transparent, white text, white border 0.25)

**Seharusnya:** Tambahkan `outlinedButtonTheme` dan `textButtonTheme` di AppTheme.

**Referensi DESIGN.md:** Button System — 5 varian.

---

### M-02. Tidak ada AppShell layout dengan stage indicator

| File | Status |
|------|--------|
| Seluruh page | Missing |

**Temuan:** DESIGN.md menspesifikasikan venture flow layout sebagai:
```
[App Shell] — back button + title + stage indicator
[Content] — single-column, max 720px
[Bottom CTA] — primary action button, sticky on mobile
```

Flutter saat ini: masing-masing page pakai `Scaffold + AppBar(title:)` sendiri-sendiri. Tidak ada shared component untuk:
- Stage indicator / progress bar
- Bottom CTA sticky
- Stage badge pill

**Referensi DESIGN.md:** Layout Structure → Venture flow.

---

### M-03. Tidak ada skeleton loading states

| File | Status |
|------|--------|
| Semua page | `CircularProgressIndicator()` doang |

**Temuan:** DESIGN.md menspesifikasikan 4 loading patterns:
- Skeleton card (gray #e7e5e4, rounded 8px, pulse animation)
- Button spinner
- Full page skeleton (3-4 cards stacked)
- Inline shimmer

Flutter hanya pakai `CircularProgressIndicator()` di semua halaman.

**Referensi DESIGN.md:** Loading States.

---

### M-04. Tidak ada empty state SVG illustrations

| File | Baris |
|------|-------|
| `mobile/lib/presentation/pages/dashboard/dashboard_page.dart` | 76 |

**Temuan:** Dashboard empty state pakai `Icon(Icons.lightbulb_outline, ...)` — Material icon ukuran 80, bukan SVG.

**Seharusnya:** SVG ilustrasi sesuai DESIGN.md:
- Dashboard: SVG ilustrasi ide kosong
- Mission board: SVG checklist
- Evidence: SVG upload icon

**Referensi DESIGN.md:** Empty States.

---

### M-05. Evidence page — Camera/image tidak terimplementasi

| File | Baris |
|------|-------|
| `mobile/lib/presentation/pages/evidence/evidence_upload_page.dart` | 97-114 |

**Temuan:** Bagian image upload cuma placeholder "Fitur foto menyusul". DESIGN.md menspesifikasikan 3 upload methods functional: 📷 [Ambil Foto], 📝 [Catatan], 🔗 [Link/URL].

**Seharusnya:** Integrasi `image_picker` untuk camera/gallery upload.

---

### M-06. Score gauge warna — harus orange, bukan score-tier color

| File | Baris |
|------|-------|
| `mobile/lib/presentation/pages/score/score_page.dart` | 61-66 |

**Temuan:** Circle gauge border menggunakan `_scoreColor` (dynamic green/orange/red). DESIGN.md bilang "Fill color: #ea580c (or dynamic based on score tier)".

Meskipun ada "or dynamic", konsistensi brand lebih baik pake orange solid untuk gauge. Score tier color bisa diterapkan di decision card, bukan di gauge utama.

**Seharusnya:** Border gauge = `Color(0xFFea580c)` dan teks score = `Color(0xFFea580c)`.

---

### M-07. Decision screen inline di ScorePage — harus page terpisah

| File | Baris |
|------|-------|
| `mobile/lib/presentation/pages/score/score_page.dart` | 100-134 |

**Temuan:** Decision result di-render di dalam ScorePage, bukan page terpisah. Juga tidak ada route `/venture/:id/decision` di router.

DESIGN.md punya halaman Final Decision terpisah dengan 2 varian (CONTINUE / STOP), masing-masing dengan layout kaya: icon besar, alasan, langkah selanjutnya, dan CTA.

**Seharusnya:** Buat `DecisionPage` terpisah, route di router, navigasi setelah "Hasilkan Keputusan Akhir" ditekan.

---

### M-08. Tidak ada Founder Courtroom page

| File | Status |
|------|--------|
| Router | Missing route |

**Temuan:** DESIGN.md menspesifikasikan Founder Courtroom screen dengan 3 perspective cards (Calon Pembeli, Operator Dapur, Reviewer Bisnis) + objection list + weakest assumptions. Tidak ada di Flutter.

**Seharusnya:** Buat `courtroom_page.dart` + route `courtroom` di router.

---

### M-09. Tidak ada Notifications

| File | Status |
|------|--------|
| Seluruh app | Missing |

**Temuan:** Tidak ada notification provider, tidak ada notification badge/bell di dashboard, tidak ada halaman notifikasi. DESIGN.md menspesifikasikan in-app notifications untuk:
- Mission deadline
- Review selesai
- Score update
- Reactivation prompt (pasif >48 jam)

---

### M-10. Venture card tidak menampilkan score dan mission stats

| File | Baris |
|------|-------|
| `mobile/lib/presentation/pages/dashboard/dashboard_page.dart` | 108-157 |

**Temuan:** Venture card hanya menampilkan: stage badge, nama, kategori, tanggal, progress bar. Tidak ada:
- Score (numerik / total)
- Mission stats (completed/total)
- Tombol [Lanjutkan] langsung ke halaman sesuai stage

DESIGN.md dashboard spec: setiap venture card punya score + progress + mission stats + tombol Lanjutkan.

---

### M-11. Login page tidak pakai logo image

| File | Baris |
|------|-------|
| `mobile/lib/presentation/pages/auth/login_page.dart` | 52 |

**Temuan:** Login hanya menampilkan teks "BukaDulu" dengan `headlineLarge`. DESIGN.md menspesifikasikan logo (40x40) + brand name.

**Seharusnya:** Pakai `Image.asset('assets/bukadulu.png', width: 40, height: 40)` seperti di splash screen.

---

### M-12. Mission board progress — tidak ada 14-day indicator

| File | Baris |
|------|-------|
| `mobile/lib/presentation/pages/mission/mission_board_page.dart` | 74-87 |

**Temuan:** Progress hanya menampilkan `done/total` ratio. DESIGN.md menspesifikasikan 14-day progress indicator:
```
📅 Hari ke-3 dari 14
████░░░░░░░░  22%
```

---

## Prioritas Minor (nice to fix)

### m-01. Tidak ada shadow sesuai DESIGN.md

| File | Detail |
|------|--------|
| `mobile/lib/config/theme.dart` | `CardThemeData(elevation: 1)` |

DESIGN.md: `rgba(80,50,20,0.12) 0px 50px 80px -40px, rgba(0,0,0,0.06) 0px 20px 40px -20px`. Flutter Material `elevation` tidak menghasilkan shadow persis sama — butuh custom `BoxShadow`.

### m-02. Tidak ada route untuk `/download` dan mentor dashboard

DESIGN.md screen tree:
- `/download` — post-auth landing
- `/mentor/dashboard` — role mentor
- `/mentor/mentee/:vid` — detail mentee

Router saat ini tidak punya route ini.

### m-03. Register page — Konfirmasi Password field

`mobile/lib/presentation/pages/auth/register_page.dart` line 88: Konfirmasi Password ada, sesuai spec UI Flow. Tapi DESIGN.md screen spec tidak menampilkannya — konsistensi dokumen perlu diperiksa.

### m-04. Mission card — priority badge pakai `Text` bukan colored pill

DESIGN.md: Mission Priority Badge dibuat sebagai container pill dengan bg `rgba(239,68,68,0.1)` untuk high. Flutter pakai container dengan `borderRadius: 8` dan icon di dalamnya — fungsi sama tapi tidak mengikuti pill pattern persis.

### m-05. Cost page — `_summaryRow` pakai `Colors.grey` bukan brand body

`mobile/lib/presentation/pages/cost/cost_page.dart` line 231: `Text(label, style: const TextStyle(color: Colors.grey))` — harusnya `Color(0xFF57534e)`.

### m-06. Score page — `Colors.grey[600]` dan `Colors.grey[700]` untuk teks

`mobile/lib/presentation/pages/score/score_page.dart` lines 146, 126: pakai `Colors.grey` — harusnya brand body #57534e.

### m-07. Evidence page — `Colors.grey[300/400/500]` untuk border/placeholder

`mobile/lib/presentation/pages/evidence/evidence_upload_page.dart`: border, placeholder, icon camera pakai `Colors.grey[300/400/500]` — harusnya brand border/warna sesuai DESIGN.md.

---

## Compliance Summary

| Kategori | Total | Critical | Major | Minor |
|----------|-------|----------|-------|-------|
| Color Palette | 6 | 4 | 1 | 1 |
| Typography | 2 | 2 | 0 | 0 |
| Component Design | 6 | 2 | 3 | 1 |
| Iconography | 2 | 2 | 0 | 0 |
| Screen/Layout | 7 | 0 | 5 | 2 |
| Data Visualization | 2 | 1 | 1 | 0 |
| Empty/Loading States | 2 | 0 | 2 | 0 |
| **TOTAL** | **28** | **9** | **12** | **7** |

---

## Quick Wins (bisa diperbaiki dalam <30 menit)

1. **C-01** — Ganti seed color `0xFF1A6B3C` → `0xFFea580c` di `theme.dart`
2. **C-05** — Ganti `FontWeight.bold` → `FontWeight.w300` atau `w400` di heading
3. **C-08** — Ganti `Colors.grey[600]` → `Color(0xFF57534e)` di semua page
4. **C-09** — Tambah explicit `backgroundColor: Color(0xFFea580c)` di ElevatedButton theme
5. **C-02** — Ganti gradient ke solid di splash screen
6. **m-05, m-06, m-07** — Ganti `Colors.grey` ke brand colors

## Perbaikan Jangka Menengah (30-120 menit)

7. **C-04** — Ganti emoji dengan SVG icons di cost/score pages
8. **C-06** — Update stage color mapping sesuai DESIGN.md
9. **M-01** — Tambah OutlinedButton/TextButton theme variants
10. **M-02** — Buat AppShell widget + stage indicator
11. **M-12** — Tambah 14-day progress bar di mission board
12. **M-11** — Tambah logo image di login page

## Perbaikan Jangka Panjang (120+ menit)

13. **C-03** — Migrasi semua Material Icons ke SVG inline
14. **M-03** — Implementasi skeleton loading states
15. **M-04** — Buat SVG illustrations untuk empty states
16. **M-07, M-08** — Buat DecisionPage + FounderCourtroom page
17. **M-09** — Implementasi notification system
