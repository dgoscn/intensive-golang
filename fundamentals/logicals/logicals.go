package main

import "fmt"

func purchases(work1, work2 bool) (bool, bool, bool) { //bought 50tv, bought 32tv or drunk ice cream
	buyTv50 := work1 && work2
	buyTv32 := work1 != work2 // exclusive or
	buyIceCream := work1 || work2

	return buyTv50, buyTv32, buyIceCream
}

func main() {
	tv50, tv32, icecream := purchases(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, Icecream: %t, Healthy: %",
		tv50, tv32, icecream, !icecream)
}
