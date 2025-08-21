package jsonhandling

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUnMarshalJson(t *testing.T) {
	jsonString := `{"FirstName":"Bismen","MiddleName":"Gusti","LastName":"Taka"}`

	// ubah dulu ke slice of byte
	// baru unmarshaling

	jsonByte := []byte(jsonString)

	customer := &Customer{}

	json.Unmarshal(jsonByte, customer)

	fmt.Println("Data Setelah di decode: ", customer)

	fmt.Println("Nama Pertama", customer.FirstName)
}
