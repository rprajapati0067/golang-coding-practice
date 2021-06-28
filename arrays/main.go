package main

import "fmt"

//		1,       1,            1,
// [20,30,10,12,14,5,25,26,29,30,15,12]
// [10,5,2] == [2,5,10]

func main() {
	res := countGrowth([]int{20, 30, 10, 12, 14, 5, 25, 26, 29, 30, 15, 12})
	fmt.Println("Result: ", res)
}

func countGrowth(metrices []int) int {
	counter := 0

	for i := 1; i < len(metrices)-1; i++ {
		if metrices[i] > metrices[i+1] {
			if metrices[i] > metrices[i-1] {
				counter++
			}
		}
	}
	return counter
}
