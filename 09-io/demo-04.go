package main

import (
	"fmt"
	"io"
	"os"
)

type AlphaReader struct {
	Src io.Reader
}

func (alphaReader AlphaReader) Read(buffer []byte) (int, error) {
	inputData := make([]byte, len(buffer))
	count, err := alphaReader.Src.Read(inputData)
	if err != nil {
		return count, err
	}
	dataCount := 0
	for i := 0; i < len(inputData); i++ {
		if (inputData[i] >= 'A' && inputData[i] <= 'Z') || (inputData[i] >= 'a' && inputData[i] <= 'z') {
			buffer[dataCount] = inputData[i]
			dataCount++
		}
	}
	return dataCount, io.EOF
}

func main() {
	//stringReader := strings.NewReader("this is a sample string with 1234 numbers and *&^%^$ special characters")
	fileReader, err := os.Open("data1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	alphaReader := AlphaReader{Src: fileReader}
	io.Copy(os.Stdout, alphaReader)
}
