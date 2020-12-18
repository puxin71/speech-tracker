package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/puxin71/talk-server/pkg"
	"github.com/puxin71/talk-server/pkg/handler"
)

func main() {
	const addr = "localhost:3000"
	idleConnsClosed := make(chan struct{})

	// Configure server routes
	router := handler.NewRouter(pkg.NewEnvReader())
	http.Handle("/", router)
	server := http.Server{
		Handler: router,
		Addr:    addr,
	}

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shut down.
		log.Println("Shutting down the server")
		if err := server.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown, error: %v", err)
		}
		close(idleConnsClosed)
	}()

	log.Println("Starting server and listens on " + addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe, error: %v", err)
	}

	<-idleConnsClosed
}
