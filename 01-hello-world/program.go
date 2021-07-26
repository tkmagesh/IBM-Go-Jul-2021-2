//package declaration
package main

//import dependent packages
import "fmt"

//pakcage level variables & types
// var appName string = "myapp"
var appName = "myapp"

//package init function

//main function
func main() {
	/*
		var msg string
		msg = "Hello World"
	*/
	/*
		var msg string = "Hello World"
	*/
	/*
		var msg = "Hello World"
	*/
	//The following syntax is applicable only in a function
	msg := "Hello World"
	fmt.Println(msg)

	//multiple variable declarations
	/*
		var x int
		var y int
		var result int
		x = 10
		y = 20
		result = x + y
	*/
	/*
		var x, y, result int
		x = 10
		y = 20
		result = x + y
	*/
	/*
		var (
			x      int
			y      int
			result int
			s      string
		)
		x = 10
		y = 20
		s = "Dummy string"
		result = x + y
	*/
	/*
		var (
			x, y, result int
			s            string
		)
		x = 10
		y = 20
		s = "Dummy string"
		result = x + y
	*/
	/*
		var (
			x, y   int = 10, 20
			result int
			s      string = "Dummy string"
		)
	*/
	/*
		var (
			x, y   = 10, 20
			result int
			s      = "Dummy string"
		)
	*/
	/*
		x, y := 10, 20
		result := x + y
		s := "Dummy string"
	*/
	x, y, s := 10, 20, "Dummy string"
	result := x + y

	fmt.Println(result, s)

	//using package level variable
	//fmt.Println(appName)

	// var myVar string = "This is a sample string"
	//myVar := "This is a sample string"
}

//other functions
