package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func main() {
	//pen := Product{100, "Pen", 10, 50, "Stationary"}
	pen := Product{Id: 100, Name: "Pen", Cost: 10, Units: 50, Category: "Stationary"}
	fmt.Println(pen)
	fmt.Println("Cost of pen => ", pen.Cost)
	applyDiscount(&pen, 10)
	fmt.Println("After applying 10% discount, Pen => ", pen)
}

func applyDiscount(product *Product, discountPercent float32) {
	product.Cost = product.Cost * ((100 - discountPercent) / 100)
}
