package http

import (
	"log"
	"net/http"
	"time"

	"github.com/amirghafdurzadeh/golink/internal/app"
)

type Server struct {
	http *http.Server
}

func NewServer(services app.Services, addr string) *Server {
	handler := newRooter(services)

	return &Server{
		http: &http.Server{
			Addr:              addr,
			Handler:           handler,
			ReadTimeout:       5 * time.Second,
			ReadHeaderTimeout: 5 * time.Second,
			WriteTimeout:      10 * time.Second,
			IdleTimeout:       60 * time.Second,
		},
	}
}

func (s *Server) Start() {
	log.Printf("server starting on %s", s.http.Addr)

	if err := s.http.ListenAndServe(); err != nil &&
		err != http.ErrServerClosed {
		log.Fatalf("server failed: %v", err)
	}
}
