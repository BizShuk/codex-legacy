package main

import (
	"fmt"
	"reflect"
)

func main() {
	tryTypeOf()
	tryValueOf()
	tryValueOfInstance()
}

func tryTypeOf() {
	var t func(a string) = func(a string) {
		fmt.Println("test")
	}

	var i interface{} = t

	tp := reflect.TypeOf(i)
	fmt.Println("TypeOf:", tp)
	fmt.Println("TypeOf.Kind():", tp.Kind())
}

func tryValueOf() {
	var t func(a string) = func(a string) {
		fmt.Println("test")
	}

	var i interface{} = t

	tp := reflect.ValueOf(i)
	fmt.Println("TypeOf:", tp)
	fmt.Println("TypeOf.Kind():", tp.Kind())
}

func tryValueOfInstance() {
	var t string = "original"
	var i interface{} = t

	tp := reflect.ValueOf(&i).Elem()
	t = "new"
	fmt.Println(i, t, tp)
	// TODO: [TODO] how to set on reflect ValueOf
}
