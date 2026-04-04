package main

import "fmt"

func inverterSinal(num int) int{
	return num * -1
}

func inverterSinalComPonteiro(num *int) {
	fmt.Println(num)
	fmt.Println(*num)
	*num = *num * - 1 
}

func main(){
	// number := 20
	// numberInverted := inverterSinal(number)
	// fmt.Println(number)
	// fmt.Println(numberInverted)

	number2 := 40
	
	inverterSinalComPonteiro(&number2)
	fmt.Println(number2)

}