package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/puxin71/talk-server/pkg/handler"
	"github.com/puxin71/talk-server/pkg/middleware"
)

func main() {
	const addr = "localhost:8080"
	idleConnsClosed := make(chan struct{})

	// Configure server routes
	mux := http.NewServeMux()
	mux.Handle("/", middleware.Logger(handler.Home{}, "index"))
	//mux.Handle("/", handler.Query{})

	server := http.Server{
		Handler: mux,
		Addr:    addr,
	}

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := server.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	fmt.Println("Starting server and listens on " + addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}
