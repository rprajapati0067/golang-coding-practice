package main

import "fmt"

func insertElemAtItsRightPosition(s []int, target int) []int {

	s = append(s, 0)
	n := len(s) - 1
	for i := 0; i < len(s); i++ {
		if s[i] < target && s[i+1] > target {
			for j := n; j > i; j-- {

				s[j] = s[j-1]

			}
			s[i+1] = target
		} else if target > s[len(s)-2] {
			s[len(s)-1] = target
		}

	}
	return s
}

func main() {
	s := []int{1, 2, 3, 4, 6, 7}
	target := 5

	fmt.Println(insertElemAtItsRightPosition(s, target))

}
