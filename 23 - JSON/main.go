package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type dog struct {
	Name string `json:"nome"` 
	Breed string `json:"breed"`
	Age uint `json:"-"` // Colocar "-" significa que não quer mostrar essa informação quando transformar em JSON
}

// json.Marshal
func usingMarshal() {
	dog1 := dog{Name:"Josh", Breed:"Doberman", Age:1}
	dog1Json, erro := json.Marshal(dog1)
	if erro != nil{
		log.Fatal(erro)
	}
	// fmt.Println(dog1)
	// fmt.Println(dog1Json)
	fmt.Println(bytes.NewBuffer(dog1Json))


	dogMap := map[string]string{
		"nome": "Rufus",
		"breed": "Poodle",
	}
	dogMapJson, erro := json.Marshal(dogMap)

	if erro != nil{
		log.Fatal(erro)
	}

	fmt.Println(bytes.NewBuffer(dogMapJson))
}


// json.Unmarshal
func usingUnMarshal(){
	dogJson := `{"nome":"Josh","breed":"Doberman","age":1}`
	var d dog

	erro := json.Unmarshal([]byte(dogJson), &d)
	if erro != nil{
		log.Fatal(erro)
	}

	fmt.Println(d)

	dogJson2 := `{"nome": "Rufus","breed": "Poodle"}`

	d2 := make(map[string]string)
	erro2 := json.Unmarshal([]byte(dogJson2), &d2)
	if erro2 != nil{
		log.Fatal(erro)
	}
	
	fmt.Println(d2)

}

// Handling File
func usingDataFromFile() {
	file, err := os.ReadFile("example.json")

	if err != nil{
		log.Fatal(err)
	}

	var doggy dog
	erro := json.Unmarshal(file, &doggy)

	if erro != nil{
		log.Fatal(erro)
	}

	fmt.Println(doggy)

}


func main() {
	usingDataFromFile()
}