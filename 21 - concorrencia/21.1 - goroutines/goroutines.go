package main

import (
	"fmt"
	"time"
)

func main(){
	go write("Bom dia") // goroutine
	write("Hello World")
}

func write(txt string){
	for {
		fmt.Println(txt)
		time.Sleep(time.Second)
	}
}