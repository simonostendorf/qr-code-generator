package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the Server",
	Long:  "Start the QR Code generation server.",
	RunE:  executeServer,
}

func executeServer(cmd *cobra.Command, args []string) error {
	if len(args) != 0 {
		return cmd.Help()
	}

	// TODO

	return nil
}

// setup specific flags
func init() {
}
