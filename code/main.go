package main

import (
	"fmt"
	"runtime"
)

func main() {
	//var temp int
	a := 5
	b := 20

	b = b + a
	a = b - a
	b = b - a


	//fmt.Printf("%d %d ", a, b)

	fmt.Println(runtime.NumCPU())

	//temp = a
	//a = b
	//b = a

	//fibo(10)
	//fmt.Printf("%d ", fiboRecursive(4))
}

func fiboRecursive(num int) int {
	if num == 0 || num == 1 {
		return 1
	}

	return fiboRecursive(num-1) + fiboRecursive(num-2)
}

func fibo(num int) {
	n1 := 0
	n2 := 1
	//n3 := 0

	for i := 0; i < num; i++ {
		n1, n2 = n2, n1+n2
		fmt.Printf("%d ", n1)
	}
}

func reverse(name string) string {
	val := []rune(name)

	var (
		startIndex = 0
		endIndex   = len(name) - 1
	)

	for startIndex < endIndex {
		val[startIndex], val[endIndex] = val[endIndex], val[startIndex]
		startIndex++
		endIndex--
	}
	return string(val)
}
