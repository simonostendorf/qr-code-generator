package cmd

import (
	"fmt"

	"github.com/simonostendorf/qr-code-generator/internal/helpers"
	"github.com/simonostendorf/qr-code-generator/pkg/qrcodegenerator"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate <url>",
	Short: "Generate a QR Code",
	Long:  "Generate a QR Code for the specified URL and save it to a file.",
	RunE:  executeGenerate,
}

func executeGenerate(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return cmd.Help()
	}

	url := args[0]
	color, _ := cmd.Flags().GetString("color")
	logoPath, _ := cmd.Flags().GetString("logo")
	logoSizeMultiplier, _ := cmd.Flags().GetInt("logo-size-multiplier")
	errorCorrection, _ := cmd.Flags().GetString("error-correction")
	transparentBackground, _ := cmd.Flags().GetBool("transparent-background")
	outFile, _ := cmd.Flags().GetString("out")

	params := qrcodegenerator.QRCodeParams{
		URL:                   url,
		Color:                 color,
		TransparentBackground: transparentBackground,
		ErrorCorrectionLevel:  qrcodegenerator.ErrorCorrectionLevel(errorCorrection),
	}

	if logoPath != "" {
		logo, err := helpers.ImageFromFile(logoPath)
		if err != nil {
			return fmt.Errorf("❌ failed to load logo image: %w", err)
		}
		params.Logo.Image = logo
		params.Logo.SizeMultiplier = logoSizeMultiplier
	}

	qrcode, err := qrcodegenerator.GenerateQRCode(&params)
	if err != nil {
		return fmt.Errorf("❌ failed to generate QR code: %w", err)
	}

	// save the QR code to file
	err = helpers.WriteToFile(outFile, qrcode)
	if err != nil {
		return fmt.Errorf("❌ failed to save QR code to file: %w", err)
	}

	fmt.Printf("✅ QR code generated and saved to \"%s\"\n", outFile)

	return nil
}

// setup specific flags
func init() {
	generateCmd.Flags().String("color", "#000000", "Color of the QR code in HEX format")

	generateCmd.Flags().String("logo", "", "Path to logo image to embed in the QR code")
	generateCmd.Flags().Int("logo-size-multiplier", 1, "Size multiplier for the logo. The logo should not exceed 1/5 of the QR code size.")

	generateCmd.Flags().String("error-correction", "medium", "Error correction level: low (7%), medium (15%), high (25%), highest (30%)")

	generateCmd.Flags().Bool("transparent-background", false, "Generate QR code with transparent background")

	generateCmd.Flags().String("out", "qrcode.png", "Output filename")
}
