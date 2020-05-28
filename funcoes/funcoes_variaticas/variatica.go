package main

import "fmt"

func media(numeros ...float64) float64 {
	total := 0.0

	for _, numero := range numeros {
		total += numero
	}

	return total / float64(len(numeros))
}

func main() {
	fmt.Println(media(2, 5, 9, 7, 9))
	fmt.Println(media(2, 1, 9))
}
