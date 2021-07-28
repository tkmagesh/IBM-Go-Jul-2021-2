package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg = &sync.WaitGroup{}
	wg.Add(1)
	go print("Hello", wg)

	wg.Add(1)
	go print("World", wg)

	fmt.Println("Done")
	wg.Wait()
}

func print(s string, wg *sync.WaitGroup) {
	time.Sleep(3 * time.Second)
	fmt.Println(s)
	wg.Done()
}
