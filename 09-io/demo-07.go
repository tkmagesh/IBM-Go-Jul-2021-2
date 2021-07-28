package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fileReader, err := os.Open("data3.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer fileReader.Close()
	scanner := bufio.NewScanner(fileReader)
	lineCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineCount++
		fmt.Printf("line # %d: %s\n", lineCount, line)
	}

}
