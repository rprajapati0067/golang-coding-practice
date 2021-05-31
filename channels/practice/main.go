package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	messages := make(chan string)
	//signal := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println(msg)

	default:
		fmt.Println("no message received")
	}
	msg := "HI"

	select {
	case messages <- msg:
		fmt.Println("Recived ", messages)
	default:
		fmt.Println("no message received")
	}

}
