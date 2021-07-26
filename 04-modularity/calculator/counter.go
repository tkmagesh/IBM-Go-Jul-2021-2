package calculator

var count = 0

func increment() {
	count++
}

func GetInvocationCount() int {
	return count
}
