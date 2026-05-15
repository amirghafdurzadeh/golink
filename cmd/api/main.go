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

func main() {
	startupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	application, err := app.New(startupCtx)
	if err != nil {
		log.Fatalf("failed to build app: %v", err)
	}
	defer application.Close()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	httpServer := transporthttp.NewServer(
		ctx,
		application.Services(),
		":"+application.Config().HTTPPort,
	)

	errCh := make(chan error, 1)
	go func() {
		log.Printf("http server started on :%s", application.Config().HTTPPort)
		if err := httpServer.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	select {
	case err := <-errCh:
		log.Fatalf("http server failed: %v", err)
	case <-ctx.Done():
		log.Println("shutdown signal received")
	}

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Printf("http shutdown failed: %v", err)
	}

	if err := application.Close(); err != nil {
		log.Printf("app close failed: %v", err)
	}

	log.Println("application stopped")
}
