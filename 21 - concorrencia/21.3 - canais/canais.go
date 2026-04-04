package main

import (
	"fmt"
	"time"
)

// func main(){
// 	// Canal
// 	canal1 := make(chan string)
// 	canal2 := make(chan string)
	
// 	go write("Hello World", canal)
// 	for message := range canal{
// 		fmt.Println(message)
// 	}

// 	// for {
// 	// 	message, isOpen := <-canal // esperando que o canal receba um valor
// 	// 	fmt.Println(message)
// 	// 	if !isOpen {
// 	// 		break
// 	// 	}
// 	// }

// 	fmt.Println("Fim do Programa")
// }

// func write(txt string, canal chan string){
// 	defer close(canal)
// 	for i := 0; i < 5; i++{
// 		canal <- txt // mandando um valor para dentro do canal
// 		time.Sleep(time.Second)
// 	}
	
// }

// func write2(txt string, canal1 <- chan string, canal2 chan <- string){

// }



func main() {
	canal1 := make(chan string)
	canal2 := make(chan string)

	go write("Hello World", canal1)
	go write2(canal1, canal2)

	for message := range canal2 {
		fmt.Println(message)
	}

	fmt.Println("Fim do Programa")
}

func write(txt string, canal chan<- string) {
	defer close(canal)
	for i := 0; i < 5; i++ {
		canal <- txt
		fmt.Printf("Peguei o número %d\n", i)
		time.Sleep(time.Second)
	}
}

// canal1 -> somente leitura
// canal2 -> somente escrita
func write2(canal1 <-chan string, canal2 chan<- string) {
	defer close(canal2)

	for msg := range canal1 {
		canal2 <- "Processado: " + msg
	}
}