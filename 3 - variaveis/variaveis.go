package main

import "fmt"

func main() {
	var variavel1 string = "Variavel 1"
	variavel2 := "Variavel 2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "oioio"
		variavel4 string = "lluulu"
	)
	fmt.Println(variavel3, variavel4)
	
	variavel5, variavel6 := "bombom", "good good" 
	fmt.Println(variavel5, variavel6)

	const ob string  = "et"
	fmt.Println(ob)
}