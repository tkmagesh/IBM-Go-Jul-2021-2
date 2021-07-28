package main

import (
	"fmt"
	"sync"
	"time"
)

//share memory by communicating (adviced)

var wg = &sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(1)
	go add(10, 20, ch)
	result := <-ch
	fmt.Println(result)
	wg.Wait()
	fmt.Println("Done")
}

func add(x, y int, ch chan int) {
	time.Sleep(5 * time.Second)
	result := x + y
	ch <- result
	wg.Done()
}
