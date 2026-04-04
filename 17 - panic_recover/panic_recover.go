package main

import (
	"fmt"
)

func recoverExecution(){
	if r := recover(); r != nil{
		fmt.Println("Função recuperada com sucesso")
	}else{
		fmt.Println(r)
	}
}

func isApproved(n1, n2  float32) bool{
	defer recoverExecution()
	media := (n1 + n2) / 2

	if media > 6{
		return true
	} else if media < 6 {
		return false
	}

	panic("A média é exatamente 6")
}


// Forma certa de tratar erro em Go
/*

func retorno(n1, n2 int) (int, error){
	if n1 < 2{
		return 0, fmt.Errorf("Pode não cabra")
	}
	return n1 + n2, nil
}

*/

func main(){
	fmt.Println("Preparing the Mean")
	fmt.Println(isApproved(5, 7))
	// result, error := retorno(1, 5)
	// fmt.Println(result, error)
	// fmt.Println(reflect.TypeOf(error))
}

