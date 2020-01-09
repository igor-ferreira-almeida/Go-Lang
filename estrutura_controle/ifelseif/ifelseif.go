package main

import "fmt"

func imprimeNota(nota float64) string {
	if nota >= 9 {
		return "A"
	} else if nota >= 7 && nota < 9 {
		return "B"
	} else if nota >= 5 && nota < 7 {
		return "C"
	} else {
		return "D"
	}
}

func main() {
	fmt.Println(imprimeNota(9))
	fmt.Println(imprimeNota(8))
	fmt.Println(imprimeNota(6))
	fmt.Println(imprimeNota(3))
}
