package json

// json object in golang using struct

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName: "erik",
		LastName:  "megan",
		Age:       20,
		Married:   true,
	}
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
