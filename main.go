package main

import (
	"fmt"

	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

func main() {
	// // read url from args "go run main.go [url] <code>"
	// if len(os.Args) < 2 || len(os.Args) > 3 {
	// 	fmt.Println("Usage: go run main.go [url] <code>")
	// 	return
	// }

	// url := os.Args[1]

	// code := ""
	// if len(os.Args) == 3 {
	// 	code = os.Args[2]
	// } else {
	// 	// generate 5 code random string with numbers and letters
	// 	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	// 	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 	randomStr := make([]byte, 5)
	// 	for i := range randomStr {
	// 		randomStr[i] = charset[rng.Intn(len(charset))]
	// 	}
	// 	code = string(randomStr)
	// }

	// fmt.Printf("Used short url code: %s\n", code)

	// // write to codes.txt
	// f, err := os.OpenFile("codes.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	fmt.Printf("could not open codes.txt: %v", err)
	// 	return
	// }
	// defer f.Close()

	// if _, err := f.WriteString(fmt.Sprintf("%s %s\n", code, url)); err != nil {
	// 	fmt.Printf("could not write to codes.txt: %v", err)
	// 	return
	// }

	//qrc, err := qrcode.NewWith("https://fsr5.de/"+code, qrcode.WithEncodingMode(qrcode.EncModeByte), qrcode.WithErrorCorrectionLevel(qrcode.ErrorCorrectionHighest))
	qrc, err := qrcode.NewWith("https://portals.fsr5.de", qrcode.WithEncodingMode(qrcode.EncModeByte), qrcode.WithErrorCorrectionLevel(qrcode.ErrorCorrectionHighest))
	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
		return
	}

	options := []standard.ImageOption{
		standard.WithLogoSizeMultiplier(1),
		standard.WithLogoImageFilePNG("logo.png"),
		standard.WithBuiltinImageEncoder(standard.PNG_FORMAT),
		standard.WithBgTransparent(),
		standard.WithLogoSafeZone(),
		//standard.WithCircleShape(),
		standard.WithBorderWidth(5),
		standard.WithFgColorRGBHex("#02b9ad"),
	}
	//w, err := standard.New(code+".png", options...)
	w, err := standard.New("qrcode.png", options...)
	if err != nil {
		fmt.Printf("QR generation failed: %v", err)
		return
	}

	// save file
	if err = qrc.Save(w); err != nil {
		fmt.Printf("could not save image: %v", err)
	}
}
