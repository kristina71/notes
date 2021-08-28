package app

import (
	"context"
	"log"
	"net/http"
	"notes/pkg/endpoints"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	server *http.Server
}

func New() *App {
	return &App{server: &http.Server{
		Addr:         ":6969",
		Handler:      endpoints.New(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}}
}

func (s *App) Run() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGABRT)

	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%+s\n", err)
		}
	}()

	log.Printf("server started")

	<-c

	log.Printf("server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := s.server.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed:%+s", err)
	}

	log.Printf("server exited properly")
}
