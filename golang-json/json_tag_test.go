package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func TestJsonTag(t *testing.T) {
	product := Product{
		Id:       "P0001",
		Name:     "iPhone 12",
		ImageUrl: "https://example.com/images/iphone-12.jpg",
	}

	bytes, _ := json.Marshal(product)

	fmt.Println(string(bytes))
}
