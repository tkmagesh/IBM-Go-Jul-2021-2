package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

type Products []Product

func (products Products) indexOf(product Product) int {
	for index, currentProduct := range products {
		if currentProduct == product {
			return index
		}
	}
	return -1
}

func (products Products) includes(product Product) bool {
	return products.indexOf(product) != -1
}

func (products Products) any(predicate func(product Product) bool) bool {
	for _, product := range products {
		if predicate(product) {
			return true
		}
	}
	return false
}

func (products Products) all(predicate func(product Product) bool) bool {
	for _, product := range products {
		if !predicate(product) {
			return false
		}
	}
	return true
}

func (products Products) filter(predicate func(product Product) bool) *Products {
	result := Products{}
	for _, product := range products {
		if predicate(product) {
			result = append(result, product)
		}
	}
	return &result
}

func main() {
	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{103, "Pencil", 2, 100, "Stationary"},
		Product{102, "Marker", 50, 20, "Stationary"},
		Product{104, "Frying Pan", 500, 5, "Utencil"},
		Product{101, "Phone", 5000, 3, "Electronics"},
		Product{100, "Bowl", 100, 50, "Utencil"},
	}
	fmt.Println(products)

	anyStationaryProducts := products.any(func(product Product) bool {
		return product.Category == "Stationary"
	})
	fmt.Println(anyStationaryProducts)
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
