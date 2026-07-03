import { useState } from 'react'
import { Link, useNavigate, useSearchParams } from 'react-router-dom'
import { register, verifyOTP, resendOTP, googleAuthUrl } from '../api/auth'
import { setToken } from '../lib/token'

export default function Register() {
  const navigate = useNavigate()
  const [params] = useSearchParams()
  const [name, setName] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [otp, setOtp] = useState('')
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')
  const [info, setInfo] = useState('')

  // ?verify=email → land directly in OTP verification mode (from Login for unverified accounts)
  const [verifyEmail, setVerifyEmail] = useState(params.get('verify') || '')

  const handleRegister = async (e: React.FormEvent) => {
    e.preventDefault()
    if (password.length < 8) {
      setError('Password minimal 8 karakter')
      return
    }
    setError('')
    setLoading(true)
    try {
      await register(name, email, password)
      setVerifyEmail(email)
      setInfo('Kode OTP telah dikirim ke email kamu. Masukkan untuk verifikasi.')
    } catch (err: any) {
      setError(err.message || 'Registrasi gagal. Coba lagi.')
    } finally {
      setLoading(false)
    }
  }

  const handleVerify = async (e: React.FormEvent) => {
    e.preventDefault()
    setError('')
    setLoading(true)
    try {
      const data = await verifyOTP(verifyEmail, otp)
      if (data.token) {
        setToken(data.token)
        navigate('/dashboard')
      } else {
        setError('Verifikasi gagal. Coba lagi.')
      }
    } catch (err: any) {
      setError(err.message || 'Verifikasi gagal.')
    } finally {
      setLoading(false)
    }
  }

  const handleResend = async () => {
    setError('')
    try {
      await resendOTP(verifyEmail)
      setInfo('Kode OTP baru telah dikirim.')
    } catch (err: any) {
      setError(err.message || 'Gagal mengirim ulang OTP.')
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

        {verifyEmail ? (
          <>
            <h1>Verifikasi Email</h1>
            <p className="auth-subtitle">Masukkan kode OTP yang dikirim ke {verifyEmail}</p>
            <form onSubmit={handleVerify}>
              <div className="form-group">
                <label htmlFor="otp">Kode OTP</label>
                <input id="otp" type="text" inputMode="numeric" pattern="[0-9]*" maxLength={6}
                  value={otp} onChange={e => setOtp(e.target.value.replace(/\D/g, ''))}
                  placeholder="6 digit" required autoFocus />
              </div>
              {error && <div className="auth-error">{error}</div>}
              {info && <div className="auth-info">{info}</div>}
              <button type="submit" className="btn btn-primary btn-block" disabled={loading}>
                {loading ? 'Memproses...' : 'Verifikasi'}
              </button>
            </form>
            <button className="btn btn-ghost btn-block" style={{ marginTop: 8 }} onClick={handleResend}>
              Kirim ulang OTP
            </button>
          </>
        ) : (
          <>
            <h1>Daftar</h1>
            <p className="auth-subtitle">Mulai validasi ide F&B kamu dalam 14 hari</p>

            <button type="button" className="btn btn-google btn-block" onClick={handleGoogle}>
              <svg width="18" height="18" viewBox="0 0 18 18"><path fill="#4285F4" d="M17.64 9.2c0-.64-.06-1.25-.17-1.84H9v3.48h4.84a4.14 4.14 0 0 1-1.8 2.71v2.26h2.92c1.7-1.57 2.68-3.88 2.68-6.61z"/><path fill="#34A853" d="M9 18c2.43 0 4.47-.81 5.96-2.18l-2.92-2.26c-.81.54-1.84.86-3.04.86-2.34 0-4.32-1.58-5.03-3.71H.96v2.33A9 9 0 0 0 9 18z"/><path fill="#FBBC05" d="M3.97 10.71A5.4 5.4 0 0 1 3.68 9c0-.59.1-1.17.29-1.71V4.96H.96A9 9 0 0 0 0 9c0 1.45.35 2.83.96 4.04l3.01-2.33z"/><path fill="#EA4335" d="M9 3.58c1.32 0 2.5.45 3.44 1.35l2.58-2.58A9 9 0 0 0 .96 4.96l3.01 2.33C4.68 5.16 6.66 3.58 9 3.58z"/></svg>
              Daftar dengan Google
            </button>
            <div className="auth-divider"><span>atau</span></div>

            <form onSubmit={handleRegister}>
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
          </>
        )}
      </div>
    </div>
  )
}
