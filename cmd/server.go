package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

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

	// stop gracefully on SIGTERM (Kubernetes) and SIGINT (Ctrl+C)
	ctx, stop := signal.NotifyContext(cmd.Context(), syscall.SIGTERM, os.Interrupt)
	defer stop()

	// create and start the server
	srv := server.NewServer(port)

	fmt.Printf("Starting server on port %d...\n", srv.Port)

	if err := srv.Start(ctx); err != nil {
		return err
	}

	fmt.Println("Server stopped")

	return nil
}

// setup specific flags
func init() {
	serverCmd.Flags().Uint("port", 8000, "Port to run the server on")
}
