package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float32 = 49.99
	var y float64 = 2.665
	z := 3.1415 // Infere-se o padrão float64

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.TypeOf(y))
	fmt.Println(reflect.TypeOf(z))
}
