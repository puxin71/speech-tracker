package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/puxin71/talk-server/pkg/database"

	"github.com/puxin71/talk-server/pkg"
	"github.com/puxin71/talk-server/pkg/handler"
)

const (
	// Our app's URL
	Addr = "localhost:3000"
	// Repeatedly ping the database server to check the connection's status
	DBConnTickerPeriod = 10 * time.Second
	// Connection timeout period used to decide if the connection to the server is down
	DBConnTimeout = 1 * time.Second
)

func main() {
	idleConnsClosed := make(chan struct{})

	// Cancel all sub-goroutines when our app exits
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Connect to MySQL database
	db, err := database.NewMySQLDB()
	if err != nil {
		return
	}
	defer db.Close()
	go monitor(ctx, db, DBConnTickerPeriod, DBConnTimeout)

	// Configure server routes
	router := handler.NewRouter(pkg.NewEnvReader())
	http.Handle("/", router)
	server := http.Server{
		Handler: router,
		Addr:    Addr,
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

	log.Println("Starting server and listens on " + Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe, error: %v", err)
	}

	<-idleConnsClosed
}

func monitor(ctx context.Context, db database.MySQLDB, timer time.Duration, timeout time.Duration) {
	ticker := time.NewTicker(timer)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			db.Ping(ctx, DBConnTimeout)
		}
	}
}
