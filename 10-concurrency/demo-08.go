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
	wg.Add(2)
	go add(10, 20, ch, 5)
	go add(100, 200, ch, 3)
	result := <-ch
	fmt.Println(result)
	result = <-ch
	fmt.Println(result)
	wg.Wait()
	fmt.Println("Done")
}

func add(x, y int, ch chan int, elapseTime time.Duration) {
	time.Sleep(elapseTime * time.Second)
	result := x + y
	ch <- result
	wg.Done()
}
