package main

import "fmt"

type course struct {
	name string
}

func main() {
	var thing interface{}
	fmt.Println(thing)

	type dynamic interface{}

	var thing2 dynamic = "Sup"
	fmt.Println(thing2)

	thing2 = true
	fmt.Println(thing2)

	thing2 = course{"Golang: code lang by google"}
	fmt.Println(thing2)
}
