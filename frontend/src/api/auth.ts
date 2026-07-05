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
  const text = await res.text().catch(() => '')
  // Coba parse JSON; kalau bukan JSON (mis. "404 page not found"), pakai teks apa adanya
  try {
    const data = JSON.parse(text)
    if (data?.error && typeof data.error === 'object') return data.error.message || 'Permintaan gagal'
    return data?.message || data?.error || 'Permintaan gagal'
  } catch {
    return text?.trim() || `Permintaan gagal (HTTP ${res.status})`
  }
}

// Helper aman baca body sukses: kalau bukan JSON, throw dengan pesan jelas.
async function parseJson(res: Response, fallbackMsg: string): Promise<any> {
  const text = await res.text().catch(() => '')
  try {
    return JSON.parse(text)
  } catch {
    throw new Error(text?.trim() ? `${fallbackMsg}: ${text.trim().slice(0, 200)}` : fallbackMsg)
  }
}

export async function login(emailOrPhone: string, password: string) {
  const res = await fetch(`${API_BASE}/auth/login`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email_or_phone: emailOrPhone, password }),
  })
  if (!res.ok) throw new Error(await parseError(res))
  return parseJson(res, 'Login gagal')
}

export async function register(fullName: string, email: string, password: string) {
  const res = await fetch(`${API_BASE}/auth/register`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ full_name: fullName, email, password }),
  })
  if (!res.ok) throw new Error(await parseError(res))
  return parseJson(res, 'Registrasi gagal')
}

export async function verifyOTP(email: string, otp: string) {
  const res = await fetch(`${API_BASE}/auth/verify-otp`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email, otp }),
  })
  if (!res.ok) throw new Error(await parseError(res))
  return parseJson(res, 'Verifikasi OTP gagal')
}

export async function resendOTP(email: string) {
  const res = await fetch(`${API_BASE}/auth/resend-otp`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email }),
  })
  if (!res.ok) throw new Error(await parseError(res))
  return parseJson(res, 'Kirim ulang OTP gagal')
}

export async function loginWithGoogle(code: string) {
  const res = await fetch(`${API_BASE}/auth/google`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ code }),
  })
  if (!res.ok) throw new Error(await parseError(res))
  return parseJson(res, 'Login Google gagal')
}

export async function getQuota(): Promise<QuotaInfo> {
  const res = await fetch(`${API_BASE}/auth/quota`, { headers: authHeader() })
  if (!res.ok) throw new Error(await parseError(res))
  return res.json()
}
