package main

import "fmt"

func main() {
	fmt.Print("Same")
	fmt.Print(" line.")

	fmt.Println(" New")
	fmt.Println("line.")

	x := 3.141516

	//fmt.Println("The x value is:" + x) - doesn't work
	xs := fmt.Sprint(x)
	fmt.Println("The x value is: " + xs)
	fmt.Println("The x value is:", x)

	fmt.Printf("The x value is %2.f", x)

	a := 1
	b := 1.9999
	c := false
	d := "galado"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)

}
