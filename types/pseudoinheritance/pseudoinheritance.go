package main

import "fmt"

type car struct {
	name            string
	currentVelocity int
}

type gurgel struct {
	car     //anonymous field
	burstOn bool
}

func main() {
	g := gurgel{}
	g.name = "Itaupu"
	g.currentVelocity = 0
	g.burstOn = true

	fmt.Printf("Is Gurgel %s with burst enabled? %v \n", g.name, g.burstOn)
	fmt.Println(g)
}
