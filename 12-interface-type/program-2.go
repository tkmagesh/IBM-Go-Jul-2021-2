package main

import "fmt"

func main() {
	fmt.Println(sum(10, 20))                                                        // => 30
	fmt.Println(sum(10, 20, 30, 40))                                                // => 100
	fmt.Println(sum(10, 20, 30, "40"))                                              // => 100
	fmt.Println(sum(10, 20, 30, "40", "abc"))                                       // => 100
	fmt.Println(sum(10, 20, []int{30, 40}))                                         // => 100
	fmt.Println(sum(10, 20, []interface{}{30, 40, []int{50, 60}}))                  // => 210
	fmt.Println(sum(10, 20, []interface{}{30, 40, []interface{}{50, "60", "abc"}})) // => 210
}

func sum() int {

}
