package main

import "fmt"

type pessoa struct {
	nome string
	idade int8
}


type estudante struct {
	pessoa
	curso string 
}


func main(){
	pessoa1 := pessoa{nome:"Daniel", idade:21}
	student := estudante{pessoa: pessoa1, curso:"GO"}
	fmt.Println(student)
}