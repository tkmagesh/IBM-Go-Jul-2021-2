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
	products := []Product{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{103, "Pencil", 2, 100, "Stationary"},
		Product{102, "Marker", 50, 20, "Stationary"},
		Product{104, "Frying Pan", 500, 5, "Utencil"},
		Product{101, "Phone", 5000, 3, "Electronics"},
		Product{100, "Bowl", 100, 50, "Utencil"},
	}
	fmt.Println(products)

	anyStationaryProducts := any(&products, func(product Product) bool {
		return product.Category == "Stationary"
	})
	fmt.Println(anyStationaryProducts)
}

func indexOf(products *[]Product, product Product) int {
	for index, currentProduct := range *products {
		if currentProduct == product {
			return index
		}
	}
	return -1
}

func includes(products *[]Product, product Product) bool {
	return indexOf(products, product) != -1
}

func any(products *[]Product, prodicate func(product Product) bool) bool {
	for _, product := range *products {
		if prodicate(product) {
			return true
		}
	}
	return false
}

func all(products *[]Product, prodicate func(product Product) bool) bool {
	for _, product := range *products {
		if !prodicate(product) {
			return false
		}
	}
	return true
}

func filter(products *[]Product, prodicate func(product Product) bool) *[]Product {
	result := []Product{}
	for _, product := range *products {
		if prodicate(product) {
			result = append(result, product)
		}
	}
	return &result
}

/*
write the following utility functions to manipulate the products
	IndexOf => returns the index of the given product
	Includes => return true/false based on the presence of the given product in the products list
	Any => return true /false based on the presence of atleast one product in the list matching the given criteria
		use case:
			Is there any cost product? (cost > 500)
			Is there any stationary product? (category == "Stationary")
	All => return true /false based on the condition that all the products in the list satisfy the given criteria
		use case:
			Are all products cost products? (cost > 500)
			Are all products stationary products? (category == "stationary")
	Filter => filter the products based on the given criteria
		use case:
			filter utencils from the products list
			filter under stocked products from the products list (units < 10)
*/
