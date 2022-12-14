package main

import (
	"fmt"
	"time"
)

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func prime_number(n int, c chan int) {
	inicio := 2
	for i := 0; i < n; i++ {
		for prime := inicio; ; prime++ {
			if isPrime(prime) {
				c <- prime
				inicio = prime + 1
				time.Sleep(time.Millisecond * 180)
				break
			}
		}
	}
	close(c)
}

func main() {
	c := make(chan int, 30)
	go prime_number(60, c)
	for prime := range c {
		fmt.Printf("%d ", prime)
	}
	fmt.Println("End!")
}
