package main

import "fmt"

func imprimirResultado(nota float64) {
	// O bloco sempre é representado com chaves ainda que seja apenas uma instrução
	if nota >= 7 {
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Reprovado")
	}
}

func main() {
	imprimirResultado(5.6)
	imprimirResultado(9.5)
}
