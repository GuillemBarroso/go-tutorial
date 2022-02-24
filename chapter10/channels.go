package main

import (
	"fmt"
	"time"
)

func pinger(c chan<- string) { // chan<- indicates that pinger only sends messages
	for i := 0; ; i++ {
		c <- "ping" // means send "ping" through channel c
	}
}

// if no direction (<-chan or chan<-) is specified, channel is bidirectional
func printer(c <-chan string) { // <-chan indicates that only receives messages
	for {
		msg := <-c // receive a message and store it in "msg".
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong" // send "pong" through channel c
	}
}

func main() {
	var c chan string = make(chan string)

	/* channel synchronizes two goroutines. When pinger attemps to send a message on
	the channel it will wait until printer is ready to receive the message.
	*/
	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
