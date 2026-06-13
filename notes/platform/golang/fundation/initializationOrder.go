package fundation

import "fmt"

// [Pattern]: [Go] Initialization Order <https://golang.org/ref/spec#Package_initialization>
var (
	InitA = InitC + InitB // == 9
	InitB = InitFunc()    // == 4
	InitC = InitFunc()    // == 5
	InitD = 3             // == 5 after initialization has finished
)

func InitFunc() int {
	InitD++
	return InitD
}

func InitOrderOnVariables() {
	fmt.Printf("InitA: %v, InitB: %v, InitC: %v, InitD: %v", InitA, InitB, InitC, InitD)
}

// [Notice]: [Go] Unspeicified initialization
var InitX = I(T{}).ab()  // InitX has an undetected, hidden dependency on InitI and InitJ
var _ = initSideEffect() // [Tip]: unspecified, may run b/a x initialization. unrelated to InitX, a, or b
var InitI = InitJ
var InitJ = 42

type I interface{ ab() []int }
type T struct{}

func (T) ab() []int { return []int{InitI, InitJ} }

func initSideEffect() bool {
	// Do some operation for InitI or InitJ
	return true
}
