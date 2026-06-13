package fundation

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// [Pattern]: [Carrige Return] back to begining of the line. ^M(ctrl+M), Mac/Unix: \r
// [Pattern]: [Go Print fotmat] cheatsheet
// https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/
func PrintProgressBar() {
	for i := 0; i < 50; i++ {
		time.Sleep(100 * time.Millisecond)
		h := strings.Repeat("=", i) + strings.Repeat(" ", 49-i)
		fmt.Printf("\r%3.0d%%[%s]", i*100/49, h)
		os.Stdout.Sync()
	}
}

func MoveCursor() {
	fmt.Printf("start: [1]")
	for i := 0; i < 10; i++ {
		fmt.Printf("\b\b\b[%d]", i)
	}
}

func RingBell() {
	fmt.Println("\a")
}

func PrintVerticalTab() { // ?
	fmt.Printf("3\v4")
}

func PrintFormFeed() { // ?
	fmt.Printf("5\f6")
}
