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

// [ABC DEF GHI ] 234
// [ADG AEH AFI BDG BEH BFI CDG CEH CFI]
func combinations(num int) []string {
	strSlice := lookupTheString(num)
	var finalArr []string
	var combi string
	actualArr := reverse(strSlice)
	k := 1
	firstString := []rune(actualArr[0])

	for i := 0; i < len(firstString); i++ {
		combi += string(firstString[i])
		for m := 0; m < len(firstString); m++ {
			for j := 1; j <= len(actualArr)-1; j++ {

				st1 := actualArr[j]
				combi += st1[m:k]

			}
			finalArr = append(finalArr, combi)
			combi = string(firstString[i])
			k++
		}
		k = 1
		combi = ""

	}
	return finalArr
}

func main() {

	fmt.Println(combinations(2345))

}
