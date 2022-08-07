package main

import "fmt"

func main() {
	x := 1
	y := 2

	//only postfix
	x++ // x += 1 or x = x + 1
	fmt.Println(x)

	y-- //y -= 1 or y = y -1
	fmt.Println(y)

	// fmt.Println(x == y--) it is not possible 5 levels of precedence in go
}
