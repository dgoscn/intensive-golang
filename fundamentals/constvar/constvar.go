package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var rad = 3.2 // float type inferred by golang

	// reduced way to create a var and sign value to it
	area := PI * m.Pow(rad, 2)
	//area = 1 but not declare again area := 1
	fmt.Println("The radius of circunference is:", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "catabi"
	fmt.Println(g, h, i)
}
