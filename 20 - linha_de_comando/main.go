package main

import (
	"fmt"
	"linha_de_comando/app"
	"log"
	"os"
)


func main(){
	fmt.Println("Ponto de partida")
	aplication := app.Gerar()
	if erro := aplication.Run(os.Args); erro != nil{
		log.Fatal(erro)
	}
	// go run main.go ip --host mercadolivre.com.br
}