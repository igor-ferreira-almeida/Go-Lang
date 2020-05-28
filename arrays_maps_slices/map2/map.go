package main

import "fmt"

func main() {
	funcionarios := map[string]float64{
		"José de Arimateia": 7564.15,
		"Maria Madalena":    5461.3,
		"João Batista":      10000.0,
	}

	fmt.Println(funcionarios["João Batista"])

	/*
		Não há erros ao tentar acessar um elemento inexistente ou ainda
		Também não ocorre erros ao tentar excluir um elemento que não está no map
	*/
	fmt.Println(funcionarios["Tales"])
	delete(funcionarios, "Tales")

}
