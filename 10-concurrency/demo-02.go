package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = &sync.WaitGroup{}

func main() {

	wg.Add(2)
	go print("Hello")
	go print("World")
	fmt.Println("Done")
	wg.Wait()
}

func print(s string) {
	time.Sleep(3 * time.Second)
	fmt.Println(s)
	wg.Done()
}
