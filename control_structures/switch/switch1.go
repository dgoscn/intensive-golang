package main

import "fmt"

func noteToConcept(n float64) string {
	var note = int(n)
	switch note {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Invalid"
	}
}

func main() {
	fmt.Println(noteToConcept(9.8))
	fmt.Println(noteToConcept(6.9))
	fmt.Println(noteToConcept(2.1))
}
