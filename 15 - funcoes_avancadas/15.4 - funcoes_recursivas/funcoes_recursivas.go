package main

import (
	"fmt"
)

func fibo(position uint) uint {
	if position <= 1 {
		return position
	}
	return fibo(position - 2) + fibo(position - 1)
}



func main(){
	fmt.Println("Funções Recursivas")
	qtd := uint(10)
	for i := uint(0); i <= qtd; i++{
		fiboResult := fibo(i)
		fmt.Println(i, fiboResult)
	} 
}