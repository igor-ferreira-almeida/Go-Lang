package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
}

type bmw struct {
}

func (bmw bmw) ligarTurbo() {
	fmt.Println("Turbo ligado")
}

func (bmw bmw) fazerBaliza() {
	fmt.Println("Baliza feita")
}

func main() {

	var carro1 esportivoLuxuoso = bmw{}

	carro1.fazerBaliza()
	carro1.ligarTurbo()

}
