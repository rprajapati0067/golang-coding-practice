package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var urls = []string{"https://godoc.org", "https://www.packtpub.com", "https://google.com", "https://kubernetes.io"}

	for i := range urls {
		wg.Add(1)
		go retrieve(urls[i], &wg)
	}
	wg.Wait()

	//waitgroups.WaitGroups()
	//waitgroups.ConCurrencyAndParallelism()
	//fmt.Println("\nmain terminated")
	//go numbers()
	//go alphabets()
	//
	//time.Sleep(3200*time.Millisecond)
	//fmt.Println("\nmain terminated")

}

func retrieve(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	res, err := http.Get(url)
	end := time.Since(start)
	if err != nil {
		panic(err)
	}
	fmt.Println(url, res.StatusCode, end)
}

func numbers() {
	rand.Seed(time.Now().UnixNano())

	for i := 1; i < 15; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%d ", rand.Intn(20)+20)
	}

}

func alphabets() {

	for i := 'C'; i < 'G'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}

}
