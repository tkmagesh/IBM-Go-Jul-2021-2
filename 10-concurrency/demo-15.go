package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	stop := make(chan bool)
	go fibonacci(ch, stop)
	go func() {
		for no := range ch {
			println(no)
		}
	}()
	var input string
	fmt.Scanln(&input)
	stop <- true
	fmt.Println("Done")
}

func fibonacci(ch chan int, stop chan bool) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			time.Sleep(500 * time.Millisecond)
			x, y = y, x+y
		case <-stop:
			fmt.Println("Stopped")
			close(ch)
			return
		}
	}

}
