package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/kcchandra/golang-hook/internal/config"
	"github.com/kcchandra/golang-hook/router"
)

type App struct {
	Config *config.Config
	Server *http.Server
}

func NewApp() *App {
	app := &App{
		Config: config.NewConfig(),
	}

	app.Server = &http.Server{
		Addr: ":" + app.Config.Port,
		Handler: func() http.Handler {
			server := http.NewServeMux()
			router.RegisterRoutes(server)
			return server
		}(),
	}

	return app
}

func (a *App) Start() error {
	go func() {
		if err := a.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	log.Printf("Server started on port %s", a.Config.Port)
	return nil
}

func (a *App) GracefulShutdown() error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	if err := a.Server.Shutdown(context.Background()); err != nil {
		return err
	}

	log.Println("Server closed")
	return nil
}
