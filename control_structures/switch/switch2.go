package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch { //true
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 18:
		fmt.Println("Evening")
	default:
		fmt.Println("Night")
	}
}
