package main

import "fmt"

func main() {
	i := 1

	var p *int = nil

	p = &i // Pegando o endereço da variável

	fmt.Println(p, *p, i)

	*p++ // Mudando o valor alocado na memória

	/** Go não possui aritmética de ponteiros
	p++
	*/

	fmt.Println(p, *p, i, &i)
}
