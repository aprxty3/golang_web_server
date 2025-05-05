package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArrayEncode(t *testing.T) {
	customer := Customer{
		FirstName: "Aji", MiddleName: "Pras", LastName: "Tiyo",
		Hobbies: []string{"Gaming", "Reading", "Coding"},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestJsonArrayDecode(t *testing.T) {
	jsonString := `{"FirstName": "Aji", "MiddleName": "Pras", "LastName": "Tiyo", "Hobbies": ["Gaming", "Reading", "Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}
