package gobasic

import (
	"fmt"
	"testing"
)

func getGoodBye(name string) (result string) {
	result = "good bye" + name
	return
}

func TestGetGoodBye(t *testing.T) {
	sayGoodBye := getGoodBye
	result := sayGoodBye("erik")
	fmt.Println(result)
}
