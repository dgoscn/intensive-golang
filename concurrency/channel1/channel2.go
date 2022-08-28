package main

import (
	"fmt"
	"time"
)

//Channel it is the way of goroutines communicate
// for instance

func twoThreeFourTimes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // sending data to the channel

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go twoThreeFourTimes(2, c)

	a, b := <-c, <-c //receiving data from the channel
	fmt.Println(a, b)

	fmt.Println(<-c)
}
