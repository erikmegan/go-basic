package gobasic

import (
	"fmt"
	"testing"
)

func forExample() {
	counter := 1
	for counter <= 10 {
		fmt.Println("perulangan ke: ", counter)
		counter++
	}

	for i := 0; i < 10; i++ {
		fmt.Println("perulangan ke: ", i)
	}

}

func forSlice() {
	slice := []string{"erik", "megan"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}

func forRange() {
	slice := []string{"erik", "megan"}
	for index, name := range slice {
		fmt.Println("index: ", index, " value: "+name)
	}
	for _, name := range slice {
		fmt.Println(name)
	}
}
func TestFor(t *testing.T) {
	forExample()
}

func TestForSlice(t *testing.T) {
	forSlice()
}
func TestForRange(t *testing.T) {
	forRange()
}
