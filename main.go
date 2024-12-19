package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/kushalShukla-web/microservice/handler"
)

func main() {
	// Here We are creating a logger object for structure logging in our Application.
	// if you wanted to learn more about structure logging visit this site https://betterstack.com/community/guides/logging/logging-in-go/
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	// we are creating an object for diffrent parts of your application.
	// example - for Purchasing coffee there gonna be a different handlers and for payment there gonna
	// be a different handler.
	product := handler.Newproduct(logger)
	// Here creating a simple server with servermux.
	sm := http.NewServeMux()
	// routing to that request i.e get post put etc.
	sm.Handle("/", product)
	// creating a new customize server server
	s := http.Server{
		Addr:         ":9090",         // this server is using port 9090 for external communication
		Handler:      sm,              // to use Default servermux or created by us.
		ReadTimeout:  5 * time.Second, // maximum time to read request
		WriteTimeout: 5 * time.Second,
	}
	// running in go routine
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			logger.Error("err", err)
			os.Exit(1)
		}
	}()
	// 1. Dont understand this part !
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Println("Got signal:", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
	///////////

}
