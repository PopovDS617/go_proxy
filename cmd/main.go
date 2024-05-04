package main

import (
	"context"
	httpServer "goproxy/internal/http"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	server := httpServer.NewHTTPServer()

	go func() {
		log.Printf("starting HTTP Proxy Server. Listening at %s", server.Addr)
		if err := server.Server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		} else {
			log.Println("server closed!")
		}
	}()

	sigquit := make(chan os.Signal, 1)
	signal.Notify(sigquit, os.Interrupt, syscall.SIGTERM)
	sig := <-sigquit
	log.Printf("caught sig: %+v", sig)
	log.Printf("gracefully shutting down server...")

	if err := server.Server.Shutdown(context.Background()); err != nil {
		log.Printf("unable to shut down server: %v", err)
	} else {
		log.Println("server stopped")
	}
}
