package gobasic

import (
	"fmt"
	"testing"
)

func createMap() {
	// var book map[string]string = make(map[string]string) // another declaration creating map
	book := make(map[string]string)
	book["title"] = "basic go-lang"
	book["author"] = "erik"
	book["ups"] = "salah"

	delete(book, "ups") // deleting map data using key
	fmt.Println(book)

}
func mapExample() {
	var person map[string]string = map[string]string{
		"name":     "erik",
		"lastName": "megans",
	}
	// fmt.Println(person)
	// fmt.Println(person["name"])
	// fmt.Println(person["lastName"])
	delete(person, "name")
	if person["name"] == "" {
		fmt.Println("kosong")
	} else {
		fmt.Println("isi")
	}
}

func TestMapExampleNumber(t *testing.T) {
	var v map[string]int = map[string]int{
		"ini": 1,
		"itu": 2,
	}
	fmt.Println(v["ini"])
	fmt.Println(v)
	// if v["sini"] == 0 {
	// 	fmt.Println("kosong")
	// } else {
	// 	fmt.Println("isi")
	// }
}
func TestMap(t *testing.T) {
	mapExample()
}
func TestCreateMap(t *testing.T) {
	createMap()
}
