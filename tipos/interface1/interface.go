package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	descricao string
	preco     float64
}

// interfaces são implementadas implicitamente
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return p.descricao
}

func imprimir(imprimivel imprimivel) {
	imprimivel.toString()
}

func main() {

	pessoa := pessoa{nome: "João", sobrenome: "Batista"}
	produto1 := produto{descricao: "Headset", preco: 499}
	var produto2 imprimivel = produto{descricao: "Headset", preco: 753}

	fmt.Println(pessoa.toString())
	fmt.Println(produto1.toString())

	fmt.Println(produto2.toString())
}
