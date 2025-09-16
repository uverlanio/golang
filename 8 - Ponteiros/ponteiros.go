package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var var1 int = 10
	var var2 int = var1
	fmt.Println(var1, var2)

	var2++
	fmt.Println(var1, var2)

	// Ponteiro é uam referência de memória - não guarda valor
	var var3 int
	var ponteiro *int

	var3 = 100
	ponteiro = &var3 //&var enderço de memória

	fmt.Println(var3, ponteiro) //desferenciação - pegando o valor do endereço de memória

	var3 = 150

	fmt.Println(var3, *ponteiro)

}
