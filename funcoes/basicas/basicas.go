package main

import "fmt"

func f1() {
	fmt.Println("F1")
}
func f2(p1 string, p2 string) {
	fmt.Printf("Parâmetro 1: %s, parâmetro 2: %s\n", p1, p2)
}

func f3() string {
	return "Retorno da função 3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("Parâmetro 1: %s, parâmetro 2: %s\n", p1, p2)
}

func f5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {
	f1()
	f2("Igor", "Almeida")

	r3, r4 := f3(), f4("Igor", "Almeida")

	fmt.Println(r3)
	fmt.Println(r4)

	r5a, r5b := f5()

	fmt.Println(r5a, r5b)
}
