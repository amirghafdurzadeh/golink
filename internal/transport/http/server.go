package http

import (
	"context"
	"net"
	"net/http"

	"github.com/amirghafdurzadeh/golink/internal/app"
	"github.com/amirghafdurzadeh/golink/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(ctx context.Context, cfg config.HTTPConfig, services app.Services) *Server {
	handler := newRouter(services)

	return &Server{
		httpServer: &http.Server{
			Addr:              ":" + cfg.Port,
			Handler:           handler,
			ReadTimeout:       cfg.ReadTimeout,
			ReadHeaderTimeout: cfg.ReadHeaderTimeout,
			WriteTimeout:      cfg.WriteTimeout,
			IdleTimeout:       cfg.IdleTimeout,
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
