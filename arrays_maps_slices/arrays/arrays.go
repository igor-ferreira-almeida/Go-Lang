package main

import "fmt"

func main() {
	var notas [3]float64

	fmt.Println(notas)
	notas[0], notas[1], notas[2] = 7.8, 8.5, 10
	fmt.Println(notas)

	var total float64 = 0.0

	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("A média é %.2f", media)
}
