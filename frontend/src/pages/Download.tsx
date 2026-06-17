import { useNavigate } from 'react-router-dom'
import { useEffect, useState } from 'react'

export default function Download() {
  const navigate = useNavigate()
  const [token, setToken] = useState<string | null>(null)

  useEffect(() => {
    const t = localStorage.getItem('bukadulu_token')
    if (!t) {
      navigate('/login')
      return
    }
    setToken(t)
  }, [navigate])

  const handleLogout = () => {
    localStorage.removeItem('bukadulu_token')
    navigate('/')
  }

  if (!token) return null

  return (
    <div className="auth-page">
      <div className="auth-container" style={{ textAlign: 'center' }}>
        <div className="download-icon">
          <svg width="64" height="64" viewBox="0 0 24 24" fill="none" stroke="var(--brand-orange)" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round">
            <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/>
            <polyline points="7 10 12 15 17 10"/>
            <line x1="12" y1="15" x2="12" y2="3"/>
          </svg>
        </div>
        <h1 style={{ marginTop: 20 }}>Selamat Datang!</h1>
        <p className="auth-subtitle">
          Akun kamu berhasil dibuat. Sekarang unduh aplikasi BukaDulu untuk memulai validasi ide F&B kamu.
        </p>

        <div className="download-buttons">
          <button className="btn btn-primary btn-lg btn-block"
            disabled
            style={{ justifyContent: 'center', marginBottom: 12, opacity: 0.6, cursor: 'not-allowed' }}>
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
              <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/>
              <polyline points="7 10 12 15 17 10"/>
              <line x1="12" y1="15" x2="12" y2="3"/>
            </svg>
            Coming Soon
          </button>
          <p className="caption" style={{ color: 'var(--body)', marginBottom: 24 }}>
            Aplikasi Android masih dalam pengembangan. Pantau terus untuk update selanjutnya.
          </p>
        </div>

        <button onClick={handleLogout} className="btn btn-ghost" style={{ marginTop: 32 }}>
          Keluar
        </button>
      </div>
    </div>
  )
}
