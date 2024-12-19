package handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"regexp"
	"strconv"

	"github.com/kushalShukla-web/microservice/data"
)

type X struct {
	x *slog.Logger
}

// this object will hold all the crud operations are performed on the product object.
func Newproduct(val *slog.Logger) *X {
	return &X{val}
}

// ServeHTTP is a function which is defined by us and used by servermux internaly .
// the reason we define ServeHTTP function by us cause it gives us flexibility to modify
// its nature.
func (c *X) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		c.getrequest(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		c.addrequest(rw, r)
		return
	}
	if r.Method == http.MethodPut {
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 {
			c.x.Error("More than one id's are passed")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			c.x.Error("Invalid URI more than one capture group")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		x, err := strconv.Atoi(g[0][1])
		if err != nil {
			c.x.Error("Error in converting string %", err)
			return
		}
		id := int32(x)
		if err != nil {
			c.x.Error("Unable to convert string to integer")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		c.putrequest(id, rw, r)
		// if we dont provide the return statement after the execution of the above function
		// the put method is interpret the incomplete state as a failure , leading to 405
		// error
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

func (c *X) addrequest(rw http.ResponseWriter, r *http.Request) {
	incoming := r.Body
	x := &data.Product{}
	// Here Fromjson is changing x i.e its getting value in json
	// and putting it in x.
	err := x.Fromjson(incoming)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}
	fmt.Println(x, "wooooo")
	data.Addata(x)
}

func (c *X) putrequest(id int32, rw http.ResponseWriter, r *http.Request) {
	// we are reading from the body
	incoming := r.Body
	x := &data.Product{}
	err := x.Fromjson(incoming)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}
	err = data.UpdateProduct(id, *x)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusNotFound)
	}

}
