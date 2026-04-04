package main

import "fmt"

func main(){
	canal := make(chan string, 2)
	
	canal <- "Bom Dia"
	canal <- "Boa Noite"
		
	mensagem := <- canal
	mensagem1 := <- canal
	
	fmt.Println(mensagem)
	fmt.Println(mensagem1)
	
}