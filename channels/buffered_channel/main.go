package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	ch := make(chan int, 3)

	wg.Add(2)
	go func() {
		for v := range []int{1, 2, 3, 4} {
			ch <- v
		}
		wg.Done()
	}()

	go func() {
		v := <-ch

		fmt.Println(v)
		wg.Done()
	}()

	wg.Wait()
}
