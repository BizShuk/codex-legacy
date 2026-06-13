package io

import (
	"fmt"
	"testing"
)

func TestOpenFile(t *testing.T) {
	OpenFile("./sampe_10timeframes.md")
	fmt.Println()
}

func TestFileRead_simlpe(t *testing.T) {
	FileRead_simlpe("./sampe_10timeframes.md")
	fmt.Println()
}
