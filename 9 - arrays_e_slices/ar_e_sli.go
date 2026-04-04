package main

import (
	"fmt"
	"reflect"
)

func main(){
	// Array
	var lista[5] int
	lista2 := [4]string {"a", "b", "c", "d"}
	lista3 := [...]int{1, 2, 3, 4, 5}

	// Slice
	slice := []int{1, 2, 3, 4, 5, 6}
	slice = append(slice, 34)

	slice2 := lista2[1:3]

	fmt.Println("Arrays e Slices")
	fmt.Println(lista)
	fmt.Println(lista2)
	fmt.Println(lista3)
	fmt.Println(slice)
	fmt.Println(slice2)

	// Para pegar o tipo das variaveis
	fmt.Println(slice)
	fmt.Println(reflect.TypeOf(lista2))

	// Arrays Internos
	fmt.Println("-----------------")
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice)) // tamanho do array
	fmt.Println(cap(slice)) // capacidade do array

	fmt.Println("-----------------")
	slice4 := make([]int64, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 12)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
	
}