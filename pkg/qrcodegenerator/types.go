package qrcodegenerator

import "image"

type ErrorCorrectionLevel string

const (
	ErrorCorrectionLow     ErrorCorrectionLevel = "low"     // 7%
	ErrorCorrectionMedium  ErrorCorrectionLevel = "medium"  // 15% (default)
	ErrorCorrectionHigh    ErrorCorrectionLevel = "high"    // 25%
	ErrorCorrectionHighest ErrorCorrectionLevel = "highest" // 30%
)

type QRCodeParams struct {
	URL                   string
	ErrorCorrectionLevel  ErrorCorrectionLevel
	Logo                  *QRCodeLogoParams
	TransparentBackground bool
	Color                 string // hex color code, black by default
}

type QRCodeLogoParams struct {
	SizeMultiplier int
	Image          *image.Image // should be at most 1/5 of the QR code size
}
