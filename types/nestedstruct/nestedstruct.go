package main

import "fmt"

type item struct {
	productID int
	qty       int
	price     float64
}

type order struct {
	userID int
	items  []item
}

func (o order) totalValue() float64 {
	total := 0.0
	for _, item := range o.items {
		total += item.price * float64(item.qty)
	}
	return total
}

func main() {
	order := order{
		userID: 1,
		items: []item{
			item{2, 1, 12.11},
			item{3, 5, 13.13},
			item{5, 13, 14.33},
		},
	}
	fmt.Printf("Total value from the order is $%.2f", order.totalValue())
}
