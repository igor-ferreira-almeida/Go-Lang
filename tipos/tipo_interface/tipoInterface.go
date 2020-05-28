package main

import (
	"fmt"
	"reflect"
)

type curso struct {
	nome string
}

func main() {
	var coisa interface{}
	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)

	coisa = "teste"
	fmt.Println(coisa)

	coisa = true
	fmt.Println(coisa)

	type generico interface{}

	var x generico = "teste"

	fmt.Println(x)
	fmt.Println(reflect.TypeOf(x))

	x = 67.5

	fmt.Println(x)
	fmt.Println(reflect.TypeOf(x))

}
