package gobasic

/*
empty interface adalah interface yg tidak memiliki deklarasi method satupun,
hal ini membuat secara otomatis semua tipe data akan menjadi implementasi nya
*/
import (
	"fmt"
	"testing"
)

func ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "ups"
	}
}

func TestEmptyInterface(t *testing.T) {
	// var data interface{} = ups(1) // another form of declaration
	data := ups(1)
	fmt.Println(data)
}
