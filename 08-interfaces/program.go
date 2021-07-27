package main

import (
	"fmt"
	"interfaces-app/models"
	"interfaces-app/shapes"
	"interfaces-app/utils"
)

func main() {
	c := shapes.Circle{Radius: 3}
	fmt.Println(c)
	/* utils.PrintArea(c)
	utils.PrintPerimeter(c) */
	utils.PrintDimension(c)

	r := shapes.Rectangle{Width: 10, Height: 12}
	fmt.Println(r)
	/* utils.PrintArea(r)
	utils.PrintPerimeter(r) */
	utils.PrintDimension(r)

	products := models.Products{
		models.Product{105, "Pen", 5, 50, "Stationary"},
		models.Product{103, "Pencil", 2, 100, "Stationary"},
		models.Product{102, "Marker", 50, 20, "Stationary"},
		models.Product{104, "Frying Pan", 500, 5, "Utencil"},
		models.Product{101, "Phone", 5000, 3, "Electronics"},
		models.Product{100, "Bowl", 100, 50, "Utencil"},
	}
	/*
		using the sort apis implemented to sort the products by id, name, cost etc and print the result
	*/
}
