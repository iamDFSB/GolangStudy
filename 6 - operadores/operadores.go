package main

import "fmt"

func main(){
	// Aritmético
	soma := 1 + 2
	subtracao := 2 - 1
	divisao := 10 /4
	mult := 10 * 5
	restoDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, mult, restoDivisao)

	var num1 int16 = 123
	var num2 int16 = 123
	soma2 := num1 + num2
	fmt.Println(soma2)

	// Operadores Relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	// Operadores Lógicos
	verdadeiro := true
	falso := false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

	// Operadores Unários
	numero := 10
	numero++
	numero--
	numero += 12 // numero = numero + 12   
	numero *= 2
	numero /= 2
	numero %= 2
	fmt.Println(numero)

	// Operador Ternário
	var txt string
	if numero > 4 {
		txt = "Maior que 5"
	} else {
		txt = "Menor que 5"
	}
	fmt.Println(txt)
}