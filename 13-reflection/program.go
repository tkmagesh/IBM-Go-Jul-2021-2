package main

import (
	"fmt"
	"reflect"
)

type Employee struct {
	Id     int
	Name   string
	Salary float32
}

func main() {
	emp := Employee{100, "Magesh", 10000}
	values := reflect.ValueOf(emp)
	typeObj := values.Type()
	for i := 0; i < values.NumField(); i++ {
		fmt.Println(typeObj.Field(i).Name, values.FieldByName(typeObj.Field(i).Name))
	}
}
