package jsonhandling

import (
	"encoding/json"
	"fmt"
	"testing"
)

// encode
func logJson(data interface{}) {
	byte, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(byte))

}

func TestMarsahalLearn(t *testing.T) {
	logJson("Bismen")
	logJson(2)
	logJson(true)
	logJson([]string{"Bismen", "Gusti", "Taka"})
}

// json object

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
}

// encode
func TestMarshalJson(t *testing.T) {
	customer1 := &Customer{
		FirstName:  "Bismen",
		LastName:   "Taka",
		MiddleName: "Gusti",
	}

	byte, err := json.Marshal(customer1)

	if err != nil {
		// ada sesutu
		panic(err)
	}

	fmt.Println(string(byte))

}
