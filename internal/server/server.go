package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/simonostendorf/qr-code-generator/internal/server/api"
)

// shutdownTimeout is how long running requests may take to finish after a
// stop signal; it stays below the 30 second Kubernetes grace period.
const shutdownTimeout = 10 * time.Second

type Server struct {
	Port uint
}

func NewServer(port uint) *Server {
	return &Server{
		Port: port,
	}
}

// Start serves the API until ctx is cancelled, then shuts down gracefully.
func (s *Server) Start(ctx context.Context) error {
	mux := http.NewServeMux()

	s.registerRoutes(mux)

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", s.Port),
		Handler:           withCORS(mux),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	errCh := make(chan error, 1)
	go func() {
		errCh <- srv.ListenAndServe()
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("graceful shutdown: %w", err)
	}
	if err := <-errCh; !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (s *Server) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/health", api.HealthHandler)
	mux.HandleFunc("/api/generate", api.GenerateHandler)
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
