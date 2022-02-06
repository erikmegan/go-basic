package gobasic

import (
	"fmt"
	"testing"
)

func ChangeAddress(Address *Address) { // the param give info that type should be pointer
	Address.City = "jakarta"
}

func TestPointerFunction(t *testing.T) {
	Address1 := Address{"jogja", "DIY", "indo"}
	ChangeAddress(&Address1)
	fmt.Println("this is address1: ", Address1)

	var address2 *Address = &Address1
	ChangeAddress(address2)
	fmt.Println("this is address2: ", address2)

}
