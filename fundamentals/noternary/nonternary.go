package main

import "fmt"

// Golang do not support ternary operator
func getResult(score float64) string {
	//return score >= 6 ? "Approved" : "You shall not pass"
	if score >= 6 {
		return "Approved"
	}
	return "You shall not pass"
}

func main() {
	fmt.Println(getResult(6.2))
}
