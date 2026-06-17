import { Link } from 'react-router-dom'
import { useState } from 'react'

const NAV_ITEMS = ['Cara Kerja', 'Fitur', 'Harga', 'FAQ'] as const

function Logo() {
  return (
    <Link to="/" className="logo">
      <img src="/bukadulu.png" alt="BukaDulu" width="32" height="32" />
      BukaDulu
    </Link>
  )
}

function NavLinks({ onNavClick }: { onNavClick?: () => void }) {
  return (
    <>
      {NAV_ITEMS.map(item => (
        <a key={item} href={`#${item.toLowerCase().replace(/\s/g, '-')}`} onClick={onNavClick}>{item}</a>
      ))}
      <Link to="/login" onClick={onNavClick}>Masuk</Link>
      <Link to="/register" className="btn btn-primary" onClick={onNavClick}>Coba Gratis</Link>
    </>
  )
}

function Navbar() {
  const [open, setOpen] = useState(false)
  return (
    <nav>
      <div className="container">
        <Logo />
        <div className={`nav-links${open ? ' open' : ''}`} id="navLinks">
          <NavLinks onNavClick={() => setOpen(false)} />
        </div>
        <button className="nav-toggle" aria-label="Menu" onClick={() => setOpen(o => !o)}>
          <span /><span /><span />
        </button>
      </div>
    </nav>
  )
}

function SectionIntro({ label, title, desc }: { label: string; title: string; desc: string }) {
  return (
    <div className="section-intro">
      <div className="micro">{label}</div>
      <h2 className="display-large">{title}</h2>
      <p className="body-large" style={{ marginTop: 16 }}>{desc}</p>
    </div>
  )
}

function ProblemCard({ icon, title, desc, delay }: { icon: React.ReactNode; title: string; desc: string; delay: string }) {
  return (
    <div className={`problem-card fade-in ${delay}`}>
      <div className="problem-card-icon">{icon}</div>
      <h3>{title}</h3>
      <p>{desc}</p>
    </div>
  )
}

function StepCard({ num, title, desc, delay }: { num: string; title: string; desc: string; delay: string }) {
  return (
    <div className={`step fade-in ${delay}`}>
      <div className="step-number">{num}</div>
      <h3>{title}</h3>
      <p>{desc}</p>
    </div>
  )
}

function FeatureCard({ tag, icon, title, desc, delay }: { tag: string; icon: React.ReactNode; title: string; desc: string; delay: string }) {
  return (
    <div className={`feature-card fade-in ${delay}`}>
      <div className="feature-tag">{tag}</div>
      <div className="feature-icon">{icon}</div>
      <h3>{title}</h3>
      <p>{desc}</p>
    </div>
  )
}

type FaqItem = { q: string; a: string }

function FaqSection({ items }: { items: FaqItem[] }) {
  const [openIdx, setOpenIdx] = useState<number | null>(null)
  return (
    <section className="faq" id="faq">
      <div className="container">
        <SectionIntro label="FAQ" title="Pertanyaan yang sering diajukan" desc="" />
        <div className="faq-list">
          {items.map((item, i) => (
            <div key={i} className="faq-item">
              <button
                className={`faq-question${openIdx === i ? ' open' : ''}`}
                onClick={() => setOpenIdx(openIdx === i ? null : i)}
              >
                {item.q}
              </button>
              <div className={`faq-answer${openIdx === i ? ' open' : ''}`}
                dangerouslySetInnerHTML={{ __html: item.a }} />
            </div>
          ))}
        </div>
      </div>
    </section>
  )
}

export default function Landing() {
  const faqItems: FaqItem[] = [
    {
      q: 'Apa bedanya BukaDulu sama aplikasi business plan generator?',
      a: 'BukaDulu bukan business plan generator. Kami tidak bikin dokumen 40 halaman yang tidak pernah dibaca. Kami bikin kamu bergerak: menentukan SKU, menghitung margin real, menjalankan misi lapangan, dan mengumpulkan bukti. Perbedaan fundamental: kami fokus ke <strong>eksekusi pra-peluncuran</strong>, bukan rencana.',
    },
    {
      q: 'Produk saya belum jadi, apa tetap bisa pakai?',
      a: 'Justru itu target kami. Kalau produk kamu sudah jadi dan laku, kamu mungkin tidak butuh BukaDulu lagi. Kami khusus untuk fase sebelum jualan — saat masih bingung, ragu, dan belum punya bukti pasar.',
    },
    {
      q: 'Apa saya harus bisa masak dulu?',
      a: 'Tidak harus. Banyak founder sukses mulai dari jualan frozen food titipan, snack kemasan, atau minuman sederhana. BukaDulu cocok buat siapa saja yang serius mau jualan makanan atau minuman, terlepas dari skill memasak.',
    },
    {
      q: 'Bagaimana kalau setelah 14 hari ide saya tidak layak?',
      a: 'Itu <strong>hasil terbaik</strong> yang bisa kamu dapat. Lebih baik tahu dari sekarang dengan modal Rp149 ribu daripada buka gerai, rugi 30 juta, dan tutup 3 bulan kemudian. BukaDulu tegas: "Stop" bukan kegagalan, tapi keputusan bisnis yang cerdas.',
    },
    {
      q: 'Apa ada mentor yang mendampingi?',
      a: 'Untuk MVP, sistem yang menjadi "mentor" kamu: tegas, berbasis data, dan tidak membiarkan kamu menunda. Ke depannya kami akan buka Mentor Dashboard untuk pendamping manusia — termasuk program cohort dengan mentor bisnis sungguhan.',
    },
    {
      q: 'Ada garansi tidak kalau saya berhasil jualan?',
      a: 'Tidak ada. BukaDulu bukan jaminan sukses. Kami hanya sistem yang memaksimalkan peluang kamu dengan data nyata. Keputusan tetap di tangan kamu. Yang kami jamin: setelah 14 hari, kamu akan punya <strong>bukti, bukan sekadar keyakinan</strong>.',
    },
  ]

  return (
    <div>
      <Navbar />

      {/* HERO */}
      <section className="hero" id="hero">
        <div className="container">
          <div className="hero-content">
            <div className="hero-badge fade-in">Riset dan overthinking menjebak. Mulai buktikan sekarang.</div>
            <h1 className="display-hero fade-in delay-1">
              Punya ide jualan makanan?<br />
              <strong style={{ fontWeight: 400, color: 'var(--brand-orange)' }}>Buktikan dulu dalam 14 hari.</strong>
            </h1>
            <p className="body-large fade-in delay-2">
              Bukan business plan. Bukan analisis AI yang tidak jelas. BukaDulu adalah sistem eksekusi
              yang memaksa kamu bergerak — dari ide mentah ke bukti pasar nyata — dalam 14 hari.
              Gratis untuk sprint pertamamu.
            </p>
            <div className="hero-cta fade-in delay-3">
              <Link to="/register" className="btn btn-primary btn-lg">Cek Ide Gratis</Link>
              <a href="#cara-kerja" className="btn btn-ghost btn-lg">Lihat Cara Kerja</a>
            </div>
            <div className="hero-stats fade-in delay-4">
              <div className="stat-item">
                <div className="stat-number">14 hari</div>
                <div className="stat-label">Dari ide ke bukti pasar</div>
              </div>
              <div className="stat-item">
                <div className="stat-number">1-3 SKU</div>
                <div className="stat-label">Fokus, bukan banyak menu</div>
              </div>
              <div className="stat-item">
                <div className="stat-number">100%</div>
                <div className="stat-label">Berdasarkan bukti nyata</div>
              </div>
            </div>
          </div>
          <div className="screenshot-area fade-in delay-5">
            <div className="screenshot-card">
              <div className="micro" style={{ color: 'var(--brand-orange)', marginBottom: 16 }}>Preview Aplikasi</div>
              <div className="app-preview">
                <div className="preview-step">
                  <div className="step-label">Langkah 1</div>
                  <h4>Compress Ide</h4>
                  <p>Tulis ide mentah, AI bantu menyusun konsep bisnis yang tajam</p>
                </div>
                <div className="preview-step">
                  <div className="step-label">Langkah 2</div>
                  <h4>Fokus Menu</h4>
                  <p>Pilih 1-3 SKU yang paling realistis buat diuji pasar</p>
                </div>
                <div className="preview-step">
                  <div className="step-label">Langkah 3</div>
                  <h4>Jalankan Misi</h4>
                  <p>Dapat tugas harian dan upload bukti sebagai progress</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* MASALAH */}
      <section className="problem" id="masalah">
        <div className="container">
          <SectionIntro label="Masalah" title="Kebanyakan calon pebisnis F&B berhenti di kepala sendiri" desc="Bukan karena ide jelek. Tapi karena tidak ada yang memaksa mereka bergerak." />
          <div className="problem-grid">
            <ProblemCard delay="" title="Bingung mulai dari mana" desc="Ide masih campur aduk. Tidak ada langkah pertama yang jelas. Akhirnya cuma dipikirin terus tanpa aksi."
              icon={<svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><circle cx="12" cy="12" r="10"/><path d="M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3"/><path d="M12 17h.01"/></svg>} />
            <ProblemCard delay="delay-1" title="Pengen jual banyak menu" desc="Pikiran langsung lompat ke menu 30 item, padahal belum ada satu pelanggan pun yang pernah mencoba."
              icon={<svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><path d="M4 15s1-1 4-1 5 2 8 2 4-1 4-1V3s-1 1-4 1-5-2-8-2-4 1-4 1z"/><line x1="4" y1="22" x2="4" y2="15"/></svg>} />
            <ProblemCard delay="delay-2" title="Tidak tahu modal realistis" desc="Hitungan masih feeling. Tidak tahu apakah margin cukup, atau jualan rugi dari hari pertama."
              icon={<svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><line x1="12" y1="1" x2="12" y2="23"/><path d="M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6"/></svg>} />
            <ProblemCard delay="delay-3" title="Takut gagal, takut malu" desc="Jadi makin parah karena tidak ada yang mengawasi. Proyek tinggal proyek, tidak pernah dieksekusi."
              icon={<svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"/></svg>} />
            <ProblemCard delay="delay-4" title="Gonta-ganti ide terus" desc="Ide A sedikit, ide B sedikit. Tidak pernah fokus cukup lama buat mendapatkan data pasar yang benar."
              icon={<svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><polyline points="23 4 23 10 17 10"/><path d="M20.49 15a9 9 0 1 1-2.12-9.36L23 10"/></svg>} />
            <ProblemCard delay="delay-5" title="Semangat palsu, nol bukti" desc="Yang ada cuma semangat dan keyakinan tanpa data. Padahal modal dan tenaga terbatas."
              icon={<svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><path d="M12 20h9"/><path d="M16.5 3.5a2.121 2.121 0 0 1 3 3L7 19l-4 1 1-4L16.5 3.5z"/></svg>} />
          </div>
        </div>
      </section>

      {/* CARA KERJA */}
      <section className="how-it-works" id="cara-kerja">
        <div className="container">
          <SectionIntro label="Cara Kerja" title="14 hari. 5 tahap. Satu keputusan." desc="Tidak perlu ribet. Kami kasih kamu sistem langkah demi langkah." />
          <div className="steps">
            <StepCard num="1" title="Compress Ide" desc="Tulis ide mentah kamu. Sistem akan membantu menyusunnya jadi konsep bisnis yang tajam dan jelas." delay="" />
            <StepCard num="2" title="Fokus Menu" desc="Pilih maksimal 1-3 SKU. Sistem kasih rekomendasi hero product mana yang paling layak diuji duluan." delay="delay-1" />
            <StepCard num="3" title="Hitung Realitas" desc="Input bahan, harga, kemasan. Lihat HPP, margin, dan modal awal yang kamu butuhkan — bukan feeling." delay="delay-2" />
            <StepCard num="4" title="Jalankan Misi" desc="Dapatkan tugas harian: sampling, pre-order, polling, titip jual. Selesai = bukti diupload." delay="delay-3" />
          </div>
          <div style={{ textAlign: 'center', marginTop: 56 }}>
            <div className="step" style={{ display: 'inline-block', maxWidth: 260 }}>
              <div className="step-number" style={{ background: 'linear-gradient(135deg,var(--brand-amber),var(--brand-orange))' }}>5</div>
              <h3>Keputusan Akhir</h3>
              <p>Sistem kasih skor kesiapan + keputusan tegas: <strong>Lanjut, Ulang, Pivot, atau Stop</strong>.</p>
            </div>
          </div>
        </div>
      </section>

      {/* FITUR */}
      <section className="features" id="fitur">
        <div className="container">
          <div className="features-header">
            <SectionIntro label="Fitur" title="Sistem yang memaksa kamu bergerak, bukan cuma baca laporan" desc="Setiap fitur dibangun untuk satu tujuan: mengubah niat jadi bukti." />
          </div>
          <div className="features-grid">
            <FeatureCard tag="AI" title="Idea Compressor" desc="Masukkan ide mentah dalam bahasa sehari-hari. AI akan menyusunnya jadi konsep bisnis, target customer, value proposition, dan asumsi utama." delay=""
              icon={<svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round"><path d="M12 2a4 4 0 0 1 4 4c0 2-2 4-4 4s-4-2-4-4a4 4 0 0 1 4-4z"/><path d="M2 22c0-4 4.5-6 10-6s10 2 10 6"/></svg>} />
            <FeatureCard tag="Fokus" title="Menu Focus Engine" desc="Pilih 1-3 SKU terbaik. Sistem menilai kompleksitas tiap menu dan merekomendasikan hero product yang paling realistis buat diuji." delay="delay-1"
              icon={<svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round"><circle cx="12" cy="12" r="10"/><circle cx="12" cy="12" r="6"/><circle cx="12" cy="12" r="2"/></svg>} />
            <FeatureCard tag="Angka" title="Reality Ledger" desc="Hitung HPP, harga jual, margin kotor, modal awal minimum, dan titik impas. Sistem akan kasih alarm kalau margin kamu tipis atau berbahaya." delay="delay-2"
              icon={<svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round"><line x1="18" y1="20" x2="18" y2="10"/><line x1="12" y1="20" x2="12" y2="4"/><line x1="6" y1="20" x2="6" y2="14"/></svg>} />
            <FeatureCard tag="Aksi" title="Field Missions" desc="Dapatkan misi harian yang konkret: polling, pre-order, sampling, atau wawancara. Tiap misi punya definisi selesai — dan kamu harus buktikan." delay="delay-3"
              icon={<svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round"><polygon points="13 2 3 14 12 14 11 22 21 10 12 10 13 2"/></svg>} />
            <FeatureCard tag="Bukti" title="Evidence Upload" desc="Upload foto produk, screenshot chat, daftar pesanan, atau feedback pelanggan. Sistem menilai apakah bukti valid, lemah, atau gagal total." delay="delay-4"
              icon={<svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round"><path d="M23 19a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h4l2-3h6l2 3h4a2 2 0 0 1 2 2z"/><circle cx="12" cy="13" r="4"/></svg>} />
            <FeatureCard tag="Keputusan" title="Founder Courtroom" desc="Review adversarial dari 3 peran: calon pembeli, operator dapur, dan reviewer. Menguji kelemahan asumsi kamu sebelum lanjut ke tahap berikutnya." delay="delay-5"
              icon={<svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round"><path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"/></svg>} />
          </div>
        </div>
      </section>

      {/* TESTIMONI */}
      <section className="proof" id="testimoni">
        <div className="container">
          <div className="micro" style={{ color: 'var(--brand-orange)', marginBottom: 12 }}>Testimoni</div>
          <h2 className="display-large" style={{ maxWidth: 500, margin: '0 auto' }}>Apa kata mereka yang sudah mencoba?</h2>
          <div className="proof-quote fade-in">
            <blockquote>
              "Awalnya saya pikir ide jualan puding cup saya sudah oke. Ternyata setelah dihitung margin,
              harga jual saya terlalu rendah. BukaDulu tidak hanya memberi tahu, tapi juga memberi misi untuk
              tes harga langsung ke calon pembeli. Hasilnya? Saya berani pivot ke produk yang lebih
              menguntungkan."
            </blockquote>
            <footer>
              <div className="proof-avatar">SA</div>
              <div>
                <div className="proof-author">Sari Anindya</div>
                <div className="proof-role">Founder — Sari's Kitchen</div>
              </div>
            </footer>
          </div>
        </div>
      </section>

      {/* HARGA */}
      <section className="pricing" id="harga">
        <div className="container">
          <SectionIntro label="Harga" title="Mulai gratis. Bayar kalau mau lanjut." desc="Sprint pertama gratis. Tidak perlu kartu kredit. Cukup buktikan idemu dulu." />
          <div className="pricing-grid">
            <div className="pricing-card fade-in">
              <div className="pricing-name">Free Sprint</div>
              <div className="pricing-desc">Coba sistem validasi 14 hari penuh</div>
              <div className="pricing-price">Rp0 <span>/ sprint</span></div>
              <div className="pricing-period">Sprint pertama gratis</div>
              <ul className="pricing-features">
                <li>Idea compression</li>
                <li>Menu focus engine</li>
                <li>Reality ledger</li>
                <li>3 misi awal</li>
                <li>Evidence upload & review</li>
                <li>Launch readiness score</li>
              </ul>
              <Link to="/register" className="btn btn-ghost">Mulai Gratis</Link>
            </div>
            <div className="pricing-card featured fade-in delay-1">
              <div className="pricing-name">Pro Sprint</div>
              <div className="pricing-desc">Validasi penuh dengan fitur lengkap</div>
              <div className="pricing-price">Rp149k <span>/ sprint</span></div>
              <div className="pricing-period">Setiap sprint 14 hari</div>
              <ul className="pricing-features">
                <li>Semua fitur Free Sprint</li>
                <li>Misi harian tidak terbatas</li>
                <li>Founder Courtroom penuh</li>
                <li>Multiple evidence per misi</li>
                <li>Prioritas review evidence</li>
                <li>Export PDF validasi</li>
              </ul>
              <Link to="/register" className="btn btn-primary">Langganan Sekarang</Link>
            </div>
          </div>
          <p className="caption" style={{ textAlign: 'center', color: 'rgba(255,255,255,0.4)', marginTop: 24 }}>
            Sudah punya akun? <Link to="/login" style={{ color: 'var(--brand-amber)' }}>Masuk di sini</Link>
          </p>
        </div>
      </section>

      <FaqSection items={faqItems} />

      {/* FINAL CTA */}
      <section className="final-cta" id="coba">
        <div className="container">
          <h2>Berhenti memikirkan. Mulai buktikan.</h2>
          <p>Ribuan calon pebisnis berhenti di fase kepikiran. Kamu tidak harus seperti mereka.</p>
          <div className="hero-cta">
            <Link to="/register" className="btn btn-lg btn-white">Cek Ide Saya Gratis</Link>
            <a href="#cara-kerja" className="btn btn-outline-light btn-lg">Pelajari Dulu</a>
          </div>
          <p className="caption" style={{ color: 'rgba(255,255,255,0.5)', marginTop: 20 }}>Gratis. Tidak perlu kartu kredit. Sprint pertama full akses.</p>
        </div>
      </section>

      <footer>
        <div className="container">
          <p>&copy; 2026 BukaDulu — execution system for pre-launch F&B validation. Built with honesty and evidence. <Link to="/login" style={{ color: 'rgba(255,255,255,0.4)', textDecoration: 'none' }}>Masuk</Link></p>
        </div>
      </footer>
    </div>
  )
}
