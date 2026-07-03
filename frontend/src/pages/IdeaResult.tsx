import { useEffect, useState } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import AppNavbar from '../components/AppNavbar'
import LottiePlayer from '../components/LottiePlayer'
import { getIdea, processIdea, confirmIdea, updateIdea, Idea } from '../api/idea'
import { createOrder } from '../api/payment'

function parseList(s?: string): string[] {
  if (!s) return []
  try {
    const arr = JSON.parse(s)
    return Array.isArray(arr) ? arr : []
  } catch {
    return []
  }
}

export default function IdeaResult() {
  const { id: ventureId = '' } = useParams()
  const navigate = useNavigate()
  const [idea, setIdea] = useState<Idea | null>(null)
  const [loading, setLoading] = useState(true)
  const [processing, setProcessing] = useState(false)
  const [confirming, setConfirming] = useState(false)
  const [editing, setEditing] = useState(false)
  const [error, setError] = useState('')

  // editable fields
  const [concept, setConcept] = useState('')
  const [customer, setCustomer] = useState('')
  const [value, setValue] = useState('')

  const load = async () => {
    setLoading(true)
    setError('')
    try {
      const i = await getIdea(ventureId)
      setIdea(i)
      if (i) {
        setConcept(i.one_line_concept || '')
        setCustomer(i.target_customer || '')
        setValue(i.value_proposition || '')
      }
    } catch (err: any) {
      setError(err.message || 'Gagal memuat ide.')
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => { load() }, [ventureId])

  const handleProcess = async () => {
    setError('')
    setProcessing(true)
    try {
      const result = await processIdea(ventureId)
      if (result.paymentRequired) {
        // Create QRIS order and go to paywall
        const order = await createOrder(ventureId, result.purpose || 'idea_validation')
        if (order.free) {
          // free path — retry process
          setProcessing(false)
          return handleProcess()
        }
        navigate(`/pay/${order.order_id}`)
        return
      }
      if (result.idea) setIdea(result.idea)
      setConcept(result.idea?.one_line_concept || '')
      setCustomer(result.idea?.target_customer || '')
      setValue(result.idea?.value_proposition || '')
    } catch (err: any) {
      setError(err.message || 'Gagal memproses ide.')
    } finally {
      setProcessing(false)
    }
  }

  const handleSaveEdit = async () => {
    setError('')
    try {
      const updated = await updateIdea(ventureId, {
        one_line_concept: concept,
        target_customer: customer,
        value_proposition: value,
      } as any)
      setIdea(updated)
      setEditing(false)
    } catch (err: any) {
      setError(err.message || 'Gagal menyimpan perubahan.')
    }
  }

  const handleConfirm = async () => {
    setError('')
    setConfirming(true)
    try {
      await confirmIdea(ventureId)
      navigate('/dashboard')
    } catch (err: any) {
      setError(err.message || 'Gagal mengunci ide.')
    } finally {
      setConfirming(false)
    }
  }

  if (loading) {
    return (
      <>
        <AppNavbar />
        <div className="container app-content">
          <p className="body-text">Memuat ide...</p>
        </div>
      </>
    )
  }

  if (error && !idea) {
    return (
      <>
        <AppNavbar />
        <div className="container app-content">
          <div className="auth-error">{error}</div>
        </div>
      </>
    )
  }

  const isProcessed = idea?.status === 'done' && idea.one_line_concept
  const assumptions = parseList(idea?.key_assumptions)
  const risks = parseList(idea?.early_risks)

  return (
    <>
      <AppNavbar />
      <div className="container app-content">
        <button className="btn btn-ghost" onClick={() => navigate('/dashboard')}>← Kembali</button>
        <h1 className="display-large" style={{ fontSize: '2rem', marginTop: 12 }}>
          {idea?.one_line_concept || 'Analisis Ide'}
        </h1>

        {error && <div className="auth-error" style={{ marginTop: 12 }}>{error}</div>}

        {!isProcessed && (
          <div className="card" style={{ marginTop: 20 }}>
            <div className="card-body" style={{ textAlign: 'center', padding: 32 }}>
              {processing ? (
                <>
                  <LottiePlayer src="/lottie/meditation.json" style={{ width: 180, margin: '0 auto' }} />
                  <h3 style={{ marginTop: 16 }}>AI sedang menganalisis ide kamu...</h3>
                  <p className="body-text">Mohon tunggu beberapa detik.</p>
                </>
              ) : (
                <>
                  <LottiePlayer src="/lottie/Creative-Idea.json" style={{ width: 180, margin: '0 auto' }} loop={false} />
                  <h3 style={{ marginTop: 16 }}>Siap dianalisis</h3>
                  <p className="body-text" style={{ marginTop: 8 }}>
                    Klik tombol di bawah untuk menjalankan AI menyusun konsep terstruktur dari ide kamu.
                  </p>
                  <button className="btn btn-primary btn-lg" style={{ marginTop: 16 }} onClick={handleProcess}>
                    Proses dengan AI
                  </button>
                </>
              )}
            </div>
          </div>
        )}

        {isProcessed && (
          <>
            <div className="raw-idea" style={{ marginTop: 20 }}>
              <div className="caption">Ide Mentah</div>
              <p className="body-text" style={{ marginTop: 4 }}>{idea?.raw_input}</p>
            </div>

            <div className="result-grid" style={{ marginTop: 24 }}>
              <div className="card">
                <div className="card-body">
                  <div className="micro">Konsep Satu Kalimat</div>
                  {editing ? (
                    <textarea value={concept} onChange={e => setConcept(e.target.value)} rows={2} />
                  ) : (
                    <p className="body-large" style={{ marginTop: 8 }}>{concept}</p>
                  )}
                </div>
              </div>
              <div className="card">
                <div className="card-body">
                  <div className="micro">Target Pelanggan</div>
                  {editing ? (
                    <textarea value={customer} onChange={e => setCustomer(e.target.value)} rows={3} />
                  ) : (
                    <p className="body-large" style={{ marginTop: 8 }}>{customer}</p>
                  )}
                </div>
              </div>
              <div className="card">
                <div className="card-body">
                  <div className="micro">Value Proposition</div>
                  {editing ? (
                    <textarea value={value} onChange={e => setValue(e.target.value)} rows={3} />
                  ) : (
                    <p className="body-large" style={{ marginTop: 8 }}>{value}</p>
                  )}
                </div>
              </div>
              <div className="card">
                <div className="card-body">
                  <div className="micro">Asumsi Kunci</div>
                  <ul className="idea-list" style={{ marginTop: 8 }}>
                    {assumptions.map((a, i) => <li key={i}>{a}</li>)}
                  </ul>
                </div>
              </div>
              <div className="card">
                <div className="card-body">
                  <div className="micro">Risiko Awal</div>
                  <ul className="idea-list" style={{ marginTop: 8 }}>
                    {risks.map((r, i) => <li key={i}>{r}</li>)}
                  </ul>
                </div>
              </div>
            </div>

            <div style={{ marginTop: 24, display: 'flex', gap: 12, flexWrap: 'wrap' }}>
              {editing ? (
                <>
                  <button className="btn btn-primary" onClick={handleSaveEdit}>Simpan</button>
                  <button className="btn btn-ghost" onClick={() => setEditing(false)}>Batal</button>
                </>
              ) : (
                <button className="btn btn-ghost" onClick={() => setEditing(true)} disabled={idea?.is_locked}>
                  Edit Hasil
                </button>
              )}
              <button className="btn btn-primary" onClick={handleConfirm} disabled={confirming || idea?.is_locked}>
                {confirming ? 'Memproses...' : idea?.is_locked ? 'Ide Dikunci' : 'Kunci Ide & Selesai'}
              </button>
            </div>
          </>
        )}
      </div>
    </>
  )
}
