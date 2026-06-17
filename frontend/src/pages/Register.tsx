import { useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { register } from '../api/auth'

export default function Register() {
  const navigate = useNavigate()
  const [name, setName] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    if (password.length < 8) {
      setError('Password minimal 8 karakter')
      return
    }
    setError('')
    setLoading(true)
    try {
      const data = await register(name, email, password)
      localStorage.setItem('bukadulu_token', data.token)
      navigate('/download')
    } catch (err: any) {
      setError(err.message || 'Registrasi gagal. Coba lagi.')
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
        <h1>Daftar</h1>
        <p className="auth-subtitle">Mulai validasi ide F&B kamu dalam 14 hari</p>

        <form onSubmit={handleSubmit}>
          <div className="form-group">
            <label htmlFor="name">Nama Lengkap</label>
            <input id="name" type="text" value={name}
              onChange={e => setName(e.target.value)} placeholder="Nama kamu"
              required autoFocus />
          </div>
          <div className="form-group">
            <label htmlFor="email">Email</label>
            <input id="email" type="email" value={email}
              onChange={e => setEmail(e.target.value)} placeholder="email@contoh.com"
              required />
          </div>
          <div className="form-group">
            <label htmlFor="password">Password</label>
            <input id="password" type="password" value={password}
              onChange={e => setPassword(e.target.value)} placeholder="Minimal 8 karakter"
              required minLength={8} />
          </div>

          {error && <div className="auth-error">{error}</div>}

          <button type="submit" className="btn btn-primary btn-block" disabled={loading}>
            {loading ? 'Memproses...' : 'Daftar'}
          </button>
        </form>

        <p className="auth-footer">
          Sudah punya akun? <Link to="/login">Masuk</Link>
        </p>
      </div>
    </div>
  )
}
