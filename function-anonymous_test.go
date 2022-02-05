package gobasic

import (
	"fmt"
	"testing"
)

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("you are blocked", name)
	} else {
		fmt.Println("welcome", name)
	}
}

func TestRegisterUser(t *testing.T) {
	list := func(name string) bool {
		return name == "anjing"
	}

	registerUser("erik", list)
}

func TestRegisterUser2(t *testing.T) {
	registerUser("anjing", func(s string) bool {
		return s == "anjing"
	})
}
