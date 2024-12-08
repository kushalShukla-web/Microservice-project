package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/kushalShukla-web/microservice/handler"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	hello := handler.NewHello(logger)
	sm := http.NewServeMux()
	sm.Handle("/", hello)

	http.ListenAndServe(":9090", nil)
}
