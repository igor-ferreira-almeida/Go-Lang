package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1

	// Não existe loop while em Go
	for i < 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println()
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	// Loop infinito
	for {
		fmt.Println("Imprimindo para sempre...")
		time.Sleep(time.Second)
	}
}
