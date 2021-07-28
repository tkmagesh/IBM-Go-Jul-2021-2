package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	wg.Add(2)
	go print("Hello", 3, ch1, ch2, wg)
	go print("World", 2, ch2, ch1, wg)
	ch1 <- 1
	wg.Wait()
	fmt.Println("Done")
}

func print(s string, delay time.Duration, in chan int, out chan int, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		<-in
		println(s)
		time.Sleep(delay * time.Second)
		out <- 1
	}
	wg.Done()
}

/*
expected output:

Hello
World
Hello
World
Hello
World
Hello
World
Hello
World
*/
