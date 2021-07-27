package shapes

import "fmt"

type Rectangle struct {
	Width  float32
	Height float32
}

func (r Rectangle) Area() float32 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle, Height = %v, Width = %v", r.Height, r.Width)
}
