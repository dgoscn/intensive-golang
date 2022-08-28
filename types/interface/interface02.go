package main

import "fmt"

type sport interface {
	enableBurst()
}

type ferrari struct {
	model           string
	burstEnabled    bool
	currentVelocity int
}

func (f *ferrari) enableBurst() {
	f.burstEnabled = true
}

func main() {
	car1 := ferrari{"F50", false, 0}
	car1.enableBurst()

	var car2 sport = &ferrari{"F50", false, 0}
	car2.enableBurst()

	fmt.Println(car1, car2)
}
