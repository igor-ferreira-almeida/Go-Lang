package main

import "fmt"

var soma = func(x, y int) int {
	return x + y
}

var subtracao = func(x, y int) int {
	return x - y
}

func main() {
	fmt.Println(soma(5, 9))
	fmt.Println(subtracao(5, 9))
}
