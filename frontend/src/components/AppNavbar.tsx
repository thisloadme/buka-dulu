import { Link, useNavigate } from 'react-router-dom'
import { clearToken } from '../lib/token'

export default function AppNavbar() {
  const navigate = useNavigate()

  const handleLogout = () => {
    clearToken()
    navigate('/login')
  }

  return (
    <nav className="app-nav">
      <div className="container app-nav-inner">
        <Link to="/dashboard" className="logo">
          <img src="/bukadulu.png" alt="BukaDulu" width="28" height="28" />
          BukaDulu
        </Link>
        <div className="app-nav-links">
          <Link to="/dashboard">Beranda</Link>
          <Link to="/history">Riwayat</Link>
          <button className="btn btn-ghost" onClick={handleLogout}>Keluar</button>
        </div>
      </div>
    </nav>
  )
}
