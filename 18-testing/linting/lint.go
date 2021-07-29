package linting

import "fmt"

func checkFlag(flag bool) bool {
	return flag
}

func errChecking() (int, error) {
	a := 2
	fmt.Println(a)
	return 0, nil
}

func callsErrChecking() int {
	val, _ := errChecking()
	return val
}
