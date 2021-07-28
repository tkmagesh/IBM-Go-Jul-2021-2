package main

import (
	"fmt"
	"sync"
	"time"
)

//share memory by communicating (adviced)

var wg = &sync.WaitGroup{}

func main() {
	//deadlock simulation
	//ch := make(chan int)

	//using buffered channel
	ch := make(chan int, 1)
	wg.Add(1)
	go add(10, 20, ch)
	wg.Wait()
	result := <-ch
	fmt.Println(result)
	fmt.Println("Done")
}

func add(x, y int, ch chan int) {
	time.Sleep(5 * time.Second)
	result := x + y
	ch <- result
	wg.Done()
}
