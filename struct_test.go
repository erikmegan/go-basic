package gobasic

import (
	"fmt"
	"testing"
)

// struct is template data, prototype data, class, object data
type Customer struct {
	Name, Address string
	Age           int
}

func (Customer Customer) sayHello() { // this is struct method
	fmt.Println("hello, this is struct method ", Customer.Name)
}
func (Customer Customer) sayHelloWithReturn() string { // this is struct method with return value
	return "hello, this is struct method with return string " + Customer.Name
}
func (Customer Customer) sayHelloWithParam(address string) {
	fmt.Println("hello, this is struct method with param ", address)
}
func TestCustomerStruct(t *testing.T) {
	var firstClass Customer
	firstClass.Name = "erik"
	firstClass.Address = "jln ini"
	firstClass.Age = 30

	fmt.Println(firstClass)
	fmt.Println("this is name: ", firstClass.Name)

	firstClass.sayHello()
	fmt.Println(firstClass.sayHelloWithReturn())
	firstClass.sayHelloWithParam(firstClass.Address)

}

func TestCustomerStructLiterals(t *testing.T) {
	joko := Customer{
		Name:    "joko",
		Address: "jln itu",
		Age:     30,
	}
	fmt.Print(joko)

	budi := Customer{"budi", "jakarta", 25}
	fmt.Println(budi)

}
