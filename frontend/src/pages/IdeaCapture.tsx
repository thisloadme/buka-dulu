import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import AppNavbar from '../components/AppNavbar'
import LottiePlayer from '../components/LottiePlayer'
import { createVenture, captureIdea } from '../api/idea'

const MIN_CHARS = 20

export default function IdeaCapture() {
  const navigate = useNavigate()
  const [name, setName] = useState('')
  const [text, setText] = useState('')
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')

  const counterColor = text.length < MIN_CHARS ? 'var(--danger)' : 'var(--success)'

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setError('')
    if (text.trim().length < MIN_CHARS) {
      setError(`Minimal ${MIN_CHARS} karakter untuk ide.`)
      return
    }
    setLoading(true)
    try {
      const ventureName = name.trim() || text.trim().slice(0, 40)
      const venture = await createVenture(ventureName)
      await captureIdea(venture.id, text.trim())
      navigate(`/idea/${venture.id}`)
    } catch (err: any) {
      setError(err.message || 'Gagal menyimpan ide.')
    } finally {
      setLoading(false)
    }
  }

  return (
    <>
      <AppNavbar />
      <div className="container app-content">
        <div className="page-grid">
          <div>
            <div className="micro">Tahap 1</div>
            <h1 className="display-large" style={{ fontSize: '2.25rem' }}>Ceritakan Idemu</h1>
            <p className="body-large" style={{ marginTop: 12 }}>
              Tulis ide bisnis F&B kamu dengan detail. Sebutkan produk, target pelanggan, harga, dan lokasi.
              Semakin detail, semakin tajam analisis AI.
            </p>

            <form onSubmit={handleSubmit} style={{ marginTop: 24 }}>
              <div className="form-group">
                <label htmlFor="name">Nama Usaha / Ide (opsional)</label>
                <input id="name" type="text" value={name}
                  onChange={e => setName(e.target.value)}
                  placeholder="Contoh: Nasi Geprek Bunda" />
              </div>
              <div className="form-group">
                <label htmlFor="text">Deskripsi Ide</label>
                <textarea id="text" value={text} onChange={e => setText(e.target.value)}
                  rows={8} minLength={MIN_CHARS} required
                  placeholder="Contoh: Saya ingin jual nasi goreng homemade dengan topping ayam geprek, target mahasiswa kampus di area Jogja, harga 15-20 ribu, buka sore-malam..." />
              </div>
              <div style={{ display: 'flex', justifyContent: 'space-between', marginTop: 4 }}>
                <span className="caption">Minimal {MIN_CHARS} karakter</span>
                <span className="caption" style={{ color: counterColor, fontWeight: 600 }}>
                  {text.length} / {MIN_CHARS}
                </span>
              </div>

              {error && <div className="auth-error" style={{ marginTop: 16 }}>{error}</div>}

              <button type="submit" className="btn btn-primary btn-lg btn-block" style={{ marginTop: 20 }} disabled={loading}>
                {loading ? 'Menyimpan...' : 'Lanjut ke Analisis AI'}
              </button>
            </form>
          </div>

          <div className="page-side">
            <LottiePlayer src="/lottie/Creative-Idea.json" style={{ width: '100%', maxWidth: 280, margin: '0 auto' }} />
            <div className="tip-card" style={{ marginTop: 16 }}>
              <strong>Tips:</strong> Sebutkan siapa pelanggan spesifik, momen konsumsi, dan keunikan produkmu.
              Hindari deskripsi terlalu umum seperti "jualan makanan".
            </div>
          </div>
        </div>
      </div>
    </>
  )
}
