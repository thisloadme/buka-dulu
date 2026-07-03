import { authHeader } from '../lib/token'

const API_BASE = import.meta.env.VITE_API_URL || 'http://localhost:8080/api/v1'

export interface Idea {
  id: string
  venture_id: string
  raw_input: string
  one_line_concept?: string
  target_customer?: string
  value_proposition?: string
  key_assumptions?: string
  early_risks?: string
  version: number
  is_locked: boolean
  status: string
  created_at: string
  updated_at: string
}

export interface Venture {
  id: string
  owner_user_id: string
  name: string
  category?: string
  region?: string
  stage: string
  current_version: number
  created_at: string
  updated_at: string
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

// Create a new venture. Returns the venture object.
export async function createVenture(name: string, category?: string, region?: string): Promise<Venture> {
  const res = await fetch(`${API_BASE}/ventures`, {
    method: 'POST',
    headers: authHeader(),
    body: JSON.stringify({ name, category, region }),
  })
  if (!res.ok) throw new Error(await parseError(res))
  return res.json()
}

export async function listVentures(): Promise<Venture[]> {
  const res = await fetch(`${API_BASE}/ventures`, { headers: authHeader() })
  if (!res.ok) throw new Error(await parseError(res))
  const data = await res.json()
  // backend may return array or {ventures:[...]}
  return Array.isArray(data) ? data : (data.ventures ?? [])
}

export async function captureIdea(ventureId: string, rawInput: string): Promise<Idea> {
  const res = await fetch(`${API_BASE}/ventures/${ventureId}/idea`, {
    method: 'POST',
    headers: authHeader(),
    body: JSON.stringify({ raw_input: rawInput }),
  })
  if (!res.ok) throw new Error(await parseError(res))
  return res.json()
}

export async function getIdea(ventureId: string): Promise<Idea | null> {
  const res = await fetch(`${API_BASE}/ventures/${ventureId}/idea`, { headers: authHeader() })
  if (res.status === 404) return null
  if (!res.ok) throw new Error(await parseError(res))
  return res.json()
}

export interface ProcessIdeaResult {
  ok: boolean
  idea?: Idea
  paymentRequired?: boolean
  ventureId?: string
  purpose?: string
}

// Process (run AI on) the venture's idea. Returns paymentRequired flag on 402.
export async function processIdea(ventureId: string): Promise<ProcessIdeaResult> {
  const res = await fetch(`${API_BASE}/ventures/${ventureId}/idea/process`, {
    method: 'POST',
    headers: authHeader(),
  })
  if (res.status === 402) {
    const data = await res.json().catch(() => ({}))
    return { ok: false, paymentRequired: true, ventureId, purpose: data.purpose || 'idea_validation' }
  }
  if (!res.ok) throw new Error(await parseError(res))
  const idea = await res.json()
  return { ok: true, idea }
}

export async function confirmIdea(ventureId: string): Promise<Idea> {
  const res = await fetch(`${API_BASE}/ventures/${ventureId}/idea/confirm`, {
    method: 'POST',
    headers: authHeader(),
  })
  if (!res.ok) throw new Error(await parseError(res))
  return res.json()
}

export async function updateIdea(ventureId: string, patch: Partial<Idea>): Promise<Idea> {
  const res = await fetch(`${API_BASE}/ventures/${ventureId}/idea`, {
    method: 'PUT',
    headers: authHeader(),
    body: JSON.stringify(patch),
  })
  if (!res.ok) throw new Error(await parseError(res))
  return res.json()
}

export interface HistoryItem {
  venture: Venture
  idea: Idea | null
}

export async function getHistory(): Promise<HistoryItem[]> {
  const res = await fetch(`${API_BASE}/history`, { headers: authHeader() })
  if (!res.ok) throw new Error(await parseError(res))
  const data = await res.json()
  return data.items ?? []
}
