package handler // as per name of Folder

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

type Hello struct {
	x *slog.Logger
}

// Don't bind this function with the above struct as because we are creating
// an object here.
func NewHello(val *slog.Logger) *Hello {
	return &Hello{val}
}

func (c *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received")

	d, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err, "oops")
		http.Error(rw, "oops", http.StatusBadGateway)
		return
	}
	fmt.Fprintf(rw, "Hello %s", d)
}
