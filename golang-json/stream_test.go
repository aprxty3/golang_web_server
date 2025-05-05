package golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamDecoder(t *testing.T) {
	reader, err := os.Open("customer.json")
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(reader)

	customer := &Customer{}

	decoder.Decode(customer)

	fmt.Println(customer)
}

func TestStreamEncoder(t *testing.T) {
	create, _ := os.Create("customer_out.json")
	encoder := json.NewEncoder(create)

	customer := Customer{
		FirstName: "Aji", MiddleName: "Pras", LastName: "Tiyo",
	}
	encoder.Encode(customer)
}
