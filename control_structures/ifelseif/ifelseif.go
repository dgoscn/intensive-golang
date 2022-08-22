package main

import "fmt"

func noteToConcept(n float64) string {
	if n >= 9 && 10 <= 10 {
		return "A"
	} else if n >= 7 && n < 9 {
		return "B"
	} else if n >= 5 && n < 8 {
		return "C"
	} else if n >= 3 && n < 5 {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	fmt.Println(noteToConcept(9.8))
	fmt.Println(noteToConcept(6.9))
	fmt.Println(noteToConcept(2.1))
}
