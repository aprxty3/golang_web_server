package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestEncodeJson(t *testing.T) {
	logJson("Aji")
	logJson(1)
	logJson(true)
	logJson([]string{"Aji", "Pras"})
}
