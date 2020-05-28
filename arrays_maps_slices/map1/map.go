package main

import "fmt"

func main() {
	// var aprovados map[int]String
	//Maps devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[123] = "Maria Madalena"
	aprovados[345] = "João Batista"

	fmt.Println(aprovados)

	for key, value := range aprovados {
		fmt.Printf("Nome: %s, CPF: %d\n", value, key)
	}

	fmt.Println(aprovados[345])

	delete(aprovados, 345)
	fmt.Println(aprovados[345])
}
