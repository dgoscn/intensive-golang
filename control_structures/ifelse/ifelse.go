package main

import "fmt"

func printResult(note float64) {
	if note >= 7 {
		fmt.Println("Approved with note:", note)
	} else {
		fmt.Println("Not Approved with note:", note)

	}
}

func main() {
	printResult(2.4)
	printResult(8.1)
}
