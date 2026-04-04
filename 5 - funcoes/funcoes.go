package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculo(n1, n2 int8) (int8, int8){
	soma := n1 + n2
	sub := n1 - n2
	return soma, sub
}

func main() {
	result := somar(1, 2)
	fmt.Println(result)

	var f = func (txt string) string{
		fmt.Println(txt)
		return txt
	}

	resultado := f("OPA")
	fmt.Println(resultado)
	
	resultSoma, _ := calculo(8, 4)
	fmt.Println(resultSoma)

}