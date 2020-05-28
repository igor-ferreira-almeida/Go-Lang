package main

import "fmt"

type nota float64

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func notaParaConceito(n nota) string {
	if n.entre(9, 10) {
		return "A"
	} else if n.entre(7, 8) {
		return "B"
	} else if n.entre(5, 6) {
		return "C"
	} else {
		return "D"
	}
}

func main() {
	fmt.Println(notaParaConceito(4))
	fmt.Println(notaParaConceito(9))
	fmt.Println(notaParaConceito(7))
}
