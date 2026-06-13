package fundation

import "fmt"

// [Tip]: [Go] init func
// func init will be run ahead of main func.
// It could have multiple init functions.
func init() {
	fmt.Println("This is init func...")
}
