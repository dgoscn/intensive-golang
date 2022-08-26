package main

import "fmt"

func inc1(n int) {
	n++
}

// a pointer is represented by *
func inc2(n *int) {
	// * is used when you want to dereference, i.e,
	// have acess to the value where the pointer points
	*n++ // it is not a pure function anymore, when you use reference -> line 23
}

func main() {
	n := 1

	inc1(n) //by value
	fmt.Println(n)

	// when used &, you take the address of the variable
	inc2(&n)
	fmt.Println(n)
}
