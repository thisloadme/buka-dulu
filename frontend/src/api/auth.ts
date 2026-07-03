import { authHeader } from '../lib/token'

const API_BASE = import.meta.env.VITE_API_URL || 'http://localhost:8080/api/v1'

// Google OAuth Client ID (public, safe to expose). Set via Vite env.
export const GOOGLE_CLIENT_ID = import.meta.env.VITE_GOOGLE_CLIENT_ID || ''
export const GOOGLE_REDIRECT_URL = `${window.location.origin}/auth/google/callback`

export function googleAuthUrl(): string {
  const params = new URLSearchParams({
    client_id: GOOGLE_CLIENT_ID,
    redirect_uri: GOOGLE_REDIRECT_URL,
    response_type: 'code',
    scope: 'openid email profile',
    prompt: 'select_account',
  })
  return `https://accounts.google.com/o/oauth2/v2/auth?${params.toString()}`
}

export interface QuotaInfo {
  free_used: number
  free_limit: number
  price: number
}

async function parseError(res: Response): Promise<string> {
  try {
    const data = await res.json()
    if (data?.error && typeof data.error === 'object') return data.error.message || 'Permintaan gagal'
    return data?.message || data?.error || 'Permintaan gagal'
  } catch {
    return 'Permintaan gagal'
  }
}

export async function login(emailOrPhone: string, password: string) {
  const res = await fetch(`${API_BASE}/auth/login`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email_or_phone: emailOrPhone, password }),
  })
  const data = await res.json()
  if (!res.ok) throw new Error(data.message || data.error || 'Login gagal')
  return data
}

export async function register(fullName: string, email: string, password: string) {
  const res = await fetch(`${API_BASE}/auth/register`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ full_name: fullName, email, password }),
  })
  const data = await res.json()
  if (!res.ok) throw new Error(data.message || data.error || 'Registrasi gagal')
  return data
}

export async function verifyOTP(email: string, otp: string) {
  const res = await fetch(`${API_BASE}/auth/verify-otp`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email, otp }),
  })
  const data = await res.json()
  if (!res.ok) throw new Error(data.message || data.error || 'Verifikasi OTP gagal')
  return data
}

export async function resendOTP(email: string) {
  const res = await fetch(`${API_BASE}/auth/resend-otp`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email }),
  })
  const data = await res.json()
  if (!res.ok) throw new Error(data.message || data.error || 'Kirim ulang OTP gagal')
  return data
}

export async function loginWithGoogle(code: string) {
  const res = await fetch(`${API_BASE}/auth/google`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ code }),
  })
  const data = await res.json()
  if (!res.ok) throw new Error(data.message || data.error || 'Login Google gagal')
  return data
}

export async function getQuota(): Promise<QuotaInfo> {
  const res = await fetch(`${API_BASE}/auth/quota`, { headers: authHeader() })
  if (!res.ok) throw new Error(await parseError(res))
  return res.json()
}
