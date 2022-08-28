package main

import "fmt"

func main() {
	p1 := Point{2.0, 2.0}
	p2 := Point{2.0, 4.0}

	fmt.Println(peccaries(p1, p2))
	fmt.Println(Distance(p1, p2))
}
