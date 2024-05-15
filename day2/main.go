package main

import (
	"fmt"
	"time"
)

func counter(x int) {
	for i := 0; i < x; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func mainOld1() { // T1
	go counter(10) // T2
	go counter(10) // T3
	counter(10) // T1
}

func mainOld2() { // T1
	channel := make(chan string)

	go func() { // T2
		channel <- "opa"
	}()

	msg := <- channel
	fmt.Println(msg) // T1
}

func worker(workerID int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}

func mainOld3() { // T1
	channel := make(chan int)

	workersQuantity := 100

	for j := 0; j < workersQuantity; j++ {
		go worker(j, channel)
	}

	for i := 0; i < 1000; i++ {
		channel <- i
	}
}

func worker2(workerID int, data chan int, out chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerID, x)
		time.Sleep(time.Second)
		out <- x
	}
}

func main() { // T1
	channel := make(chan int)

	workersQuantity := 1000000

	for j := 0; j < workersQuantity; j++ {
		go worker(j, channel)
	}

	for i := 0; i < 10000000; i++ {
		channel <- i
	}
}
