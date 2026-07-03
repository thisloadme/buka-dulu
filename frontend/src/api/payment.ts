import { authHeader } from '../lib/token'

const API_BASE = import.meta.env.VITE_API_URL || 'http://localhost:8080/api/v1'

export interface Order {
  id: string
  user_id: string
  venture_id?: string
  purpose: string
  amount: number
  total_amount?: string
  status: string // pending|paid|expired|failed
  qris_url?: string
  qris_image?: string
  expired_at?: string
  paid_at?: string
  fulfilled: boolean
  created_at: string
}

export interface CreateOrderResult {
  order_id: string
  free: boolean
  qris_image?: string
  total_amount?: string
  expired_at?: string
  status: string
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

export async function createOrder(ventureId: string, purpose = 'idea_validation'): Promise<CreateOrderResult> {
  const res = await fetch(`${API_BASE}/payments/order`, {
    method: 'POST',
    headers: authHeader(),
    body: JSON.stringify({ venture_id: ventureId, purpose }),
  })
  if (!res.ok) throw new Error(await parseError(res))
  return res.json()
}

export async function getOrder(orderId: string): Promise<Order> {
  const res = await fetch(`${API_BASE}/payments/order/${orderId}`, { headers: authHeader() })
  if (!res.ok) throw new Error(await parseError(res))
  return res.json()
}
