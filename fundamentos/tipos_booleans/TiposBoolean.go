package main

import (
	"fmt"
	"reflect"
)

func main() {

	b := true

	var b2 bool = false

	fmt.Println(b)
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(b2)
}
