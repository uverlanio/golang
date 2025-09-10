package main

import (
	"errors"
	"fmt"
)

func main() {
	var num int8 = 100
	fmt.Println(num)

	// Não aceita número negativo
	var num2 uint32 = 1000000
	fmt.Println(num2)

	// alias
	// INT32 = RUNE
	var num3 rune = 100000
	fmt.Println(num3)

	// UINT8 = BYTE
	var num4 byte = 100
	fmt.Println(num4)

	// NUMEROS REAIS

	var numReal1 float32 = 123.45
	fmt.Println(numReal1)

	var numReal2 float64 = 12300000000.45
	fmt.Println(numReal2)

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)

	var texto string
	fmt.Println(texto)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
