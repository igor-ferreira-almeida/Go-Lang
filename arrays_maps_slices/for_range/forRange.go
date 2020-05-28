package main

import "fmt"

func main() {
	numeros := [...]int{7, 8, 9, 4, 5}

	for i, numero := range numeros {
		fmt.Printf("índice: %d, valor: %d\n", i, numero)
	}

	/*
		Ignorar o índice, pois se não ignorar é preciso usar,
		caso contrário haverá um erro de compilação. Este
		comportamento é semelhante ao foreach
	*/
	for _, numero := range numeros {
		fmt.Printf("valor: %d\n", numero)
	}

	/*
		Caso seja usado apenas um retorno, sempre será o índice
	*/
	for i := range numeros {
		fmt.Printf("índice: %d\n", i)
	}
}
