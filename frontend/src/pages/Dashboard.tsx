import { useEffect, useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import AppNavbar from '../components/AppNavbar'
import LottiePlayer from '../components/LottiePlayer'
import { getHistory, HistoryItem } from '../api/idea'
import { getQuota, QuotaInfo } from '../api/auth'

export default function Dashboard() {
  const navigate = useNavigate()
  const [items, setItems] = useState<HistoryItem[]>([])
  const [quota, setQuota] = useState<QuotaInfo | null>(null)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')

  const load = async () => {
    setLoading(true)
    setError('')
    try {
      const [hist, q] = await Promise.all([getHistory(), getQuota()])
      setItems(hist)
      setQuota(q)
    } catch (err: any) {
      setError(err.message || 'Gagal memuat data.')
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => { load() }, [])

  const validatedCount = items.filter(i => i.idea?.status === 'done').length
  const freeLeft = quota ? Math.max(0, quota.free_limit - quota.free_used) : 0

  return (
    <>
      <AppNavbar />
      <div className="container app-content">
        <div className="dashboard-hero">
          <div>
            <div className="micro">Dasbor</div>
            <h1 className="display-large">Selamat datang kembali 👋</h1>
            <p className="body-large" style={{ marginTop: 12, maxWidth: 560 }}>
              {quota
                ? (freeLeft > 0
                    ? <>Kamu punya <strong>{freeLeft} validasi gratis</strong> tersisa. Ide berikutnya: Rp{quota.price.toLocaleString('id-ID')}.</>
                    : <>Validasi ide berikutnya: <strong>Rp{quota.price.toLocaleString('id-ID')}</strong> per ide.</>)
                : 'Mulai validasi ide F&B kamu sekarang.'}
            </p>
            <button className="btn btn-primary btn-lg" style={{ marginTop: 20 }} onClick={() => navigate('/idea/new')}>
              + Validasi Ide Baru
            </button>
          </div>
          <div className="dashboard-hero-lottie">
            <LottiePlayer src="/lottie/Creative-Idea.json" style={{ width: '100%', maxWidth: 320 }} />
          </div>
        </div>

        <section style={{ marginTop: 48 }}>
          <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'baseline', marginBottom: 16 }}>
            <h2 className="display-large" style={{ fontSize: '1.75rem' }}>Riwayat Validasi</h2>
            <Link to="/history" className="btn btn-ghost">Lihat semua</Link>
          </div>

          {loading ? (
            <p className="body-text">Memuat...</p>
          ) : error ? (
            <div className="auth-error">{error}</div>
          ) : items.length === 0 ? (
            <div className="empty-state">
              <LottiePlayer src="/lottie/Creative-Idea.json" style={{ width: 200, margin: '0 auto' }} loop={false} />
              <h3 style={{ marginTop: 16 }}>Belum ada ide</h3>
              <p className="body-text">Mulai ide pertama kamu — gratis.</p>
              <button className="btn btn-primary" style={{ marginTop: 12 }} onClick={() => navigate('/idea/new')}>
                Validasi Ide Pertama
              </button>
            </div>
          ) : (
            <div className="history-list">
              {items.slice(0, 5).map(item => (
                <Link to={`/idea/${item.venture.id}`} key={item.venture.id} className="history-card">
                  <div className="history-card-main">
                    <div className="history-card-title">{item.idea?.one_line_concept || item.venture.name}</div>
                    <div className="history-card-meta">
                      {new Date(item.venture.updated_at).toLocaleDateString('id-ID', { day: 'numeric', month: 'short', year: 'numeric' })}
                      {' · '}
                      <span className={`status-badge status-${item.idea?.status || 'pending'}`}>
                        {item.idea?.status === 'done' ? 'Selesai' : item.idea?.status === 'processing' ? 'Memproses' : item.idea?.status === 'failed' ? 'Gagal' : 'Draft'}
                      </span>
                    </div>
                  </div>
                  <span className="history-card-arrow">→</span>
                </Link>
              ))}
            </div>
          )}
          {validatedCount > 0 && (
            <p className="caption" style={{ marginTop: 16 }}>Total {validatedCount} ide sudah divalidasi.</p>
          )}
        </section>
      </div>
    </>
  )
}
