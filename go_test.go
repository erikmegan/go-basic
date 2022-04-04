package gobasic

import (
	"fmt"
	"strings"
	"testing"
)

func TestGo(t *testing.T) {
	// var n int64 = 126
	// biner := strconv.FormatInt(n, 2)
	// fmt.Println(biner)

	s := strings.Split("asdfg", "")
	fmt.Println(s)
}
func simpleArraySum(ar []int32) int32 {
	// Write your code here
	var res int32
	for _, arr := range ar {
		res += arr
	}
	return res
}
