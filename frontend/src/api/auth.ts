const API_BASE = import.meta.env.VITE_API_URL || 'http://localhost:8080/api/v1'

export async function login(emailOrPhone: string, password: string) {
  const res = await fetch(`${API_BASE}/auth/login`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email_or_phone: emailOrPhone, password }),
  })
  const data = await res.json()
  if (!res.ok) throw new Error(data.message || 'Login gagal')
  return data
}

export async function register(fullName: string, email: string, password: string) {
  const res = await fetch(`${API_BASE}/auth/register`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ full_name: fullName, email, password }),
  })
  const data = await res.json()
  if (!res.ok) throw new Error(data.message || 'Registrasi gagal')
  return data
}
