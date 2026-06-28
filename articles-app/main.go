package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nios-x/articles-go/internal/config"
)

func main() {
	cfg := config.MustLoad()

	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	// Create the server
	server := &http.Server{
		Addr:    cfg.HttpServer.Addr,
		Handler: router,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	fmt.Println("Starting Server at ", "http://"+cfg.HttpServer.Addr)

	// Start the server
	go func() {
		if err := server.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			panic(err)
		}
	}()

	// Wait for Ctrl+C
	<-done
	fmt.Println("Shutting down server...")

	// Give active requests 5 seconds to finish
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("Shutdown error:", err)
	}

	fmt.Println("Server stopped")
}
