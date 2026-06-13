package fundation

import "fmt"

// [Pattern]: [Go Primitive] Integer/Floating-point/Imagenary literals, https://go.dev/ref/spec#Integer_literals
func NumberPresent() {
	// decimal := 42
	// million := 1_000_000
	// binary := 0b10_1010
	// octal := 0o52
	// hexa := 0x2A
	// for i := 0; i < binary; i++ {
	// }
	// fmt.Println(decimal, million, binary, octal, hexa)

	// fmt.Println(.12345e+5) // e<num> = 10^<num>
	fmt.Println(0x5p+2)  // p<num> = 2^<num>
	fmt.Println(0x.1p+1) // 0x.1p+1 = 0x.2 => 2/16 =1/8 = 0.125
}
