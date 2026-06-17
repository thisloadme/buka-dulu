# State Machine & Business Rules — BukaDulu MVP
## Versi 1.0 | Stage-Gate System

---

## 1. State Machine Diagram

```
                    ┌──────────────────────────────────────────────┐
                    │              VENTURE STATE MACHINE            │
                    └──────────────────────────────────────────────┘

     ┌──────────┐
     │  DRAFT   │
     └────┬─────┘
          │ US-021: Idea structured & confirmed
          ▼
  ┌────────────────┐
  │  IDEA_DEFINED  │◀──── (reset if user re-does idea)
  └───────┬────────┘
          │ US-030: Customer segment defined & locked
          ▼
  ┌────────────────────┐
  │ CUSTOMER_DEFINED   │
  └───────┬────────────┘
          │ US-041/042: SKU chosen & locked (max 3, min 1)
          ▼
  ┌────────────────┐
  │  SKU_FOCUSED   │
  └───────┬────────┘
          │ US-051: Cost & margin evaluation confirmed
          ▼
  ┌───────────────────┐
  │  COST_EVALUATED   │
  └───────┬───────────┘
          │ US-060: Missions generated & first mission accepted
          ▼
  ┌──────────────────┐
  │  MISSION_ACTIVE  │◀──── (loop: mission → evidence → review)
  └───────┬──────────┘
          │ US-062/070: Evidence uploaded for a mission
          ▼
  ┌──────────────────────┐
  │  EVIDENCE_SUBMITTED  │──── Auto-transition after AI review
  └───────┬──────────────┘
          │ US-080: AI review completed
          ▼
  ┌─────────────────────┐
  │  EVIDENCE_REVIEWED  │──┐
  └───────┬─────────────┘  │ (more missions remain)
          │ All missions complete + all evidence reviewed
          ▼
  ┌──────────────────┐
  │  READY_TO_DECIDE │
  └───────┬──────────┘
          │ US-091: Final decision generated
          ▼
  ┌──────────┐  ┌────────┐  ┌───────┐  ┌──────┐
  │ CONTINUE │  │ REPEAT │  │ PIVOT │  │ STOP │
  └──────────┘  └────────┘  └───────┘  └──────┘
       │             │          │          │
       │         (repeat       │          │
       │          mission      │          │
       │          stage)       │          │
       └───────────────────────┴──────────┘
              End state — no further transitions
```

---

## 2. Gate Transition Rules

Setiap transisi state punya **gate** yang harus dipenuhi:

### Gate 1: Draft → Idea Defined
```
REQUIRED:
  ✅ Raw ide mentah (≥ 20 karakter)
  ✅ AI sudah menghasilkan structured concept
  ✅ Founder sudah konfirmasi (is_locked = true)
  
VALIDATION:
  - One-line concept tidak boleh kosong
  - Target customer terisi
  - Setidaknya 1 key assumption teridentifikasi
```

### Gate 2: Idea Defined → Customer Defined
```
REQUIRED:
  ✅ Customer segment sudah dipilih/dibuat
  ✅ Setidaknya 2 dari 5 atribut terisi (usia, konteks beli, budget, lokasi, momen)
  
VALIDATION:
  - Sistem mengecek apakah segment terlalu general ("semua orang")
  - Jika terlalu general → warning tapi tidak block gate
```

### Gate 3: Customer Defined → SKU Focused
```
REQUIRED:
  ✅ Minimal 1 menu dengan status 'active'
  ✅ Maksimal 3 menu dengan status 'active' (dipaksa sistem)
  ✅ Hero SKU sudah ditentukan
  
VALIDATION:
  - Menu yang tidak dipilih otomatis 'deferred'
  - Complexity score sudah dihitung untuk semua menu active
```

### Gate 4: SKU Focused → Cost Evaluated
```
REQUIRED:
  ✅ Cost summary sudah dihitung untuk semua SKU active
  ✅ Founder sudah konfirmasi cost data (is_locked = true)
  
VALIDATION:
  - HPP per porsi > 0
  - Margin tidak boleh NULL
  - Jika margin 'berbahaya' → strong warning, tapi tidak block gate
```

### Gate 5: Cost Evaluated → Mission Active
```
REQUIRED:
  ✅ Minimal 1 mission sudah di-generate oleh sistem
  ✅ Minimal 1 mission sudah di-accept oleh founder
  
TRANSITION:
  - Sistem generate daily missions berdasarkan stage dan risk profile
  - Founder tidak bisa start missions tanpa melewati gate ini
```

### Gate 6: Mission Active ↔ Evidence Submitted (looping)
```
MISSION → EVIDENCE:
  ✅ Founder mengunggah evidence untuk mission aktif
  
EVIDENCE → REVIEWED:
  ✅ AI review selesai (verdict + rationale generated)
  
REVIEWED → MISSION ACTIVE (back):
  ⬅ Ada mission lain yang belum selesai
  
ALL MISSIONS COMPLETE:
  ➡ Semua mission aktif sudah completed + evidence reviewed
```

### Gate 7: Ready to Decide → Final Decision
```
REQUIRED FOR ALL:
  ✅ Semua komponen score terisi (tidak ada yang 0)
  ✅ Setidaknya 1 evidence dengan verdict 'valid'
  
THEN APPLY DECISION MATRIX:
  → Lihat section 4
```

---

## 3. Scoring Engine

### 3.1 Score Components

| Komponen | Range | Bobot | Sumber Data |
|---|---|---|---|
| Clarity | 0-100 | 10% | Idea concept completeness, customer segment specificity |
| Focus | 0-100 | 10% | Jumlah SKU active, hero identified, deferred menu ratio |
| Economics | 0-100 | 25% | Gross margin percentage, break-even feasibility |
| Execution | 0-100 | 20% | Mission completion rate, on-time vs delayed ratio |
| Evidence | 0-100 | 25% | Jumlah & kualitas evidence valid, review verdict distribution |
| Market Response | 0-100 | 10% | Pre-order count, survey positive rate, sampling feedback |

### 3.2 Formula

```
Total Score = (Clarity × 0.10) + (Focus × 0.10) + (Economics × 0.25) 
            + (Execution × 0.20) + (Evidence × 0.25) + (MarketResponse × 0.10)
```

### 3.3 Component Detail

**Clarity Score (0-100):**
- 80-100: Konsep jelas, customer spesifik, asumsi terdefinisi dengan baik
- 50-79: Konsep cukup jelas, customer agak umum
- 0-49: Konsep vague atau customer terlalu general

**Focus Score (0-100):**
- 80-100: 1-2 SKU, hero jelas, deferred menu terdokumentasi
- 50-79: 3 SKU, hero cukup jelas
- 0-49: >3 SKU coba dipaksa, hero tidak jelas

**Economics Score (0-100):**
- 80-100: Margin >40% (sehat), break-even realistis
- 50-79: Margin 20-40% (tipis)
- 0-49: Margin <20% (berbahaya)

**Execution Score (0-100):**
- `completed_missions / total_missions_accepted × 100`
- Bisa +10 bonus jika semua selesai sebelum deadline

**Evidence Score (0-100):**
- `(valid_count × 100 + weak_count × 50) / total_evidence`
- Suspicious/invalid = 0
- Bonus +10 jika ada evidence dari ≥3 tipe berbeda (image, text, link)

**Market Response Score (0-100):**
- Dari AI review evidence: analisis sentimen respons pasar
- Pre-order konfirmasi: +30
- Testimoni positif: +20
- Feedback negatif: -10

---

## 4. Decision Matrix

Output keputusan berdasarkan **Total Score** + **Evidence Quality** + **Economics Health**:

```
                    ┌─────────────────────────────────────────────────────┐
                    │                  EVIDENCE QUALITY                    │
                    │    HIGH (≥60)        MEDIUM (30-59)     LOW (<30)   │
┌─────────────────┼─────────────────────────────────────────────────────┤
│ ECONOMICS       │                     │                     │          │
│ HEALTHY (≥70)   │   CONTINUE          │   CONTINUE          │  REPEAT  │
│ (margin sehat)  │                     │   (dengan warning)  │  mission │
├─────────────────┼─────────────────────┼─────────────────────┼──────────┤
│ ECONOMICS       │                     │                     │          │
│ MODERATE (40-69)│   CONTINUE          │   REPEAT            │  STOP    │
│ (margin tipis)  │   (rekomendasi      │   mission           │          │
│                 │    pivot nanti)     │                     │          │
├─────────────────┼─────────────────────┼─────────────────────┼──────────┤
│ ECONOMICS       │                     │                     │          │
│ LOW (<40)       │   PIVOT             │   PIVOT             │  STOP    │
│ (margin bahaya) │   (ubah model       │   atau STOP         │          │
│                 │    jualan)          │                     │          │
└─────────────────┴─────────────────────┴─────────────────────┴──────────┘
```

### Decision Outcomes

| Outcome | Arti | Next Action |
|---|---|---|
| **CONTINUE** | Ide layak, bukti cukup. Founder siap ke fase berikutnya (uji jual skala kecil). | Saran: mulai produksi terbatas, setup pre-order. |
| **REPEAT** | Evidence belum cukup kuat. Ulangi misi dengan scope lebih baik. | Saran: perbaiki kualitas evidence, target responden lebih spesifik. |
| **PIVOT** | Asumsi inti tidak terdukung. Ubah salah satu: customer, produk, atau channel jual. | Saran: ubah target customer atau ganti hero SKU. |
| **STOP** | Ide tidak layak dalam bentuk sekarang. Tidak disarankan lanjut. | Saran: dokumentasikan pelajaran, coba ide lain. |

---

## 5. Mission Generation Rules

### 5.1 Mission Template Selection

```
IF stage = 'MISSION_ACTIVE' AND cost margin > 40%:
  Prioritaskan: polling, pre-order, sampling
  
IF stage = 'MISSION_ACTIVE' AND cost margin < 20%:
  Prioritaskan: harga_test, interview (cari tahu kenapa margin rendah)

IF first mission:
  Always: polling ringan + interview (2 mission first day)

IF no mission completed in 48h:
  Generate: reactivation prompt + 1 mission sangat ringan (15 min)
```

### 5.2 Priority Queue

Mission diurutkan berdasarkan:
1. Deadline terdekat
2. Priority (high > medium > low)
3. Estimated time (shortest first untuk daily queue)

### 5.3 Evidence Requirement per Mission Type

| Mission Type | Min Evidence Required |
|---|---|
| Polling | Screenshot hasil polling + jumlah responden |
| Pre-order | Screenshot chat/list order + jumlah konfirmasi |
| Sampling | Foto produk dikirim + feedback yang diterima |
| Titip jual | Foto produk di titip + laporan penjualan |
| Interview | Catatan/recap wawancara minimal 3 responden |
| Harga test | Screenshot hasil test harga + respon |
| Observation | Foto lokasi + catatan |

---

## 6. Business Rules Summary

| ID | Rule | Level | Enforcement |
|---|---|---|---|
| BR-001 | Founder hanya bisa akses venture sendiri | Authorization | Middleware, handler |
| BR-002 | Max 3 active SKU per venture | Data integrity | DB trigger (PG) / app logic (SQLite) |
| BR-003 | Min 1 active SKU required before cost stage | Business | Service layer |
| BR-004 | Mission cannot be completed without evidence | Business | Service layer |
| BR-005 | Score hanya bisa dihitung jika semua komponen > 0 | Business | Scoring service |
| BR-006 | Human override must log reviewer identity | Audit | Review service |
| BR-007 | Cost re-calculation required jika ingredient berubah | Data integrity | Service layer |
| BR-008 | Score re-calculation required jika evidence baru | Event-driven | Scoring service |
| BR-009 | Decision hanya bisa di-generate 1x per venture (sampai reset) | Business | Decision service |
| BR-010 | Venture yang sudah CONTINUE/STOP tidak bisa diedit (read-only) | Business | Middleware |
