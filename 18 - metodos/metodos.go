package main

import (
	"fmt"
)

type user struct {
	name string
	age uint8
}

// Criando métodos para o struct
func (u user) save(){
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados \n", u.name)
}

func (u user) isAdult() bool{
	return u.age >= 18
}

func (u *user) beAdult(){
	if !u.isAdult(){
		diff := 18 - u.age
		u.age = u.age + diff
	}
}


func main(){
	// user1 := user{"Daniel", 21}
	// user1.save()
	// fmt.Println(user1.isAdult())

	user2 := user{"Juvenal", 13}
	user2.save() // Usando os métodos do struct logo após a instanciação
	fmt.Println(user2.isAdult())
	
	user2.beAdult()
	fmt.Println(user2.isAdult())
}