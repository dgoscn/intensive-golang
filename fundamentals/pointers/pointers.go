package main

import "fmt"

func main() {
	i := 1

	var p *int = nil
	p = &i //receiving the address of variable i
	*p++   //dereferencing (extracting the exactly value of i)
	i++

	//Go do not have pointer aritimetics
	// p++

	fmt.Println(p, *p, i, &i)
}
