package main

import "fmt"

func main() {
	funcionarios := map[string]map[string]float64{
		"A": {
			"Amanda": 8000.5,
			"Abel":   15000,
			"Ana":    6525.5,
		},
		"B": {
			"Beatriz": 8000.5,
			"Bruno":   15000,
		},
		"C": {
			"Carla":  8000.5,
			"Carina": 15000,
		},
	}

	// fmt.Println(funcionarios)

	for letra, funcionariosLetra := range funcionarios {
		fmt.Println()
		fmt.Printf("Funcionários com letra: %s\n", letra)
		fmt.Println("------------------------------------")
		for nome, salario := range funcionariosLetra {
			fmt.Printf("Nome: %s, salário: %.2f\n", nome, salario)
		}
	}

}
