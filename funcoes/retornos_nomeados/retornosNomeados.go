package main

import "fmt"

func trocar(p1, p2 int) (segundo, primeiro int) {
	primeiro = p1
	segundo = p2

	return
}

func main() {
	a, b := trocar(5, 2)
	fmt.Println(a, b)
}
