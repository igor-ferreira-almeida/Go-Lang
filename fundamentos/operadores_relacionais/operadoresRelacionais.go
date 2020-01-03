package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Igual =", "x" == "y")
	fmt.Println("Igual =", 5 == 5.0)
	fmt.Println("Diferente =", 5 != 4)
	fmt.Println("Maior =", 5 > 4)
	fmt.Println("Menor =", 2 < 4)
	fmt.Println("Maior ou igual =", 5 > 4)
	fmt.Println("Menor ou igual =", 2 < 4)

	data1 := time.Unix(0, 0)
	data2 := time.Unix(0, 0)

	fmt.Println("Igualdade entre datas =", data1 == data2)
	fmt.Println("Igualdade entre datas =", data1.Equal(data2))

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"João"}
	p2 := Pessoa{"João"}

	fmt.Println("Igualdade entre struct =", p1 == p2)
}
