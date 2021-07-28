package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lineCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineCount++
		fmt.Printf("line # %d: %s\n", lineCount, line)
	}

}
