package main

import (
	"log"

	"github.com/kcchandra/golang-hook/internal/app"
)

func main() {
	app := app.NewApp()

	if err := app.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	if err := app.GracefulShutdown(); err != nil {
		log.Fatalf("Failed to shutdown server: %v", err)
	}
}
