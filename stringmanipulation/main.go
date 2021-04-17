package main

import (
	"fmt"
	"strings"
)

func main() {
	stringFunctions()
	//s := sets.NewSet()
	//s.Add("ravi")
	//s.Add("ravi")
	//
	//fmt.Println(s)

	//fmt.Println(reverseStringSecondApproach("Hello"))

	//fmt.Println(reverseStringFirstApproach("Hello"))
	//fmt.Println(reverseWords("How are you?"))
}



func stringFunctions() {
	fmt.Println("Good Morning")

	s := "Good bye"
	s = s + "f"

}

func reverseStringFirstApproach(word string) string {
	chars := []rune(word)
	starIndex := 0
	endIndex := len(word) - 1
	for starIndex < endIndex {
		chars[starIndex], chars[endIndex] = chars[endIndex], chars[starIndex]
		starIndex++
		endIndex--
	}
	return string(chars)
}

func reverseStringSecondApproach(word string) string {
	chars := []rune(word)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}
func reverseWords(sentance string) string {
	words := strings.Fields(sentance)
	starIndex := 0
	endIndex := len(words) - 1
	for starIndex < endIndex {
		words[starIndex], words[endIndex] = words[endIndex], words[starIndex]
		starIndex++
		endIndex--
	}
	return strings.Join(words, " ")
}
