package main

import (
	"fmt"
	"time"
)

func main() {
	go print("Hello")
	go print("World")
	fmt.Println("Done")
	time.Sleep(2 * time.Second)
	//runtime.Gosched()
}

func print(s string) {
	fmt.Println(s)
}
