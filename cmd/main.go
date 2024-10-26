package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	http "github.com/ramasuryananda/dummy-cv/internal/server"
)

func main() {
	ctx := context.Background()
	server := http.NewHTTP(ctx)

	go func() {
		err := server.Run().ListenAndServe()
		if err != nil {

			log.Fatal("server.Run().ListenAndServe() error - mainHTTP")
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	fmt.Printf("Server stopped\n")

	os.Exit(0)
}
