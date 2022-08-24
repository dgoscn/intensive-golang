package main

import "fmt"

func main() {
	numbers := [...]int{1, 2, 3, 4, 5} //[...]The count of numbers of elements is made by the compiler

	for i, numbers := range numbers {
		fmt.Printf("%d) %d\n", i+5, numbers)
	}

	for _, num := range numbers { //ignoring the indexes number with _
		fmt.Println(num)
	}
}
