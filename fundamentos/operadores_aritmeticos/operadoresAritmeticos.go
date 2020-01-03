package main

import "fmt"

func main() {
	a := 4
	b := 3

	fmt.Println("soma = ", a+b)
	fmt.Println("subtração = ", a-b)
	fmt.Println("multiplicação = ", a*b)
	fmt.Println("divisão = ", a/b)
	fmt.Println("módulo = ", a%b)

	//bitwise
	fmt.Println("AND = ", a&b)
	fmt.Println("OR = ", a|b)
	fmt.Println("XOR = ", a^b)
}
