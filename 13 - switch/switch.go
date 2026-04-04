package main

import "fmt"

func control(num int) string{
	var txt string

	switch num{
		case 1:
			txt = "Up"
			// fallthrough // Para jogar para o próximo case, no caso o 2
		case 2:
			txt = "Down"
		case 3:
			txt = "Left"
		case 4:
			txt = "Right"
		default:
			txt = "No one"
	}

	return txt
}


func greetingTimer(num int) string{
	switch {
		case num <= 12:
			return "Bom dia"
		case (num < 18) :
			return "Boa tarde"
		default:
			return "Boa noite"
	}
}

func main(){
	fmt.Println("Switch")
	controlNumber := 1
	greetingNumber := 13

	fmt.Println(control(controlNumber))
	fmt.Println(greetingTimer(greetingNumber))
}