package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID          int32   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Sku         string  `json:"sku"`
	Createdon   string  `json:"-"`
	Updateon    string  `json:"-"`
	Deletedon   string  `json:"-"`
}

func (p *Product) Fromjson(x io.Reader) error {
	// NewDecoder is used to Decode the Json data into Struct
	// the NewDecoder main use is to input data from x (io.Reader)
	// and convert it into struct data. and its decoded in struct p after
	// calling e.Decode function.
	e := json.NewDecoder(x)
	return e.Decode(p)
}

var ErrProductNotFound = fmt.Errorf("Product not found")
var Samples = []*Product{
	{
		ID:          1,
		Name:        "latte",
		Description: "data1",
		Price:       100,
		Sku:         "bc123",
		Createdon:   time.Now().UTC().String(),
		Updateon:    time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Cappuccino",
		Description: "data2",
		Price:       200,
		Sku:         "abc1234",
		Createdon:   time.Now().UTC().String(),
		Updateon:    time.Now().UTC().String(),
	},
}

func Getdata() []*Product {
	return Samples
}
func Addata(x *Product) {
	x.ID = getnextid()
	Samples = append(Samples, x)
}

func getnextid() int32 {
	fmt.Println(Samples[len(Samples)-1].ID+1, "Hi from Length")
	return Samples[len(Samples)-1].ID + 1
}

func UpdateProduct(id int32, p Product) error {
	i := findIndexbyid(id)
	if i != -1 {
		p.ID = id
		Samples[i] = &p
		return nil
	}
	return ErrProductNotFound
}

func findIndexbyid(id int32) int32 {
	for i, p := range Samples {
		if p.ID == id {
			return int32(i)
		}

	}
	return -1
}
