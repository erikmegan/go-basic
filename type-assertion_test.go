package gobasic

/*
type assertions is the ability to change data type to the type we wanted
this feature used when we face emtpy interface
*/

import (
	"fmt"
	"testing"
)

func random() interface{} {
	return "ok"
}

func TestRandom(t *testing.T) {
	// var result interface{} = random() // another form of declaration
	result := random()
	resultString := result.(string) //conversion to string
	fmt.Println("this is resultString: " + resultString)

	// type assertions will do panic if the type doesnt match
	// then its better using switch expression to cover it
	resultUsingSwitch := random()
	switch value := resultUsingSwitch.(type) {
	case string:
		fmt.Println("the value is string: ", value)
	case int:
		fmt.Println("the value is int: ", value)
	case bool:
		fmt.Println("the value is boolean: ", value)
	default:
		fmt.Println("the value unknown type")
	}
}
