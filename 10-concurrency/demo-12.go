package main

func main() {
	/* split the data into two sets, add them concurrently, collate the final result and display the final result */
	data := []int{4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55}

	firstSet := data[0 : len(data)/2]
	secondSet := data[len(data)/2:]
	resultCh1 := make(chan int)
	resultCh2 := make(chan int)
	go sum(resultCh1, firstSet...)
	go sum(resultCh2, secondSet...)
	finalResult := <-resultCh1 + <-resultCh2
	println(finalResult)
}

func sum(resultCh chan int, nos ...int) {
	result := 0
	for _, no := range nos {
		result += no
	}
	resultCh <- result
}
