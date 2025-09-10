package main

import "fmt"

func main() {
	var var1 string = "var 1"
	var2 := "var 2"
	fmt.Println(var1)
	fmt.Println(var2)

	const const1 string = "const 1"
	fmt.Println(const1)

	var3, var4 := "var3", "var4"

	var3, var4 = var4, var3

	fmt.Println(var3)
	fmt.Println(var4)
}
