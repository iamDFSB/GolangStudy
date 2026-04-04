package main

import (
	"fmt"
	// "time"
)


func main(){
	fmt.Println("Loops")
	
	////  É tipo um while 
	/*

	num := 0
	for num < 3 {
		num++
		fmt.Println("Incrementando num")
		time.Sleep(time.Second)
	}

	fmt.Println(num)

	*/

	//// For-loop normal
	/*

	for j:=0; j <= 6; j+=3{
		fmt.Println("Incrementando j: ", j)
		time.Sleep(time.Second)
	}

	*/

	//// Iterar por um array
	/*

	nomes := [3]string{"Joao", "Vera", "Magnólia"}
	
	for _, nome := range nomes{
		time.Sleep(time.Second)
		fmt.Println(nome)
	}

	*/

	//// Iterar sobre uma String
	/*

	for i, letra := range "PALAVRA"{
		fmt.Println(i, string(letra)) // Necessário fazer o casting para string para não retornar o valor ASCII
	}

	*/

	//// Iterar sobre um Map/Dicionário

	user := map[string]string{
		"p1": "Pessoa 1",
		"p2": "Pessoa 2",
		"p3": "Pessoa 3",
	}

	for key, value := range user{
		fmt.Println(key, value)
	}


}