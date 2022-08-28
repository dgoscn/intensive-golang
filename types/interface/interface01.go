package main

import "fmt"

type printable interface {
	toString() string
}

type person struct {
	name      string
	last_name string
}

type product struct {
	name  string
	price float64
}

//interfaces are implicitly implemented

func (p person) toString() string {
	return p.name + " " + p.last_name
}

func (p product) toString() string {
	return fmt.Sprintf("%s - $ %.2f", p.name, p.price)
}

func print(x printable) {
	fmt.Println(x.toString())
}

func main() {
	var thing printable = person{"El", "Catita"}
	fmt.Println(thing.toString())
	print(thing)

	thing = product{"Blue Jeans", 99.99}
	fmt.Println(thing.toString())
	print(thing)

	p2 := product{"Blue Jeans", 199.99}
	print(p2)
}
