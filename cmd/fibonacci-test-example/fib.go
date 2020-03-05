package main

import (
	log "github.com/sirupsen/logrus"
)

// fib Fibonacci function definition
type fibFunc func(int64) int64

func setFormatter() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors:   false,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
}

func main() {
	setFormatter()
	log.Infof("Fibonacci recursive: fibRec(10) = %d", fibRec(10))
	log.Infof("Fibonacci iterative: fibIteration(10) = %d", fibIteration(10))
	log.Infof("Fibonacci dynamic map implementation: fibDynamicMap(10) = %d", fibDynamicMap(10))
	log.Infof("Fibonacci dynamic array implementation: fibDynamicArr(10) = %d", fibDynamicArr(10))
}

// fibRec Recursive implementation of Fibonacci
func fibRec(n int64) int64 {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibRec(n-1) + fibRec(n-2)
}

// fibIteration Fibonacci implementation which calculates Fibonacci by iterating
func fibIteration(n int64) int64 {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	n1 := int64(0)
	n2 := int64(1)
	res := int64(1)
	for i := int64(2); i < n+1; i++ {
		res = n1 + n2
		n1 = n2
		n2 = res
	}
	return res
}

// Map for dynamic implementation
var dynamic = map[int64]int64{0: 0, 1: 1}

// fibDynamic Dynamic programming implementation of Fibonacci using a map
func fibDynamicMap(n int64) int64 {
	val, ok := dynamic[n]
	if ok {
		return val
	}
	res := fibDynamicMap(n-1) + fibDynamicMap(n-2)
	dynamic[n] = res
	return res
}

// Array for dynamic implementation
var dynamicArr = make([]int64, 10000)

// fibDynamic Dynamic programming implementation of Fibonacci using an array
func fibDynamicArr(n int64) int64 {
	val := dynamicArr[n]
	if val != 0 {
		return val
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	res := fibDynamicArr(n-1) + fibDynamicArr(n-2)
	dynamicArr[n] = res
	return res
}
