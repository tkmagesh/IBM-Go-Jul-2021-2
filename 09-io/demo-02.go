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
			if count >= len(p) {
				return count, nil
			}
		}
	}
	return count, io.EOF
}

func copy(writer io.Writer, reader io.Reader) {
	totalBytesRead := 0
	for {
		buffer := make([]byte, 10)
		bytesRead, err := reader.Read(buffer)

		totalBytesRead += bytesRead
		fmt.Println("Total bytes read ", totalBytesRead)
		if bytesRead > 0 {
			writer.Write(buffer[:bytesRead])
		}
		if err == io.EOF {
			break
		}
	}
}

func main() {
	alphaReader := AlphaReader("this is a sample string with 1234 numbers and *&^%^$ special characters")
	copy(os.Stdout, alphaReader)
}
