package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")

	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	p1 := pessoa{nome: "João", sobrenome: "Batista"}

	fmt.Println(p1.getNomeCompleto())

	p2 := pessoa{nome: "Maria", sobrenome: "Madalena"}

	fmt.Println(p2.getNomeCompleto())

	p3 := pessoa{}
	p3.setNomeCompleto("José Alencar")

	fmt.Println(p3.getNomeCompleto())
}
