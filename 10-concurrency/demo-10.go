package main

import (
	"fmt"
	"time"
)

func main() {
	go print("Hello", 3)
	go print("World", 2)
	var input string
	fmt.Scanln(&input)
}

func print(s string, delay time.Duration) {
	for i := 0; i < 5; i++ {
		println(s)
		time.Sleep(delay * time.Second)
	}
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
