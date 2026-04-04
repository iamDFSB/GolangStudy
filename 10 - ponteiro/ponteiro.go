package main

import "fmt"

func main(){
	var var1 int  = 10
	var var2 int = var1
	fmt.Println(var1, var2)
	
	var2++
	fmt.Println(var1, var2)

	// Ponteiro é uma referência de memória
	var var3 int
	var ponteiro *int

	var3 = 100
	ponteiro = &var3 // salvando o endereço de memória da variavel
	fmt.Println(var3, ponteiro)
	fmt.Println(var3, *ponteiro)
	ponteiro = &var2
	fmt.Println(var3, ponteiro)
	fmt.Println(var3, *ponteiro)

}