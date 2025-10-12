package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"net/http"
	"strings"

	"github.com/simonostendorf/qr-code-generator/pkg/qrcodegenerator"
)

type GenerateBody struct {
	URL                   string            `json:"url"`
	ErrorCorrectionLevel  string            `json:"errorCorrectionLevel,omitempty"`
	Logo                  *GenerateBodyLogo `json:"logo,omitempty"`
	TransparentBackground bool              `json:"transparentBackground,omitempty"`
	Color                 string            `json:"color,omitempty"` // hex, e.g. "#000000"
}

type GenerateBodyLogo struct {
	ImageBase64    string `json:"imageBase64,omitempty"`
	SizeMultiplier int    `json:"sizeMultiplier,omitempty"`
}

func GenerateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// decode request body
	var body GenerateBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, fmt.Sprintf("400 bad request: invalid json: %v", err.Error()), http.StatusBadRequest)
		return
	}

	// map request body to qr code params
	params, err := mapBodyToParams(&body)
	if err != nil {
		http.Error(w, fmt.Sprintf("400 bad request: could not map body to qr-code params: %v", err.Error()), http.StatusBadRequest)
		return
	}

	// generate qr code
	code, err := qrcodegenerator.GenerateQRCode(params)
	if err != nil {
		http.Error(w, fmt.Sprintf("500 internal server error: could not generate qr-code: %v", err.Error()), http.StatusInternalServerError)
		return
	}

	// write api response
	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(http.StatusOK)
	w.Write(code)
}

func mapBodyToParams(body *GenerateBody) (*qrcodegenerator.QRCodeParams, error) {
	params := &qrcodegenerator.QRCodeParams{
		URL:                   body.URL,
		ErrorCorrectionLevel:  qrcodegenerator.ErrorCorrectionLevel(body.ErrorCorrectionLevel),
		TransparentBackground: body.TransparentBackground,
		Color:                 body.Color,
	}

	if body.Logo != nil && body.Logo.ImageBase64 != "" {
		img, err := decodeBase64Image(body.Logo.ImageBase64)
		if err != nil {
			return nil, fmt.Errorf("invalid base64 logo: %v", err)
		}
		params.Logo = &qrcodegenerator.QRCodeLogoParams{
			SizeMultiplier: body.Logo.SizeMultiplier,
			Image:          &img,
		}
	}

	return params, nil
}

func decodeBase64Image(s string) (image.Image, error) {
	// data:image/png;base64,... entfernen
	if strings.Contains(s, ",") {
		s = strings.SplitN(s, ",", 2)[1]
	}
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(bytes.NewReader(data))
	return img, err
}
