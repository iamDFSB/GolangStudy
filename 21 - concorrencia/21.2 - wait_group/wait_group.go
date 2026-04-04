package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)
	
	go func() { 
		write("Bom dia")
		waitGroup.Done() // -1
	}()

	go func() { 
		write("Hello World")
		waitGroup.Done() // -1
	}()
	
	waitGroup.Wait() //  Espera a contagem das goroutines terminar
}

func write(txt string){
	for i := 0; i < 5; i++{
		fmt.Println(txt)
		time.Sleep(time.Second)
	}
}