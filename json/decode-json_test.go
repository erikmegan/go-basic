package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	jsonReq := `{"FirstName":"erik","LastName":"megan","Age":20,"Married":true}`
	jsonBytes := []byte(jsonReq)

	Customer := Customer{} // or u can declare pointer & here
	err := json.Unmarshal(jsonBytes, &Customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(Customer)
	fmt.Println(Customer.FirstName)
	fmt.Println(Customer.LastName)
}
