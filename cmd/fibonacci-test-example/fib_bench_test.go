package main

import (
	"testing"
)

const (
	fibNum int64 = 30
)

func BenchmarkFibRec(b *testing.B) {
	runBench(b, fibRec)
}

func BenchmarkFibIteration(b *testing.B) {
	runBench(b, fibIteration)
}
func BenchmarkFibDynamicMap(b *testing.B) {
	runBench(b, fibDynamicMap)
}

func BenchmarkFibDynamicArr(b *testing.B) {
	runBench(b, fibDynamicArr)
}

func runBench(b *testing.B, f fibFunc) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		f(fibNum)
	}
}
