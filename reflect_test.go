package gobasic

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {
	number := 23
	reflectValue := reflect.ValueOf(number)

	fmt.Println("type data variable number: ", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variable: ", reflectValue.Int())
	}
}

func (s *Student) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)
	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}
	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("name :", reflectType.Field(i).Name)
		fmt.Println("type data :", reflectType.Field(i).Type)
		fmt.Println("value :", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}

func TestGetPropertyInfo(t *testing.T) {
	var s1 = &Student{Name: "erik", Grade: 3}
	s1.getPropertyInfo()
}
