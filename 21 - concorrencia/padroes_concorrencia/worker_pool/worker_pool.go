package main

import (
	"fmt"
	"time"
)


func fibo(position uint) uint {
	if position <= 1 {
		return position
	}
	return fibo(position - 2) + fibo(position - 1)
}

func worker(tasks <- chan uint, results chan <- uint){
	defer close(results)
	for num := range tasks{
		results <- fibo(num)
	}
}

func worker2(tasks chan <- uint){
	defer close(tasks)
	for i:=uint(0);i<=45;i++{
		tasks <- i
	}
}


func main(){
	start := time.Now()
	fmt.Println("Funções Recursivas")
	taskChannel := make(chan uint, 45)
	resultChannel := make(chan uint, 45)

	go worker(taskChannel, resultChannel)
	go worker(taskChannel, resultChannel)
	worker2(taskChannel)

	for result := range resultChannel{
		fmt.Println(result)
	}
	fmt.Println("Time Result: ", time.Since(start))
	// Com 1 worker: 9.8560422s
	// Com 2 worker: 3.7882274s
}
