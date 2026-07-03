const TOKEN_KEY = 'bukadulu_token'

export function getToken(): string | null {
  return localStorage.getItem(TOKEN_KEY)
}

export function setToken(token: string): void {
  localStorage.setItem(TOKEN_KEY, token)
}

export function clearToken(): void {
  localStorage.removeItem(TOKEN_KEY)
}

export function isLoggedIn(): boolean {
  return !!getToken()
}

// authHeader returns fetch headers with the bearer token if present.
export function authHeader(extra: Record<string, string> = {}): Record<string, string> {
  const headers: Record<string, string> = { 'Content-Type': 'application/json', ...extra }
  const token = getToken()
  if (token) headers['Authorization'] = `Bearer ${token}`
  return headers
}
