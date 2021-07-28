package main

import "sync"

//communicate by sharing memory (not adviced)
var result int

var mutex = &sync.Mutex{}
var wg = &sync.WaitGroup{}

func main() {
	wg.Add(2)
	go add(10, 20)
	go add(20, 30)
	wg.Wait()
}

func add(x, y int) {
	mutex.Lock()
	{
		result = x + y
	}
	mutex.Unlock()
	wg.Done()
}
