package main

import (
	"fmt"
	"os"
)

func main() {

	page := os.Getpagesize()
	fmt.Print(page)
}
