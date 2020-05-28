package main

import "fmt"

type carro struct {
	nome       string
	velocidade int
}

type ferrari struct {
	carro       // campos anônimos
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "Testarossa"
	f.velocidade = 290
	f.turboLigado = true

	fmt.Println(f)
}
