package main

import "fmt"

func multiplicar(x, y int) int {
	return x * y
}

func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	resultado := exec(multiplicar, 3, 5)
	fmt.Println(resultado)
}
