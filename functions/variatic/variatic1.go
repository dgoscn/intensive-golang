package main

import "fmt"

func avrg(numbers ...float64) float64 {
	total := 0.0
	for _, num := range numbers {
		total += num
	}
	return total / float64(len(numbers))
}

func main() {
	fmt.Printf("Avrg: %.2f", avrg(1.3, 31.1, 3.13))
}
