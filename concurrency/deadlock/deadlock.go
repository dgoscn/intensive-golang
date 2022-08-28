package main

import (
	"fmt"
	"time"
)

func routine(c chan int) {
	time.Sleep(time.Second)
	c <- 1 //blocker operation
	fmt.Println("Only executed after the read")
}

func main() {
	c := make(chan int) //channel without buffer

	go routine(c)

	fmt.Println(<-c) //blocker operation
	fmt.Println("Read")
	fmt.Println(<-c) //deadlock
	fmt.Println("End")
}
