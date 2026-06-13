package fundation

import "fmt"

type fmtInterface struct{}

func (this *fmtInterface) Error() string {
	return fmt.Sprintln("This is in Error()")
}

func (this fmtInterface) String() string {
	return fmt.Sprintln("This is in String()")
}

// [Notice]: [Go] fmt priority Error() > String()
// This will take Stringer interface(String method) to print.
// However, if there is error interface(Error method) implement. error interface has higher priority.

func FmtInterface() {
	f := &fmtInterface{}
	fmt.Print(f)
}
