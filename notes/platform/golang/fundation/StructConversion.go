package fundation

import "fmt"

// [Pattern]: [Go] Struct Coversion
// Only if underlying data structure are the same.
// [Ref](https://www.sohamkamani.com/golang/type-assertions-vs-type-conversions/#why-is-it-an-assertion)

type Person struct {
	name string
	age  int
}

type Child struct {
	name string
	age  int
}

// type Pet struct {
// 	name string
// }

func StrcutConvert() {
	bob := Person{
		name: "bob",
		age:  15,
	}
	babyBob := Child(bob)
	// "babyBob := Pet(bob)" would result in a compilation error
	fmt.Println(bob, babyBob)
}
