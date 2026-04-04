package main

import "fmt"

func main(){
	fmt.Println("Maps")
	// map[<tipo da chave>]<tipo dos valores>

	user := map[string]string{
		"nome": "Daniel",
		"sobrenome": "Solsa",
	}

	fmt.Println(user)

	user2 := map[string]map[string]string{
		"user1": user,
		"user2": {
			"nome": "Daniel",
			"sobrenome": "Jonas",
		},
	}

	fmt.Println(user2)
	delete(user2, "user2")
	// Adicionando um user
	user2["user3"] = map[string]string{
		"nome": "Lucas",
		"sobrenome" : "Jonas",
	}
	fmt.Println(user2)
}