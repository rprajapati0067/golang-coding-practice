package main

import "fmt"

func main() {
	total := 30
	numArray := [6]int{10, 40, 20, 30, 13}

	//for _, v1 := range numArray {
	//	for _, v2 := range numArray {
	//
	//	}
	//}
	//

	//[6]int{10, 40, 20, 30, 13}


	for i := 0; i < len(numArray); i++ {
		for j := i; j < len(numArray); j++ {

			if i == j {
				return
			} else {
				if numArray[i]+numArray[j] == total {
					fmt.Println(numArray[i], numArray[j])
					break
				}
			}

		}

	}

}
