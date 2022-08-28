package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1 //sending data to a channel (write)
	<-ch

	ch <- 2
	fmt.Println(<-ch)
}
