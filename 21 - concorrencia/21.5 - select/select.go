package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("Select Code")
	canal1, canal2 := make(chan string), make(chan string)

	go func(){
		for {
			time.Sleep(time.Second / 2)
			canal1 <- "Canal 1"
		}
	}()
	
	go func(){
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {
		select {
			case mensagemCanal1 := <- canal1:
				fmt.Println(mensagemCanal1)
			case mensagemCanal2 := <- canal2:
				fmt.Println(mensagemCanal2)
		}
		
		// Forma mais lenta
		// mensagemCanal1 := <- canal1
		// mensagemCanal2 := <- canal2 
		// fmt.Println(mensagemCanal1)
		// fmt.Println(mensagemCanal2)
	}
	
}