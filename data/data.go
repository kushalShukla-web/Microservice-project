package data

import (
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

var Samples = []*Product{
	{
		ID:   1,
		Name: "latte",
		Description: "The iconic latte is loved in coffee shops all over the world. " +
			"The subtle coffee taste and creamy texture makes it a coffee that’s " +
			"universally cherished by even the most casual of coffee drinkers.",
		Price:     100,
		Sku:       "bc123",
		Createdon: time.Now().UTC().String(),
		Updateon:  time.Now().UTC().String(),
	},
	{
		ID:   2,
		Name: "Cappuccino",
		Description: "The name comes from the Capuchin friars, referring to the color of their habits, " +
			"and in this context, referring to the color of the beverage when milk is added in small portion to dark, brewed coffee.",
		Price:     200,
		Sku:       "abc1234",
		Createdon: time.Now().UTC().String(),
		Updateon:  time.Now().UTC().String(),
	},
}

func Getdata() []*Product {
	return Samples
}
