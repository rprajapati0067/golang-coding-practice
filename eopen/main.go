package main

import "fmt"

var lookup map[int]string

func init() {
	lookup = make(map[int]string, 8)
	lookup[2] = "ABC"
	lookup[3] = "DEF"
	lookup[4] = "GHI"
	lookup[5] = "JKL"
	lookup[6] = "MNO"
	lookup[7] = "PQRS"
	lookup[8] = "TUV"
	lookup[9] = "WXYZ"
}

// ABC DEF GHI

// func separateNum(n int) []string { // 234

// }

func lookupTheString(num int) []string {
	var s []string
	var numArr []int
	i := 0
	rem := 0
	for num > 0 {
		rem = num % 10
		num = num / 10
		numArr = append(numArr, rem)
		//i++

		for k, v := range lookup {
			if k == rem {
				s = append(s, v)
				i++
			}
		}
	}

	return s
}

func reverse(strSlice []string) []string {

	startIndex := 0
	endIndex := len(strSlice) - 1

	for startIndex < endIndex {
		strSlice[startIndex], strSlice[endIndex] = strSlice[endIndex], strSlice[startIndex]
		startIndex++
		endIndex--
	}
	return strSlice
}

// [ABC
// DEF

// GHI]

// [105 440]
// [170  180]
// [130 550] 2

// [105 440]
// [170  180]
// [130 550] 2

// [105 440]
// [170  180]
// [130 550]

func combinations(num int) []string {
	strSlice := lookupTheString(num)
	var finalArr []string
	var combi string
	actualArr := reverse(strSlice)
	for i := 0; i < len(actualArr); i++ {
		st := actualArr[i]

		alpha := string(st[i])
		combi += alpha
		finalArr = append(finalArr, combi)
	}
	return finalArr
}

func main() {

	fmt.Println(combinations(234))

	//fmt.Println(combinations(2))
}
