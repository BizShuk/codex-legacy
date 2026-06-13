package lib

import (
	"fmt"
	"net/url"
)

// output:
// test=code
// 0
func UrlValue() {
	v := url.Values{
		"test": {"code"},
	}

	fmt.Println(v.Encode())

	var a []string
	fmt.Println(len(a))
}
