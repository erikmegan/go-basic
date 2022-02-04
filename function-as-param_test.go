package gobasic

import (
	"fmt"
	"testing"
)

func sayHelloWithFilter(name string, filter func(string) string) {
	fmt.Println("hello ", filter(name))
}

type Filter func(string) string // function type declaration
// used if there are so many parameter included

func sayHelloFuncTypeDeclare(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("hello ", nameFiltered)
}
func spamFilter(name string) (result string) {
	if name == "anjing" {
		result = "..."
	} else {
		result = name
	}
	return
}

func TestSayHello(t *testing.T) {
	sayHelloWithFilter("anjing", spamFilter)
}

func TestTypeDeclarationFunction(t *testing.T) {
	sayHelloFuncTypeDeclare("erik", spamFilter)
}
