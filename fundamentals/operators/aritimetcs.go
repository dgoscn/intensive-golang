package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Sum =>", a+b)
	fmt.Println("Sub =>", a-b)
	fmt.Println("Div =>", a/b)
	fmt.Println("X =>", a*b)
	fmt.Println("Mod =>", a%b)

	//bitwise
	fmt.Println("And =>", a&b) // 11 & 10 = 10
	fmt.Println("Or =>", a|b)  // 11 | 10 = 11
	fmt.Println("XOR =>", a^b) // 11 ^ 10 = 01

	c := 3.0
	d := 2.0

	fmt.Println("The biggest =>", math.Max(float64(a), float64(b)))
	fmt.Println("Min =>", math.Min(c, d))
	fmt.Println("Expo =>", math.Pow(c, d))
}
