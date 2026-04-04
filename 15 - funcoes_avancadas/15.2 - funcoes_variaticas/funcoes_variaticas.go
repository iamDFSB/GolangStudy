package main

import (
	"fmt"
	"reflect"
)

// Basicamente um empacotamento de valores
func sum(numbers ...int) int{
	fmt.Println(numbers)
	fmt.Println(reflect.TypeOf(numbers))
	var total int
	for _, num := range numbers{
		total+=num
	}
	return total
}

func write(texto string, times ...int){
	for _, t := range times{
		fmt.Println(texto, t)
	}
}

func main(){
	totalSoma := sum(1, 2, 3, 4, 5)
	fmt.Println(totalSoma)
	write("John Doo", 1, 2, 3, 4, 5)
}