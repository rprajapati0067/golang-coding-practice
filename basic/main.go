package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println(secondLargest([]int64{4, 3, 1, 7, 9, 5}))
	fibonacci(10)
	//fmt.Println(factorial(5))
	//fmt.Println(reverse("ravi"))
	defer time.Now()
	time.Sleep(2 * time.Second)
}

func testGo() {
	fmt.Println("hi")
}

func reverse(st string) string {
	runes := []rune(st)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func factorial(num int64) int64 {
	var factorial int64 = 1
	for i := 1; int64(i) <= num; i++ {
		factorial = factorial * int64(i)
	}
	return factorial
}

func fibonacci(num int64) {
	var n1, n2, n3 int64
	n1, n2, n3 = -1, 1, 0

	for i := 1; int64(i) <= num; i++ {
		n3 = n1 + n2
		n1, n2 = n2, n3
		fmt.Printf("%d ", n3)
	}

}

func secondLargest(arr []int64) int64 {
	largest := arr[0]
	secondLargest := arr[0]
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > largest {
			secondLargest, largest = largest, arr[i]
		} else if arr[i] > secondLargest {
			secondLargest = arr[i]
		}
	}
	return secondLargest

}
