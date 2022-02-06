package gobasic

import (
	"fmt"
	"testing"
)

type Man struct {
	name string
}

func (man *Man) Married() {
	man.name = "Mr. " + man.name
	fmt.Println("di method: ", man.name)
}

func TestStructMethod(t *testing.T) {
	erik := Man{
		name: "erik",
	}
	erik.Married()         // no need use & in the struct method
	fmt.Println(erik.name) // will echo "Mr. erik"

}
