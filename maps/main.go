package main

import (
	"fmt"
)

func main() {

	// var s []int
	// s = append(s, 1)
	// s = append(s, 2)
	// fmt.Println(len(s), cap(s))

	// var m map[int]bool
	// m[1] = true
	// fmt.Println(m)

	// mapPractice()

}

func mapPractice() {
	dict := map[string]string{"101": "ravi", "102": "amit"}
	mapl := make(map[string]string)
	mapl["ravi"] = "101"
	for k, v := range dict {
		fmt.Printf("key : %s value: %s ", k, v)
	}
	//	fmt.Println(dict)
}
