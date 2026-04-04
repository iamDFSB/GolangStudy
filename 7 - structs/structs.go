package main

import (
	"fmt"
)

type usuario struct {
	nome string
	idade uint8
	address endereco
}

type endereco struct {
	rua string
	numero int16
}


func main(){
	fmt.Println("Arquivo structs")
	// Forma 1
	var u usuario
	u.nome = "DANIEL"
	u.idade = 21
	
	enderecoExample := endereco{rua:"Rua Pato Vermelho", numero: 3423}
	// Forma 2
	usuario2 := usuario{"Felipe", 34, enderecoExample}
	usuario3 := usuario{idade:45}
	
	fmt.Println(u)
	fmt.Println(usuario2)
	fmt.Println(usuario3)
}