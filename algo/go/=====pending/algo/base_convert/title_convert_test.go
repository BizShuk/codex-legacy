package base_convert

import (
	"fmt"
	"testing"
)

func TestConvertTitle(t *testing.T) {
	fmt.Printf("start\n")
	convertToTitle2(27)
	convertToTitle2(28)
	convertToTitle2(52)
	convertToTitle2(53) // BA
	convertToTitle2(26 * 26)
	convertToTitle2(1048) // NAH
	/*
		titleToNumber("ANH")
		titleToNumber("BA")
		titleToNumber("AA")
		titleToNumber("AB")
	*/
}
