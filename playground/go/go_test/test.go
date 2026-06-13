package main

import (
	"fmt"
	"os"
)

type T struct {
	a int
	b float64
	c string
}

type Stringer interface {
	String() string
}

func test_i() string {
	var value = Stringer{} // Value provided by caller.
	switch str := value.(type) {
	case string:
		return str
	case Stringer:
		return str.String()
	}
	return ""
}

func main() {

	for i := 0; i < 10; i++ {
		defer fmt.Printf("%d", i)
	}

	//	a := [...]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	m := map[string]string{"a": "nvalid argument", "b": "1", "c": "2"}
	for k, v := range m {
		fmt.Println(k, v)

	}

	v := make([]int, 100, 1000)
	for i, w := range v {
		fmt.Println(i, w)
	}
	fmt.Println("upda")

	test_map := make([]int, 100)
	for i, j := range test_map {
		fmt.Println(i, j)
	}

	fmt.Printf("Hello %d\n", 23)
	fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
	fmt.Println("Hello", 23)
	fmt.Println(fmt.Sprint("Hello ", 23))

	t := &T{7, -2.35, "abc\tdef"}
	fmt.Printf("%v\n", t)
	fmt.Printf("%+v\n", t)
	fmt.Printf("%#v\n", t)
	fmt.Printf("%q\n", t)
	fmt.Printf("%#q\n", t)
	fmt.Printf("%T\n", t)

	fmt.Println(test_i())

	a := 1
	b := 1

	fmt.Println(a ^ b)
	fmt.Println(a & b)
}
