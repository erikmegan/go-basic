package gobasic

import (
	"fmt"
	"testing"
)

func array() {
	var names [3]string
	names[0] = "erik"
	names[1] = "megans"

	fmt.Println(names[0])
	fmt.Println(names[1])

	fmt.Println(names)

	var values = [3]int{
		10,
		20,
		30,
	}
	fmt.Println(values)

	fmt.Println("len names: ", len(names))
	fmt.Println("len values: ", len(values))

}

func TestArray(t *testing.T) {
	array()
}
