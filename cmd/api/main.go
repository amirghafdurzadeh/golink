package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/amirghafdurzadeh/golink/internal/app"
	transporthttp "github.com/amirghafdurzadeh/golink/internal/transport/http"
)

const shutdownTimeout = 10 * time.Second

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	application, err := buildApplication(ctx)
	if err != nil {
		log.Fatalf("failed to build app: %v", err)
	}

	if err := run(ctx, application); err != nil {
		log.Fatalf("application error: %v", err)
	}

	log.Println("application stopped")
}

func buildApplication(ctx context.Context) (app.Application, error) {
	startupCtx, cancel := context.WithTimeout(ctx, shutdownTimeout)
	defer cancel()

	return app.New(startupCtx)
}

func run(ctx context.Context, application app.Application) error {
	httpServer := transporthttp.NewServer(
		ctx,
		application.Services(),
		":"+application.Config().HTTPPort,
	)

	errCh := make(chan error, 1)

	go func() {
		log.Printf(
			"http server started on :%s",
			application.Config().HTTPPort,
		)

		if err := httpServer.Start(); err != nil &&
			!errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	select {
	case err := <-errCh:
		return err

	case <-ctx.Done():
		log.Println("shutdown signal received")
	}

	return shutdown(application, httpServer)
}

func shutdown(application app.Application, httpServer *transporthttp.Server) error {
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Printf("http shutdown failed: %v", err)
	}

	if err := application.Close(); err != nil {
		return err
	}

	return nil
}
