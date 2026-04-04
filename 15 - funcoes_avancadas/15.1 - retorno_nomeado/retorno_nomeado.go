package main

import "fmt"

func calculos(a int, b int) (soma int, sub int) {
	soma = a + b
	sub = a - b
	return
}

func main() {
	fmt.Println(
		calculos(3, 4),
	)
}