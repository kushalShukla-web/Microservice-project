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
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	hello := handler.NewHello(logger)
	product := handler.Newproduct(logger)
	sm := http.NewServeMux()
	sm.Handle("/", hello)
	sm.Handle("/product", product)
	// creating a new customize server server
	s := http.Server{
		Addr:         ":9090",
		Handler:      sm,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
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

}
