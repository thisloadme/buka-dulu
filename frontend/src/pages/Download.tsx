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
          <a href="/apk/bukadulu.apk" className="btn btn-primary btn-lg btn-block"
            style={{ justifyContent: 'center', marginBottom: 12 }}>
            <svg width="20" height="20" viewBox="0 0 24 24" fill="currentColor">
              <path d="M3.609 1.814L13.792 12 3.61 22.186a.996.996 0 0 1-.61-.92V2.734a1 1 0 0 1 .609-.92zm10.89 10.893l2.302 2.302-10.937 6.333 8.635-8.635zm3.199-3.199l2.807 1.626a1 1 0 0 1 0 1.732l-2.807 1.626L15.206 12l2.492-2.492zM5.864 2.658L16.8 8.99l-2.302 2.302-8.634-8.634z"/>
            </svg>
            Download APK Android
          </a>
          <p className="caption" style={{ color: 'var(--body)', marginBottom: 24 }}>
            Versi iOS segera hadir. Untuk sementara, gunakan perangkat Android.
          </p>
        </div>

        <div className="download-steps">
          <h3>Setelah download:</h3>
          <ol>
            <li>Buka file APK yang sudah diunduh</li>
            <li>Izinkan instalasi dari sumber tidak dikenal (jika diminta)</li>
            <li>Buka aplikasi dan masuk dengan akun kamu</li>
          </ol>
        </div>

        <button onClick={handleLogout} className="btn btn-ghost" style={{ marginTop: 32 }}>
          Keluar
        </button>
      </div>
    </div>
  )
}
