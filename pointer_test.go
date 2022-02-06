package gobasic

import (
	"fmt"
	"testing"
)

func TestPassByValue(t *testing.T) {
	address1 := Address{"jogja", "DIY", "indo"}
	fmt.Println("this is address1: ", address1)

	address2 := address1
	address2.City = "jakarta" // doesnt change address1 (pass by value)
	fmt.Println("this is address2: ", address2)
}

func TestPassByRef(t *testing.T) {
	address1 := Address{"jogja", "DIY", "indo"}
	fmt.Println("this is address1: ", address1)

	address2 := &address1
	// var address2 *Address = &address1 // another form of declaration pointer
	address2.City = "jakarta" // change the reference (address1) only the city field
	fmt.Println("this is address1 after change: ", address1)
	fmt.Println("this is address2: ", address2)
}

func TestPointerStarSign(t *testing.T) {
	address1 := Address{"jogja", "DIY", "indo"}
	fmt.Println("this is address1: ", address1)

	address2 := &address1
	address2 = &Address{ // pass by value struct because all of variable part in the struct has change
		City:     "jakarta",
		Province: "DKI",
		Country:  "indo",
	}
	fmt.Println("this is address1 after change: ", address1) // doesnt change address1
	fmt.Println("this is address2: ", address2)
	fmt.Println("==================================")

	address3 := &address1
	*address3 = Address{ // change all struct which pointing Address1
		City:     "jakarta",
		Province: "DKI",
		Country:  "indo",
	}

	fmt.Println("this is address1 after change using star: ", address1)
	fmt.Println("this is address2 after change using star: ", address2) // address2 will affected because pointing address1 data
	fmt.Println("this is address3:  ", address3)
	fmt.Println("==================================")

	var address4 *Address = new(Address) // create new pointer
	address4.City = "address 4"
	fmt.Println(address4)
	fmt.Println("==================================")
	var address5 Address
	address5.City = "address 5" // what is the difference ?
	fmt.Println(address5)
}
