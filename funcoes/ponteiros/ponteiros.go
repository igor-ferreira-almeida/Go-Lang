package main

import "fmt"

func incremento1(n int) {

}

func incremento2(n *int) {
	*n++
}

func main() {
	n := 1
	incremento1(n)
	fmt.Println(n)

	incremento2(&n)
	fmt.Println(n)
}
