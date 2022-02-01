package gobasic

/*
slice has 3 data: pointer, length, capacity
-pointer adalah pennunjuk data pertama di array pada slice
-length adalah panjang dr slice
-capacity adalah kapasitas slice, dmn length tidak boleh lebih dr capacity

slice[low:high] -> create slice from low index to index before high
slice[low:] 	-> create slice from low index to last index in array
slice[:high]	-> create slice from index 0 to index before high
slice[:] 		-> create slice from index 0 to last index in array
*/

import (
	"fmt"
	"testing"
)

func SliceAndArray() {
	// isArray := [2]int{1,2,3} // this is another form of array declare
	isArray := [...]int{1, 2, 3} // this is array
	isSlice := []int{1, 2, 3}    // this is slice

	fmt.Println("this is array: ", isArray)
	fmt.Println("this is slice: ", isSlice)
}

func createNewSlice() {
	newSlice := make([]string, 2, 5) // create new slice form new array[5]
	newSlice[0] = "slice 0"
	newSlice[1] = "slice 1"
	fmt.Println("this is new slice: ", newSlice)
}

func slice() {
	var months = [...]string{ // another form declaration array without declare how many data included
		"jan",  // 0
		"feb",  // 1
		"mar",  // 2
		"apr",  // 3
		"mei",  // 4
		"jun",  // 5
		"jul",  // 6
		"aug",  // 7
		"sept", // 8
		"oct",  // 9
		"nov",  // 10
		"dec",  // 11
	}

	var slice1 = months[4:7]
	fmt.Println("this is print slice: ", slice1)
	fmt.Println("this is len slice: ", len(slice1)) // to get length slice
	fmt.Println("this is cap slice: ", cap(slice1)) //  to get capacity slice

	slice1[0] = "meiAgain" // changing slice does affect to its array
	fmt.Println("this is array after changes from slice: ", months)

	var slice2 = months[10:] // create new slice form index 10 to end
	fmt.Println("this is slice2: ", slice2)

	var slice3 = append(slice2, "newMoon") // create new slice with adding new data in case capacity is full. append will increase capacity automatic
	// if append create new capacity, the slice will create new array (pass by value). it will not affect base array in the future if there is changes
	fmt.Println("this is slice3: ", slice3)
	fmt.Print("this is slice2 after slice3 append from slice2: ", slice2)
}

func TestSlice(t *testing.T) {
	slice()
}
func TestNewSlice(t *testing.T) {
	createNewSlice()
}
func TestArrayAndSlice(t *testing.T) {
	SliceAndArray()
}
