package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName: "Aji", MiddleName: "Pras", LastName: "Tiyo",
	}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))
}
