package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Multiplexador")
	mainChannel := multplex(writer("Hello World!!"), writer("Good Vibes"))

	for i:=0; i <= 20; i++{
		fmt.Println(<- mainChannel)
	}
}

func multplex(entryChannel1, entryChannel2 <- chan string) <- chan string{
	exitChannel := make(chan string)

	go func(){
		for {
			select {
				case message := <- entryChannel1:
					exitChannel <- message
				case message := <- entryChannel2:
					exitChannel <- message
			}
		}
	}()

	return exitChannel
}

func writer(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Meu texto: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return channel
}