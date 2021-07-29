package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(sum(10, 20))                                                        // => 30
	fmt.Println(sum(10, 20, 30, 40))                                                // => 100
	fmt.Println(sum(10, 20, 30, "40"))                                              // => 100
	fmt.Println(sum(10, 20, 30, "40", "abc"))                                       // => 100
	fmt.Println(sum(10, 20, []int{30, 40}))                                         // => 100
	fmt.Println(sum(10, 20, []interface{}{30, 40, []int{50, 60}}))                  // => 210
	fmt.Println(sum(10, 20, []interface{}{30, 40, []interface{}{50, "60", "abc"}})) // => 210
}

func sum(nos ...interface{}) int {
	result := 0
	for _, no := range nos {
		switch value := no.(type) {
		case int:
			result += value
		case string:
			if no, err := strconv.Atoi(value); err == nil {
				result += no
			}
		case []int:
			intValList := make([]interface{}, len(value))
			for i, v := range value {
				intValList[i] = v
			}
			result += sum(intValList...)
		case []interface{}:
			result += sum(value...)
		default:
			continue
		}
	}
	return result
}
