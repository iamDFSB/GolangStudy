package main

import (
	"errors"
	"fmt"
)

func main() {
	// Números
	var num int = 1000000000000000000
	fmt.Println(num)

	var num2 uint32 = 3450343359
	fmt.Println(num2)

	// alias
	// INT32 = RUNE
	var num3 rune = 1234500000 // Representam numeros da tabela ASCII
	fmt.Println(num3)

	// BYTE = UINT8
	var num4 byte = 130
	fmt.Println(num4)

	// Float(float32 e float64)
	var numeroReal float32 = 123450.34
	var numeroReal2 float64 = 12345009.34
	fmt.Println(numeroReal)
	fmt.Println(numeroReal2)

	// String
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	// Char
	char := '-'
	fmt.Println(char) //  Retorna o número do caracter na tabela ASCII

	var t int
	fmt.Println(t) // Retorna 0 mostrando que é uma variavel vazia

	// Bool (default sendo false)
	var booleano bool = true
	fmt.Println(booleano)

	// Erro <nil>
	var erro error = errors.New("Erro Interno")
	fmt.Println(erro)

}