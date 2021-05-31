package main

import "fmt"

func main() {
	res := Solution([]int{51, 71, 17, 42})
	fmt.Println(res)

}
func sum(num int) int {
	sum := 0
	for num != 0 {
		sum = sum + num%10
		num = num / 10
	}
	return sum
}

func Solution(A []int) int {
	// write your code in Go 1.4
	fmt.Println(sum(23))
	max := -1
	for i := 0; i < len(A)-1; i++ {
		for j := 1 + i; j < len(A); j++ {
			if sum(A[i]) == sum(A[j]) {
				max = A[i] + A[j]
			}
		}

	}
	return max
}
