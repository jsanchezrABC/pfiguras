package figuras

import "fmt"

type Calculate interface {
	area() float64
	perimeter() float64
}

func CalculateArea(calculate Calculate) {
	fmt.Println(calculate.area())
}
func CalculatePerimeter(calculate Calculate) {
	fmt.Println(calculate.perimeter())
}
