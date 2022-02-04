package gobasic

import (
	"fmt"
	"testing"
)

func multpleReturnValue() (name string, old int, address string) {
	name = "erik"
	old = 20
	address = "jln ini"
	return
}

func TestMultipleReturnValue(t *testing.T) {
	a, b, c := multpleReturnValue()
	fmt.Println("name: ", a)
	fmt.Println("old", b)
	fmt.Println("address: ", c)
}
