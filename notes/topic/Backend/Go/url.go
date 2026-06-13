package main

import (
	"fmt"
	"net/url"
)

func main() {
	v := url.Values{
		"test": {"code"},
	}

	fmt.Println(v.Encode())

	var a []string
	fmt.Println(len(a))
}
