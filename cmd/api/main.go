package main

import (
	"context"
	"log"
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

	transporthttp.NewServer(
		application.Services(),
		":"+application.Config().HTTPPort,
	).Start()
}
