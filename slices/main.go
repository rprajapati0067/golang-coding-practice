package main

import "fmt"

func main() {
	fmt.Println(findDuplicate([]int{2, 4, 1, 1, 6, 6, 9, 0, 12, 11}))
}

/* finding duplicate and unique
v != 1 on line 25 (duplicate elements)
v == 1 on line 25 (unique elements)
*/

func findDuplicate(intSlice []int) []int {
	resMap := make(map[int]int)
	resultantList := []int{}
	for _, entry := range intSlice {
		if _, val := resMap[entry]; !val {
			resMap[entry] = 1
		} else {
			resMap[entry] = resMap[entry] + 1
		}
	}

	for k, v := range resMap {
		if v != 1 {
			resultantList = append(resultantList, k)
		}

	}
	return resultantList
}
