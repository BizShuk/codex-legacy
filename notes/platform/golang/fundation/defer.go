package fundation

import "fmt"

func DeferCall() {
	defer fmt.Println("Run at end of DeferCall()")
	fmt.Println("Running DeferCall()")
}
