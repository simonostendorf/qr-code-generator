package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "qr-code-generator",
	Version: "0.0.1",
	Short:   "Simple QR Code Generator",
	Long:    "A simple QR code generator that creates QR codes from URLs and has the option to add a logo in the middle.",
}

func Execute() {
	// cmd as command with Run instead of RunE, so no error is expected
	// nolint:errcheck
	rootCmd.Execute() //gosec:disable G104
}

// setup global flags
func init() {
	// none
}
