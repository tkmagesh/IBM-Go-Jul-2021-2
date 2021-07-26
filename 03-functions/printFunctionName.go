package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	// "seps"
)

func foo() {
}

func GetFunctionName(i interface{}, seps ...rune) string {
	// Get the function name
	fn := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()

	// split with seps
	fields := strings.FieldsFunc(fn, func(sep rune) bool {
		for _, s := range seps {
			if sep == s {
				return true
			}
		}
		return false
	})

	// fmt.Println(fields)

	if size := len(fields); size > 0 {
		return fields[size-1]
	}
	return ""
}

func main() {
	// This will print "name: main.foo"
	fmt.Println("name:", GetFunctionName(foo))
}
