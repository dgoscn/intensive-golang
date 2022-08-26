package main

import "fmt"

func getApprovedValue(number int) int {
	defer fmt.Println("In the end...")
	if number > 13 {
		fmt.Println("LulaLa...")
	}
	fmt.Println("NotLulaLa")
	return number
}

func main() {
	fmt.Println(getApprovedValue(1222))
	fmt.Println(getApprovedValue(01))
}
