import { useEffect, useRef } from 'react'
import { DotLottie } from '@lottiefiles/dotlottie-web'

interface Props {
  src: string
  autoplay?: boolean
  loop?: boolean
  className?: string
  style?: React.CSSProperties
}

// Lightweight wrapper around dotlottie-web. Renders a canvas that plays the
// given .json or .lottie animation. Ponytail: no react-specific bindings.
export default function LottiePlayer({ src, autoplay = true, loop = true, className, style }: Props) {
  const canvasRef = useRef<HTMLCanvasElement>(null)
  const instanceRef = useRef<DotLottie | null>(null)

  useEffect(() => {
    if (!canvasRef.current) return
    const dotLottie = new DotLottie({
      canvas: canvasRef.current,
      src,
      autoplay,
      loop,
    })
    instanceRef.current = dotLottie
    return () => {
      dotLottie.destroy()
      instanceRef.current = null
    }
  }, [src, autoplay, loop])

  return <canvas ref={canvasRef} className={className} style={style} />
}
