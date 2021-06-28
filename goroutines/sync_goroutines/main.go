package main

import "fmt"

var chan1 = make(chan bool, 1)
var chan2 = make(chan bool)
var index = make(chan bool, 1)

func func1() {
	for i := 1; i <= 10; i++ {
		// 2, take the element in chan1, continue to run, print two numbers
		<-chan1
		fmt.Println("Hello ")
		//	fmt.Print(i + 1)
		// 3, put an element into chan2 and wait to remove the element
		chan2 <- true
	}
}

func func2() {
	for i := 1; i <= 10; i++ {
		// 4, chan2 takes out the element and executes printing two characters,
		<-chan2
		fmt.Println("World ")
		//	fmt.Print(string(i + 1))
		// 5, chan1 receives an element, enters the blocking state, waits for the element to be removed, enters step 2, step 2345 and loops until printing is completed
		chan1 <- true
	}
	// 6, end the loop, the index channel receives an element, enters the blocked state, and waits for removal
	index <- true
}

func main() {
	// 1, put a value in chan1, enter the blocking state, wait for the element to be taken
	chan1 <- true
	go func1()
	go func2()
	// 7, the index channel takes the element and continues to execute
	<-index
	// Result: 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ
}
