package cmd

import (
	"fmt"

	"github.com/simonostendorf/qr-code-generator/internal/server"
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
	// validate arguments
	if len(args) != 0 {
		return cmd.Help()
	}

	// parse flags and arguments
	port, _ := cmd.Flags().GetUint("port")

	// create and start the server
	srv := server.NewServer(port)

	fmt.Printf("Starting server on port %d...\n", srv.Port)

	return srv.Start()
}

// setup specific flags
func init() {
	serverCmd.Flags().Uint("port", 8000, "Port to run the server on")
}
