package main

import "fmt"

type product struct {
	name     string
	price    float64
	discount float64
}

// method: receiver function (receptor)
func (p product) priceWithDiscount() float64 {
	return p.price * (1 - p.discount)
}

func main() {
	var product1 product
	product1 = product{
		name:     "Latinha",
		price:    1.79,
		discount: 0.04,
	}
	fmt.Println(product1, product1.priceWithDiscount())

	product2 := product{"Cajuzinho", 8.59, 0.89}
	fmt.Println(product2.priceWithDiscount())
}
