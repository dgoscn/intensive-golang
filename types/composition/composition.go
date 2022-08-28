package main

import "fmt"

type sport interface {
	enableBurst()
}

type lux interface {
	doGoal()
}

type sportLux interface {
	sport
	lux
	//it is possible add more methods
}

type bmw7 struct{}

func (b bmw7) enableBurst() {
	fmt.Println("Burst...")
}

func (b bmw7) doGoal() {
	fmt.Println("Gooal...")
}

func main() {
	var b sportLux = bmw7{}
	b.enableBurst()
	b.doGoal()
}
