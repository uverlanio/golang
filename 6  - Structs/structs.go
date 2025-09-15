package main

import "fmt"

type endereco struct {
	logradouro string
	numero     uint8
}

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario

	u.nome = "Zoe"
	u.idade = 1
	enderecoZoe := endereco{"Rua da Zoe", 1}
	u.endereco = enderecoZoe

	fmt.Println(u)

	enderecoUver := endereco{"Rua do Uver", 1}
	usuario2 := usuario{"Uver", 37, enderecoUver}

	fmt.Println(usuario2)

	enderecoJesus := endereco{"Rua de Jesus", 1}
	usuarioJesus := usuario{nome: "Jesus", endereco: enderecoJesus}
	fmt.Println(usuarioJesus)
}
