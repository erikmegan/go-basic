package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	bytes, err := json.Marshal(data) // encode json
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T) {
	logJson("erik")
	logJson(1)
	logJson(true)
	logJson([]string{"erik", "megan"}) // slice
	logJson([2]int{0, 1})              // array
}
