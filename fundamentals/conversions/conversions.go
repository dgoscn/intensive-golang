package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	score := 6.9
	finalScore := int(score)
	fmt.Println(finalScore)

	//warning
	//fmt.Println("Test " + string(97)) //ascii value to concatanate

	//int to string
	fmt.Println("Test" + strconv.Itoa(123))

	//string to int
	num, _ := strconv.Atoi("123") // num relates to a number and _ to a possible error
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("True")
	}
}
