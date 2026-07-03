import { useEffect, useState } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import AppNavbar from '../components/AppNavbar'
import LottiePlayer from '../components/LottiePlayer'
import { getOrder, Order } from '../api/payment'

export default function Paywall() {
  const { orderId = '' } = useParams()
  const navigate = useNavigate()
  const [order, setOrder] = useState<Order | null>(null)
  const [loading, setLoading] = useState(true)
  const [paid, setPaid] = useState(false)
  const [error, setError] = useState('')

  useEffect(() => {
    let cancelled = false

    const poll = async () => {
      try {
        const o = await getOrder(orderId)
        if (cancelled) return
        setOrder(o)
        setLoading(false)
        if (o.status === 'paid') {
          setPaid(true)
          // Wait briefly to show success animation, then return to idea processing
          setTimeout(() => {
            if (o.venture_id) navigate(`/idea/${o.venture_id}`)
          }, 2500)
        }
        if (o.status === 'expired') {
          setError('Waktu pembayaran habis. Silakan buat ide baru.')
        }
      } catch (err: any) {
        if (!cancelled) {
          setError(err.message || 'Gagal memuat order.')
          setLoading(false)
        }
      }
    }

    poll()
    const interval = setInterval(poll, 3000)
    return () => { cancelled = true; clearInterval(interval) }
  }, [orderId, navigate])

  if (loading) {
    return (
      <>
        <AppNavbar />
        <div className="container app-content"><p className="body-text">Menyiapkan pembayaran...</p></div>
      </>
    )
  }

  if (paid) {
    return (
      <>
        <AppNavbar />
        <div className="container app-content">
          <div className="card" style={{ maxWidth: 480, margin: '40px auto', textAlign: 'center' }}>
            <div className="card-body" style={{ padding: 32 }}>
              <LottiePlayer src="/lottie/Confetti.json" style={{ width: 220, margin: '0 auto' }} loop={false} />
              <h2 style={{ marginTop: 16 }}>Pembayaran Berhasil!</h2>
              <p className="body-text">Mengarahkan kembali untuk memproses ide kamu...</p>
            </div>
          </div>
        </div>
      </>
    )
  }

  return (
    <>
      <AppNavbar />
      <div className="container app-content">
        <div className="paywall" style={{ maxWidth: 460, margin: '24px auto' }}>
          <h1 className="display-large" style={{ fontSize: '1.75rem' }}>Validasi Ide Berikutnya</h1>
          <p className="body-large" style={{ marginTop: 8 }}>
            Kuota gratis kamu sudah terpakai. Bayar <strong>Rp{order?.amount.toLocaleString('id-ID')}</strong> untuk validasi ide ini.
          </p>

          {error && <div className="auth-error" style={{ marginTop: 12 }}>{error}</div>}

          {order?.qris_image && (
            <div className="qris-box" style={{ marginTop: 20, textAlign: 'center' }}>
              <img src={order.qris_image} alt="QRIS" style={{ width: '100%', maxWidth: 280, margin: '0 auto' }} />
              <p className="caption" style={{ marginTop: 8 }}>Scan QRIS di atas untuk membayar</p>
              {order.total_amount && (
                <p className="body-large" style={{ marginTop: 8 }}>
                  Total: <strong>Rp{Number(order.total_amount).toLocaleString('id-ID')}</strong>
                </p>
              )}
              {order.expired_at && (
                <p className="caption">Bayar sebelum {new Date(order.expired_at).toLocaleTimeString('id-ID')}</p>
              )}
              <p className="caption" style={{ marginTop: 12 }}>Menunggu pembayaran...</p>
            </div>
          )}

          <button className="btn btn-ghost btn-block" style={{ marginTop: 16 }} onClick={() => navigate('/dashboard')}>
            Batalkan
          </button>
        </div>
      </div>
    </>
  )
}
