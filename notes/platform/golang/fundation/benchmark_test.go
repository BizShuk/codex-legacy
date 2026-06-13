package fundation

import (
	"testing"
)

// [Pattern]: [Go Benchmark] https://blog.logrocket.com/benchmarking-golang-improve-function-performance/
// run from project root, go test -bench=. -count=5 -benchmem -run=^# <package path>
// -run=^#, no run test case
// TODO: [Go testflag] check go help testflag

// [Benchstat]:
// go test -bench=Prime -count 5 | tee old.txt
// go install golang.org/x/perf/cmd/benchstat@latest
// benchstat old.txt ...
func BenchmarkCase1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Case1(i)
	}
}
