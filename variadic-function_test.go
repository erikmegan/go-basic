package gobasic

import (
	"fmt"
	"testing"
)

func variadicSum(numbers ...int) (total int) {
	for _, i := range numbers {
		total += i
	}
	return
}

func variadicExample(operation string, numbers ...int) (result int) {
	switch operation {
	case "tambah":
		for _, i := range numbers {
			result += i
		}
	case "kurang":
		for _, i := range numbers {
			result -= i
		}
	}
	return
}

func TestVariadicSum(t *testing.T) {
	result := variadicSum(1, 2, 3, 4, 5)
	fmt.Println("total: ", result)
}
func TestVariadicExample(t *testing.T) {
	result := variadicExample("kurang", 1, 2, 3, 4, 5)
	fmt.Println(result)
}
func TestVariadicUsingSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	result := variadicExample("tambah", slice...)
	fmt.Println(result)
}
func TestVariadicUsingArray(t *testing.T) {
	arr := [5]int{1, 2, 3, 4, 5}
	// result := variadicExample("tambah", arr...)
	result := variadicSum(arr[:]...)
	fmt.Println(result)

}
