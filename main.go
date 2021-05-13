package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func executeURL(urls []string) {
	for _, v := range urls {
		wg.Add(1)
		res, err := http.Get(v)

		if err != nil {
			fmt.Println(v, res.Status)
		}

	}

}

func main() {

	//channel := make(chan string)
	urls := []string{"https://www.google.com/", "https://developers.google.com/protocol-buffers/docs/gotutorial", "https://golang.org/"}

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(url, res.Status)
		}(url)
	}

	wg.Wait()

	// total := 30

	// numArray := [6]int{10, 40, 20, 30, 13}

	//for _, v1 := range numArray {
	//	for _, v2 := range numArray {
	//
	//	}
	//}
	//

	//[6]int{10, 40, 20, 30, 13}

	// for i := 0; i < len(numArray); i++ {
	// 	for j := i; j < len(numArray); j++ {

	// 		if i == j {
	// 			return
	// 		} else {
	// 			if numArray[i]+numArray[j] == total {
	// 				fmt.Println(numArray[i], numArray[j])
	// 				break
	// 			}
	// 		}

	// 	}

	// }

}
