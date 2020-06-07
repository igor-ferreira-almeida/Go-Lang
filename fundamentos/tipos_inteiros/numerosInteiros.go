package main

import (
	"fmt"
	"math"
	"math/big"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro", reflect.TypeOf(32000))

	// Apenas números positivos (conjunto dos Naturais)
	var a uint8 = 5
	var b uint16 = 20
	var c uint32 = 500
	var d uint64 = 35441125

	var a2 byte = 8

	fmt.Println("Literal byte", reflect.TypeOf(a2))
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println("")
	fmt.Println("Números máximos de cada tipo de int")
	fmt.Println("------------------------------------")
	fmt.Println("int8", math.MaxInt8)
	fmt.Println("int16", math.MaxInt16)
	fmt.Println("int32", math.MaxInt32)
	fmt.Println("int64", math.MaxInt64)

	fmt.Println("")
	fmt.Println("Números máximos de cada tipo de uint")
	fmt.Println("------------------------------------")
	fmt.Println("uint8", math.MaxUint8)
	fmt.Println("uint16", math.MaxUint16)
	fmt.Println("uint32", math.MaxUint32)
	// fmt.Println("uint64", math.MaxUint64)

	// Representa um inteiro na tabela unicode
	var e rune = 'a'
	var f rune = 'b'
	var g rune = 'c'

	fmt.Println("Literal rune é", reflect.TypeOf(e))
	fmt.Println("unicode de a:", e)
	fmt.Println("unicode de b:", f)
	fmt.Println("unicode de c:", g)

	x := big.NewInt(10)
	// y := big.NewInt(5)

	fmt.Println(x.Add(big.NewInt(10), big.NewInt(5)))
	fmt.Println(x.Mul(big.NewInt(10), big.NewInt(5)))
	fmt.Println(x.Sub(big.NewInt(9), big.NewInt(5)))
	fmt.Println(x.Div(big.NewInt(6), big.NewInt(2)))
}
