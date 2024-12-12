package handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/kushalShukla-web/microservice/data"
)

type X struct {
	x *slog.Logger
}

func Newproduct(val *slog.Logger) *X {
	return &X{val}
}

func (c *X) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		c.getrequest(rw, r)
		return
	}
	// Here StatusMethodNotAllowed is a constant with value 405.
	// so when this line executed , it will set Header to 405.i.e Method not Allowed.
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (c *X) getrequest(rw http.ResponseWriter, r *http.Request) {
	products := data.Getdata()
	value, err := json.Marshal(products)
	if err != nil {
		fmt.Println("Error while parsing", err)
	}
	rw.Write(value)
}
