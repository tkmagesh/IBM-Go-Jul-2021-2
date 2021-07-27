package utils

type ShareWithArea interface {
	Area() float32
}

func PrintArea(sa ShareWithArea) {
	println("Area = ", sa.Area())
}

type ShareWithPerimeter interface {
	Perimeter() float32
}

func PrintPerimeter(sa ShareWithPerimeter) {
	println("Perimeter = ", sa.Perimeter())
}

/* type ShapeWithDimension interface {
	Area() float32
	Perimeter() float32
} */

type ShapeWithDimension interface {
	ShareWithArea
	ShareWithPerimeter
}

func PrintDimension(sd ShapeWithDimension) {
	println("Area = ", sd.Area())
	println("Perimeter = ", sd.Perimeter())
}
