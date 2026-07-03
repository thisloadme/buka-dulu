import { Routes, Route } from 'react-router-dom'
import Landing from './pages/Landing'
import Login from './pages/Login'
import Register from './pages/Register'
import Download from './pages/Download'
import GoogleCallback from './pages/GoogleCallback'
import Dashboard from './pages/Dashboard'
import IdeaCapture from './pages/IdeaCapture'
import IdeaResult from './pages/IdeaResult'
import Paywall from './pages/Paywall'
import History from './pages/History'
import ProtectedRoute from './components/ProtectedRoute'

export default function App() {
  return (
    <Routes>
      <Route path="/" element={<Landing />} />
      <Route path="/login" element={<Login />} />
      <Route path="/register" element={<Register />} />
      <Route path="/auth/google/callback" element={<GoogleCallback />} />
      <Route path="/download" element={<Download />} />

      <Route path="/dashboard" element={<ProtectedRoute><Dashboard /></ProtectedRoute>} />
      <Route path="/idea/new" element={<ProtectedRoute><IdeaCapture /></ProtectedRoute>} />
      <Route path="/idea/:id" element={<ProtectedRoute><IdeaResult /></ProtectedRoute>} />
      <Route path="/pay/:orderId" element={<ProtectedRoute><Paywall /></ProtectedRoute>} />
      <Route path="/history" element={<ProtectedRoute><History /></ProtectedRoute>} />
    </Routes>
  )
}
