package main

import (
	"fmt"
	"modulo/helper"
	// "github.com/badoux/checkmail"
)

func main(){
	fmt.Println("Escrevendo no arquivo main")
	helper.Writing()
	helper.Writing2()
	// erro := checkmail.ValidateFormat("devbook@gmail.com")
	// fmt.Println(erro)
}