package main

import "fmt"

func main() {

	result := countingValleys(8, "DUDDUUUUDDD")
	fmt.Println("Valley Count: ", result)

	//result := sockMerchant(9, []int32{10, 20, 20, 10, 10, 30, 50, 10, 20})
	//fmt.Println(result)

	//fmt.Println(funWithAnagrams([]string{"code","deco","anagram","gramana"}))
	//funWithAnagrams([]string{"code","deco","anagram","gramana"})
	//fmt.Println(sorted)

}

/*
 * Complete the 'funWithAnagrams' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts STRING_ARRAY text as parameter.
 */

var (
	sorted       []string
	anagramSlice []string
)

func funWithAnagrams(text []string) []string {

	// Write your code here
	for _, value := range text {
		sorted = append(sorted, string(sortChar(value)))
	}

	for i := 0; i < len(sorted); i++ {
		for j := 1; j < len(sorted); j++ {
			if sorted[i] != sorted[j] {
				anagramSlice = append(anagramSlice, sorted[i])

			}

		}
	}

	return anagramSlice
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func sortChar(input string) []byte {
	chars := []byte(input)
	for i, _ := range chars {
		for j := 0; j < len(chars)-1-i; j++ {
			if chars[j] > chars[j+1] {
				chars[j], chars[j+1] = chars[j+1], chars[j]
			}
		}
	}
	return chars

}
