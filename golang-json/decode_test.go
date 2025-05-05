package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	jsonString := `{"FirstName": "Aji", "MiddleName": "Pras", "LastName": "Tiyo"}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}
