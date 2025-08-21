package jsonhandling

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

func TestJsonTags(t *testing.T) {
	// Kalau secara default akan di map sesuai nama struct di golang untuk jsonnya. jika kita ingin jsonnya berbeda maka bisa mengunakan json tag

	product1 := &Product{
		Id:    "3",
		Name:  "Baju",
		Price: "300000000",
	}

	byte, err := json.Marshal(product1)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(byte))
}
