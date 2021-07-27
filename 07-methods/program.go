package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (product *Product) applyDiscount(discountPercent float32) {
	product.Cost = product.Cost * ((100 - discountPercent) / 100)
}

//using type alias
//type PrintableProduct product

//inorder to use the applyDiscount method, we should use type composition
type PrintableProduct struct {
	Product
}

func (p *PrintableProduct) Format() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func main() {
	//pen := Product{100, "Pen", 10, 50, "Stationary"}
	pen := Product{Id: 100, Name: "Pen", Cost: 10, Units: 50, Category: "Stationary"}
	fmt.Println(pen)
	fmt.Println("Cost of pen => ", pen.Cost)
	//applyDiscount(&pen, 10)
	//(&pen).applyDiscount(10)
	pen.applyDiscount(10)
	fmt.Println("After applying 10% discount, Pen => ", pen)

	//typecasting for an aliased type
	/*
		pp := PrintableProduct{400, "Paper", 1000, 1, "Stationary"}
		fmt.Println(pp.Format())
		printablePen := PrintableProduct(pen)
		fmt.Println(printablePen.Format())
	*/

	//using type composition
	pp := PrintableProduct{Product{400, "Paper", 1, 1000, "Stationary"}}
	fmt.Println(pp.Format())
	printablePen := PrintableProduct{pen}
	fmt.Println(printablePen.Format())

	pp.applyDiscount(10)
	fmt.Println(pp.Format())
}
