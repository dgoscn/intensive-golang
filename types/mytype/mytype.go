package main

import "fmt"

type note float64

func (n note) enter(start, end float64) bool {
	return float64(n) >= start && float64(n) <= end
}

func conceptNote(n note) string {
	if n.enter(9.0, 10.0) {
		return "A"
	} else if n.enter(7.0, 8.99) {
		return "B"
	} else if n.enter(5.0, 7.99) {
		return "C"
	} else if n.enter(3.0, 4.99) {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	fmt.Println(conceptNote(9.8))
	fmt.Println(conceptNote(6.9))
	fmt.Println(conceptNote(1.3))
}
