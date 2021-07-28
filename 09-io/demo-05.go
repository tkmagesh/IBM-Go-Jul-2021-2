package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fileContents, err := ioutil.ReadFile("data1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(fileContents))
}
