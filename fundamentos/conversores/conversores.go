package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 54.8
	y := 5

	// Operações não funcionam sem converter para tipos semelhantes
	fmt.Println(x / float64(y))

	//Converter para inteiro
	fmt.Println(int(x))

	//Converter int para string
	fmt.Println("Concatenação de string com números: " + strconv.Itoa(y))

	//Converter string para int
	num, err1 := strconv.Atoi("2")
	fmt.Println(num + y)

	num, err2 := strconv.Atoi("x")

	fmt.Println(err1)
	fmt.Println(err2)

	/*
		Também é possível converter passando 0 ou 1 para false e true respectivamente,
		além disso, pode-se passar em caixa alta ou caixa baixa

	*/
	b, _ := strconv.ParseBool("TRUE")
	fmt.Println(b)

	if b {
		fmt.Println("Verdadeiro")
	}
}
