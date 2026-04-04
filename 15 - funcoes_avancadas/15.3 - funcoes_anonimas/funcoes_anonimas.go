package main

import "fmt"


func main() {
	retorno := func(text string) string{
		fmt.Println(text)
		return fmt.Sprintf("Recebido -> %s", text)
	}("Hello")
	
	fmt.Println(retorno)
}