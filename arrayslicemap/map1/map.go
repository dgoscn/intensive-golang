package main

import "fmt"

func main() {

	winners := make(map[int]string)
	winners[043377] = "Cade"
	winners[034222] = "Lolo"
	winners[131313] = "Lolo"
	fmt.Println(winners)

	for cpf, name := range winners {
		fmt.Printf("%s (CPF: %d)\n", name, cpf)
	}

	fmt.Println(winners[131313])
	delete(winners, 131313)
	fmt.Println(winners[131313])
}
