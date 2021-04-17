package main

import "fmt"

func main() {
	var array = new([6]int)

	length := 0

	for i := 0; i < 3; i++ {
		array[i] = i * i
		length++
	}
	fmt.Println("Capacity", len(array))
	fmt.Println("Length", length)

}
