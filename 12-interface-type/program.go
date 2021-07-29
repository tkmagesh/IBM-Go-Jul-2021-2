package main

import "fmt"

type Employee struct {
	Name string
}

func main() {
	var x interface{}
	x = 20
	x = "Dummy string"
	x = true
	fmt.Println(x)

	//x =
	if no, ok := x.(int); ok {
		result := no + 100
		fmt.Println(result)
	} else {
		fmt.Println("Not an int")
	}

	//x = Employee{"John"}
	//x = 200
	//x = "Dummy string"
	x = true
	switch val := x.(type) {
	case int:
		fmt.Println("Double of x is ", val*2)
	case bool:
		fmt.Println("x is boolean type with value = ", val)
	case string:
		fmt.Println("Len of x is ", len(val))
	case Employee:
		fmt.Println("x is an Employee with name ", val.Name)
	default:
		fmt.Println("x is of unknown type")
	}

}
