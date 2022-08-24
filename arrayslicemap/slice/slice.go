package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} //array
	s1 := []int{1, 2, 3}  //slice, i.e. variable size
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	//slice isn`t an array. Slice just defines a "slice" of an array
	s2 := a2[1:3]
	fmt.Println(a2, s2)

	s3 := a2[:2] // new slice, appointing to the same array
	fmt.Println(a2, s3)

	//Imagine slice as: a pointer to an element from an array
	s4 := s2[:1]
	fmt.Println(s2, s4)
}
