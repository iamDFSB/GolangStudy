package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("Generator")
	canal := writer("Hello World!!")
	count := 10
	for i := 0; i < count; i++ {
		fmt.Println(<- canal)
	}

}

func writer(text string) <- chan string{
	canal := make(chan string)

	go func (){
		for {
			canal <- fmt.Sprintf("Meu texto: %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}