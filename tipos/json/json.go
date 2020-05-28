package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct para json
	p1 := produto{1, "Notebook", 2543.15, []string{"promoção", "eletrônico"}}
	fmt.Println(p1)

	p1Json, _ := json.Marshal(p1)
	fmt.Println(p1Json)
	fmt.Println(string(p1Json))

	// json para struct
	var p2 produto
	jsonStruct := `{"id": 2, "nome": "Caneta", "preco": 1.45, "tags": ["papelaria"]}`
	json.Unmarshal([]byte(jsonStruct), &p2)
	fmt.Println(p2.Tags[0])
}
