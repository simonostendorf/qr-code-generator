package qrcodegenerator

import (
	"bytes"
	"fmt"

	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

type WriterCloser struct {
	*bytes.Buffer
}

func (w *WriterCloser) Close() error {
	return nil
}

func GenerateQRCode(params *QRCodeParams) ([]byte, error) {
	if params.URL == "" {
		return nil, fmt.Errorf("url is required")
	}

	// setup qr-code
	code, err := qrcode.NewWith(params.URL, qrcode.WithEncodingMode(qrcode.EncModeByte), mapErrorCorrectionLevel(params.ErrorCorrectionLevel))
	if err != nil {
		return nil, fmt.Errorf("could not setup qr-code: %v", err)
	}

	// map code options from params
	options := mapOptions(params)

	// setup buffer to write to
	wc := &WriterCloser{&bytes.Buffer{}}
	writer := standard.NewWithWriter(wc, options...)

	// write to buffer
	err = code.Save(writer)
	if err != nil {
		return nil, fmt.Errorf("could not generate qr-code: %v", err)
	}

	return wc.Bytes(), nil
}

func mapOptions(params *QRCodeParams) []standard.ImageOption {
	// set default options
	options := []standard.ImageOption{
		standard.WithBuiltinImageEncoder(standard.PNG_FORMAT),
		standard.WithBorderWidth(5),
	}

	// set color
	color := "#000000" // default black
	if params.Color != "" {
		color = params.Color
	}
	options = append(options, standard.WithFgColorRGBHex(color))

	// set logo options
	if params.Logo != nil {
		options = append(options, []standard.ImageOption{
			standard.WithLogoSizeMultiplier(params.Logo.SizeMultiplier),
			standard.WithLogoImage(*params.Logo.Image),
			standard.WithLogoSafeZone(),
		}...)
	}

	// set transparent background
	if params.TransparentBackground {
		options = append(options, standard.WithBgTransparent())
	}

	return options
}

func mapErrorCorrectionLevel(level ErrorCorrectionLevel) qrcode.EncodeOption {
	switch level {
	case ErrorCorrectionLow:
		return qrcode.WithErrorCorrectionLevel(qrcode.ErrorCorrectionLow)
	case ErrorCorrectionMedium:
		return qrcode.WithErrorCorrectionLevel(qrcode.ErrorCorrectionMedium)
	case ErrorCorrectionHigh:
		return qrcode.WithErrorCorrectionLevel(qrcode.ErrorCorrectionQuart)
	case ErrorCorrectionHighest:
		return qrcode.WithErrorCorrectionLevel(qrcode.ErrorCorrectionHighest)
	default:
		return qrcode.WithErrorCorrectionLevel(qrcode.ErrorCorrectionMedium) // default
	}
}
