import { useEffect, useState } from 'react'
import { Link } from 'react-router-dom'
import AppNavbar from '../components/AppNavbar'
import LottiePlayer from '../components/LottiePlayer'
import { getHistory, HistoryItem } from '../api/idea'

export default function History() {
  const [items, setItems] = useState<HistoryItem[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')

  useEffect(() => {
    getHistory()
      .then(setItems)
      .catch(err => setError(err.message || 'Gagal memuat riwayat.'))
      .finally(() => setLoading(false))
  }, [])

  const statusLabel = (s?: string) => {
    switch (s) {
      case 'done': return 'Selesai'
      case 'processing': return 'Memproses'
      case 'failed': return 'Gagal'
      default: return 'Draft'
    }
  }

  return (
    <>
      <AppNavbar />
      <div className="container app-content">
        <h1 className="display-large" style={{ fontSize: '2rem' }}>Riwayat Validasi</h1>
        <p className="body-text" style={{ marginTop: 8 }}>Semua ide yang pernah kamu validasi.</p>

        {loading ? (
          <p className="body-text" style={{ marginTop: 24 }}>Memuat...</p>
        ) : error ? (
          <div className="auth-error" style={{ marginTop: 16 }}>{error}</div>
        ) : items.length === 0 ? (
          <div className="empty-state" style={{ marginTop: 24 }}>
            <LottiePlayer src="/lottie/meditation.json" style={{ width: 200, margin: '0 auto' }} />
            <h3 style={{ marginTop: 16 }}>Belum ada riwayat</h3>
            <p className="body-text">Validasi ide pertamamu untuk mulai.</p>
            <Link to="/idea/new" className="btn btn-primary" style={{ marginTop: 12 }}>+ Validasi Ide Baru</Link>
          </div>
        ) : (
          <div className="history-list" style={{ marginTop: 24 }}>
            {items.map(item => (
              <Link to={`/idea/${item.venture.id}`} key={item.venture.id} className="history-card">
                <div className="history-card-main">
                  <div className="history-card-title">{item.idea?.one_line_concept || item.venture.name}</div>
                  <div className="history-card-sub">{item.idea?.target_customer || item.idea?.raw_input?.slice(0, 80)}</div>
                  <div className="history-card-meta">
                    {new Date(item.venture.updated_at).toLocaleDateString('id-ID', { day: 'numeric', month: 'long', year: 'numeric' })}
                    {' · '}
                    <span className={`status-badge status-${item.idea?.status || 'pending'}`}>
                      {statusLabel(item.idea?.status)}
                    </span>
                  </div>
                </div>
                <span className="history-card-arrow">→</span>
              </Link>
            ))}
          </div>
        )}
      </div>
    </>
  )
}
