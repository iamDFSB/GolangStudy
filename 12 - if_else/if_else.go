package main

import "fmt"

func main(){
	fmt.Println("Estruturas de Controle")

	numero := 15

	if numero > 15{
		fmt.Println("Maior que 15")
	} else if numero == 15{
		fmt.Println("Igual a 15")
	} else {
		fmt.Println("Menor que 15")
	}

	// IF INIT

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Numero é maior que 0")
	}

	// fmt.Println(outroNumero) Não existe pois limita a variavel dentro do escopo do if-init

}
