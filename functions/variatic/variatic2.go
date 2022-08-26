package main

import "fmt"

func printApproved(approved ...string) {
	fmt.Println("Approved List")
	for i, approved := range approved {
		fmt.Printf("%d) %s\n", i+1, approved)
	}
}

func main() {
	aproved := []string{"Mary", "On", "Cross", "Ghost"}
	printApproved(aproved...)
}
