package receiver

import "github.com/sirupsen/logrus"

type T1 struct{}

type T2 int

func TypeAddress() {
	x := &T1{}
	y := T2(1) // [Notice]: [Go] &T2(1) => compiler error, ref: https://stackoverflow.com/questions/49270700/in-golang-why-an-interface-variable-can-get-address-of-a-struct-variable-but-can
	z := &y
	logrus.Infoln("x := &T1{} :", x, "y :=T2(1):", y, "z := &y:", *z)
}
