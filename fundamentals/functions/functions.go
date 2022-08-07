package main

import "fmt"

func sum(a int, b int) int { //asignature of the funcion - have to return
	return a + b
}

func print(value int) { //no needed a return
	fmt.Println(value)
}

func main() {

	result := sum(4, 3)
	print(result)

}

//Create a new main.go inside the folder and execute go run *.go

/*package main

func main() {

	result := sum(4, 3)
	print(result)

}
*/
