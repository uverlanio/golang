package main

import (
	"fmt"
)

func main() {

	soma := 1 + 2
	subrtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subrtracao, divisao, multiplicacao, restoDaDivisao)

	var num1 int16 = 20
	var num2 int16 = 36

	somaNums := num1 + num2
	fmt.Println(somaNums)

	var var1 string = "string1"
	var2 := "string2"

	fmt.Println(var1, var2)

	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	verdadeiro := true
	falso := false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

	numero := 10
	numero++
	numero += 10
	fmt.Println(numero)

	var n string = "n"
	if numero > 5 {
		fmt.Println(n + " maior que 5")
	} else {
		fmt.Println(n + "menor que 5")
	}

}
