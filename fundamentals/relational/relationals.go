package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strings:", "Caju" == "Caju")
	fmt.Println("!=:", 3 != 2)
	fmt.Println("<:", 3 < 2)
	fmt.Println(">:", 3 > 2)
	fmt.Println("<=:", 3 <= 3)
	fmt.Println(">=:", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Dates:", d1 == d2)
	fmt.Println("Dates", d1.Equal(d2))

	type Person struct {
		Name string
	}

	p1 := Person{"Xibata"}
	p2 := Person{"Livre"}
	fmt.Println("People:", p1 == p2)
}
