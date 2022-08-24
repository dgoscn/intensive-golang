package main

import "fmt"

func main() {
	funcPerWord := map[string]map[string]float64{
		"G": {
			"Galego Birosca": 131313.50,
			"Gade o Lolo":    131313.50,
		},
		"J": {
			"JR JR":   1001000.20,
			"Jacarys": 9999.37,
		},
		"P": {
			"Piranha": 1090909.1,
		},
	}

	delete(funcPerWord, "P")

	for word, funcs := range funcPerWord {
		fmt.Println(word, funcs)
	}
}
