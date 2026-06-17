import { useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { login } from '../api/auth'

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
      localStorage.setItem('bukadulu_token', data.token)
      navigate('/download')
    } catch (err: any) {
      setError(err.message || 'Login gagal. Coba lagi.')
    } finally {
      setLoading(false)
    }
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
