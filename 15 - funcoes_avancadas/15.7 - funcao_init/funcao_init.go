package main

import "fmt"

var n int


// Podemos ter uma função init por arquivo
func init(){
	fmt.Println("Executando uma função init")
	n=10
}

func main(){
	fmt.Println("Executando a função main: ", n)
}
