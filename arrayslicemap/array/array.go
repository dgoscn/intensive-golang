package main

import "fmt"

func main() {
	//homogeneous and static
	var note [3]float64
	fmt.Println(note)

	note[0], note[1], note[2] = 1.3, 3.1, 1.3
	fmt.Println(note)

	total := 0.0
	for i := 0; i < len(note); i++ {
		total += note[i]
	}

	avga := total / float64(len(note))
	fmt.Printf("Avg %.2f\n", avga)
}
