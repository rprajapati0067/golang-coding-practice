package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	//var data int

	wg.Add(2)

	// go func() {
	// 	defer wg.Done()
	// 	data++
	// }()

	// wg.Wait()

	// fmt.Printf("The value of data is %v\n", data)
	// fmt.Println("Done...")
	go Hello()

	go Bye()

	wg.Wait()

}

func Hello() {
	defer wg.Done()
	fmt.Println("Hello")
}
func Bye() {
	defer wg.Done()
	fmt.Println("Bye")
}
