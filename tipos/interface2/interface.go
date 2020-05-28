package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo      string
	velocidade  uint
	turboLigado bool
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	var f40 esportivo = &ferrari{"F40", 320, false}
	f40.ligarTurbo()

	testarossa := ferrari{modelo: "Testarossa", velocidade: 300, turboLigado: false}
	testarossa.ligarTurbo()

	fmt.Println(testarossa)
	fmt.Println(f40)
}
