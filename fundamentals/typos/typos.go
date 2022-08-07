package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//integer numbers
	fmt.Println(1, 2, 1000)
	fmt.Println("Integer literal:", reflect.TypeOf(3200))

	//without signal - uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("Byte is:", reflect.TypeOf(b))

	//with signal - int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("The maximum int value is:", i1)
	fmt.Println("i1 type is:", reflect.TypeOf(i1))

	var i2 rune = 'a' //represents a map from the Unicode table (int32)
	fmt.Println("The rune is:", reflect.TypeOf(i2))
	fmt.Println(i2)

	//real numbers
	var x float32 = 49.99
	fmt.Println("The x type is:", reflect.TypeOf(x))
	fmt.Println("The typo of literal 49.99 is:", reflect.TypeOf(49.99)) //float64

	//boolean
	bo := true
	fmt.Println("The typo of bo is:", reflect.TypeOf(bo))
	fmt.Println(!bo)

	//string
	s1 := "Esse homi é um galado"
	fmt.Println(s1 + "!!!")
	fmt.Println("The size of string s1 is:", len(s1))

	//string with multiple lines
	s2 := `Ow Josué
	eu nunca vi
	tamanha desgraça
	quanto mais miséria tem
	mais urubu ameaça
	`
	fmt.Println("The size of string s2 is:", len(s2))

	//char doesn't exist ie var x char = 'b' - int32
	char := 'a'
	fmt.Println("The typo of char is:", reflect.TypeOf(char))
	fmt.Println(char)
}
