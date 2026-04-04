package main

import (
	"fmt"
	"tests/enderecos"
)

func main()  {
	fmt.Println("Começou Main")
	result := enderecos.VerifyEnderecoType("Rua dos Patinhos")
	fmt.Println(result)
}