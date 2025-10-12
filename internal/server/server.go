package server

import (
	"fmt"
	"net/http"

	"github.com/simonostendorf/qr-code-generator/internal/server/api"
)

type Server struct {
	Port uint
}

func NewServer(port uint) *Server {
	return &Server{
		Port: port,
	}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()

	s.registerRoutes(mux)

	return http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", s.Port), withCORS(mux))
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
