package main

import "fmt"

type item struct {
	produtoID  int
	quantidade int
	preco      float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0

	for _, item := range p.itens {
		total += item.preco * float64(item.quantidade)
	}

	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			{
				produtoID:  1,
				quantidade: 5,
				preco:      7.45,
			},
			{
				produtoID:  2,
				quantidade: 3,
				preco:      4.5,
			},
		},
	}

	fmt.Printf("Valor total do pedido: R$ %.2f\n", pedido.valorTotal())
}
