package main

import (
	"fmt"
	"io"
	"os"
)

type AlphaReader string

func (alphaReader AlphaReader) Read(p []byte) (n int, err error) {
	fmt.Println("Buffer size ", len(p))
	count := 0
	for idx := 0; idx < len(alphaReader); idx++ {
		if (alphaReader[idx] >= 'A' && alphaReader[idx] <= 'Z') || (alphaReader[idx] >= 'a' && alphaReader[idx] <= 'z') {
			p[count] = alphaReader[idx]
			count++
		}
	}
	return count, io.EOF
}

func main() {
	alphaReader := AlphaReader("this is a sample string with 1234 numbers and *&^%^$ special characters")
	io.Copy(os.Stdout, alphaReader)
}
