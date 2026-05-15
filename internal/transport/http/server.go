package http

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/amirghafdurzadeh/golink/internal/app"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(ctx context.Context, services app.Services, addr string) *Server {
	handler := newRouter(services)

	return &Server{
		httpServer: &http.Server{
			Addr:              addr,
			Handler:           handler,
			ReadTimeout:       5 * time.Second,
			ReadHeaderTimeout: 5 * time.Second,
			WriteTimeout:      10 * time.Second,
			IdleTimeout:       60 * time.Second,
			BaseContext: func(_ net.Listener) context.Context {
				return ctx
			},
		},
	}
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
