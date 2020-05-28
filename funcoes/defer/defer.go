package main

import "fmt"

func obterValorAprovado(numero int) int {
	defer fmt.Println("fim")

	if numero < 5000 {
		fmt.Println("Valor alto")
		return numero
	} else {
		fmt.Println("Valor baixo")
		return numero
	}
}

func main() {
	fmt.Println(obterValorAprovado(6000))
}
