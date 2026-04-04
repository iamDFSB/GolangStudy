package main

import "fmt"

func funcao1(){
	fmt.Println("Executando func1")
}

func funcao2(){
	fmt.Println("Executando func2")
}

func isApproved(n1, n2  float32) bool{
	defer fmt.Println("Média calculada") // Esse vai ser executado primeiro do que o funcao1, pois a ordem é LIFO
	fmt.Println("Aluno está aprovado?")
	media := (n1 + n2) / 2

	if media > 6{
		return true
	}

	return false
}

func main(){
	defer funcao1() // adiar a função até o último momento possível
	funcao2()
	result := isApproved(5, 5)
	fmt.Println(result)
}