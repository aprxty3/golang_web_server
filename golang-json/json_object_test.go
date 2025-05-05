package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode string
}
type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
	Addresses  []Address
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName: "Aji", MiddleName: "Pras", LastName: "Tiyo",
	}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))
}
