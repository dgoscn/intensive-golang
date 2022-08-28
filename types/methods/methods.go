package main

import (
	"fmt"
	"strings"
)

type person struct {
	name      string
	last_name string
}

func (p person) getFullName() string {
	return p.name + " " + p.last_name
}

func (p *person) setFullName(fullName string) {
	parts := strings.Split(fullName, " ")
	p.name = parts[0]
	p.last_name = parts[1]
}

func main() {
	p1 := person{"Ednaldo", "Pereira"}
	fmt.Println(p1.getFullName())

	p1.setFullName("Whata IsTheBrother")
	fmt.Println(p1.getFullName())
}
