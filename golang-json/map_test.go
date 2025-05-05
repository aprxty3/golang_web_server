package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonMapDecode(t *testing.T) {
	jsonString := `{"id": "P001", "name":"Apple iPhone 13", "price": 13000}`
	jsonBytes := []byte(jsonString)

	var product map[string]interface{}

	json.Unmarshal(jsonBytes, &product)

	fmt.Println(product)
}

func TestJsonMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "P0001",
		"name":  "Apple Macbook",
		"price": 13000,
	}

	bytes, _ := json.Marshal(product)

	fmt.Println(string(bytes))

}
