package main

import (
	"reflect"
)

type Person struct {
	name string "namestr"
	age  int
}

func showtag(i interface{}) {
	switch t := reflect.TypeOf(i); t.Kind() {
	case reflect.Ptr:
		tag := t.Elem().Field(0).Tag
		print(tag)
	default:
		break
	}
}

func main() {
	var xx Person
	xx.age = 4
	xx.name = "hello"

	showtag(xx)
}
