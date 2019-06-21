package main

import (
	"./server"
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// This is the domain the server should accept connections for.

	handler := server.NewRouter()
	// http.ListenAndServe( ":443", handler)
 
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}


	go func() {
		srv.ListenAndServe()
	

	}()

	// Wait for an interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Attempt a graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)

}
