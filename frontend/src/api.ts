// The API is served on the same origin: nginx proxies /api/ to the backend in
// the image, the vite dev server does the same in development.
export const apiBase = import.meta.env.VITE_API_BASE ?? ''

export interface GenerateBodyLogo {
  imageBase64?: string
  sizeMultiplier?: number
}

export interface GenerateBody {
  url: string
  errorCorrectionLevel?: string
  transparentBackground?: boolean
  color?: string
  logo?: GenerateBodyLogo
}

export async function generateQRCode(params: GenerateBody): Promise<string> {
  const response = await fetch(`${apiBase}/api/generate`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(params),
  })

  if (!response.ok) {
    throw new Error(`Failed to generate QR: ${response.statusText}`)
  }

  const blob = await response.blob()
  return URL.createObjectURL(blob)
}
