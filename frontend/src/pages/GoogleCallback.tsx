import { useEffect, useState } from 'react'
import { useNavigate, useSearchParams } from 'react-router-dom'
import { loginWithGoogle } from '../api/auth'
import { setToken } from '../lib/token'

export default function GoogleCallback() {
  const navigate = useNavigate()
  const [params] = useSearchParams()
  const [error, setError] = useState('')

  useEffect(() => {
    const code = params.get('code')
    if (!code) {
      setError('Kode otorisasi tidak ditemukan.')
      return
    }
    loginWithGoogle(code)
      .then(data => {
        if (data?.token) {
          setToken(data.token)
          navigate('/dashboard')
        } else {
          setError('Login Google gagal: token tidak diterima.')
        }
      })
      .catch(err => setError(err.message || 'Login Google gagal.'))
  }, [params, navigate])

  return (
    <div className="auth-page">
      <div className="auth-container" style={{ textAlign: 'center' }}>
        {error ? (
          <>
            <h1>Login Gagal</h1>
            <p className="auth-error" style={{ marginTop: 16 }}>{error}</p>
            <a href="/login" className="btn btn-primary" style={{ marginTop: 16 }}>Kembali ke Login</a>
          </>
        ) : (
          <>
            <h1>Memproses login Google...</h1>
            <p className="auth-subtitle">Mohon tunggu</p>
          </>
        )}
      </div>
    </div>
  )
}
