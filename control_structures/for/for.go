package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Println("%d", i)
	}

	fmt.Println("\nMisturando... ")
	for i := i; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Even")
		} else {
			fmt.Print("Odd")
		}
	}

	for {
		fmt.Println("Forever ")
		time.Sleep(time.Second * 5)
	}
}
