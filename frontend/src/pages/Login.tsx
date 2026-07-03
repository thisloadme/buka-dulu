import { useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { login, googleAuthUrl, GOOGLE_CLIENT_ID } from '../api/auth'
import { setToken } from '../lib/token'

export default function Login() {
  const navigate = useNavigate()
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setError('')
    setLoading(true)
    try {
      const data = await login(email, password)
      if (data.token) {
        setToken(data.token)
        navigate('/dashboard')
      } else if (data.user?.email) {
        // Account not verified yet — go to OTP page
        navigate(`/register?verify=${encodeURIComponent(data.user.email)}`)
      }
    } catch (err: any) {
      setError(err.message || 'Login gagal. Coba lagi.')
    } finally {
      setLoading(false)
    }
  }

  const handleGoogle = () => {
    window.location.href = googleAuthUrl()
  }

  return (
    <div className="auth-page">
      <div className="auth-container">
        <Link to="/" className="auth-logo">
          <img src="/bukadulu.png" alt="BukaDulu" width="40" height="40" />
          <span>BukaDulu</span>
        </Link>
        <h1>Masuk</h1>
        <p className="auth-subtitle">Lanjutkan validasi ide F&B kamu</p>

        {GOOGLE_CLIENT_ID && (
          <>
            <button type="button" className="btn btn-google btn-block" onClick={handleGoogle}>
              <svg width="18" height="18" viewBox="0 0 18 18"><path fill="#4285F4" d="M17.64 9.2c0-.64-.06-1.25-.17-1.84H9v3.48h4.84a4.14 4.14 0 0 1-1.8 2.71v2.26h2.92c1.7-1.57 2.68-3.88 2.68-6.61z"/><path fill="#34A853" d="M9 18c2.43 0 4.47-.81 5.96-2.18l-2.92-2.26c-.81.54-1.84.86-3.04.86-2.34 0-4.32-1.58-5.03-3.71H.96v2.33A9 9 0 0 0 9 18z"/><path fill="#FBBC05" d="M3.97 10.71A5.4 5.4 0 0 1 3.68 9c0-.59.1-1.17.29-1.71V4.96H.96A9 9 0 0 0 0 9c0 1.45.35 2.83.96 4.04l3.01-2.33z"/><path fill="#EA4335" d="M9 3.58c1.32 0 2.5.45 3.44 1.35l2.58-2.58A9 9 0 0 0 .96 4.96l3.01 2.33C4.68 5.16 6.66 3.58 9 3.58z"/></svg>
              Masuk dengan Google
            </button>
            <div className="auth-divider"><span>atau</span></div>
          </>
        )}

        <form onSubmit={handleSubmit}>
          <div className="form-group">
            <label htmlFor="email">Email atau Nomor HP</label>
            <input id="email" type="text" value={email}
              onChange={e => setEmail(e.target.value)} placeholder="email@contoh.com"
              required autoFocus />
          </div>
          <div className="form-group">
            <label htmlFor="password">Password</label>
            <input id="password" type="password" value={password}
              onChange={e => setPassword(e.target.value)} placeholder="Minimal 8 karakter"
              required minLength={8} />
          </div>

          {error && <div className="auth-error">{error}</div>}

          <button type="submit" className="btn btn-primary btn-block" disabled={loading}>
            {loading ? 'Memproses...' : 'Masuk'}
          </button>
        </form>

        <p className="auth-footer">
          Belum punya akun? <Link to="/register">Daftar</Link>
        </p>
      </div>
    </div>
  )
}
