package main

import "fmt"

func main() {
	AccessingClosedChannel()
	// NilChannelBlocks()
}

/*
NOTE: A nil channel always blocks weather we are sending and receiving from nil channel
*/
func NilChannelBlocks() {
	var ch chan bool
	ch <- true        // sending to nil chan block
	fmt.Println(<-ch) // receving from nil chan also blocks
}

/*
NOTE: A closed channel never blocks
Once a channel has been closed, you cannot send a value on this channel, but you can still receive from the channel.
*/
func AccessingClosedChannel() {
	ch := make(chan int16)

	close(ch)
	// Sending to closed channel always panic
	// ch <- 14

	for i := 0; i < 2; i++ {
		fmt.Println(<-ch)
	}
}
