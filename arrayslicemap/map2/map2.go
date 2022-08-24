package main

import "fmt"

func main() {
	funcESalary := map[string]float64{
		"Cade O":    1111.9,
		"Lolo Lolo": 13131.13,
		"Lolo Lo":   181818.13,
	}

	funcESalary["Only Love Hurts"] = 1350.25
	delete(funcESalary, "NonExists")

	for name, salary := range funcESalary {
		fmt.Println(name, salary)
	}
}
