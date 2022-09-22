package figuras

import "math"

type Circle struct {
	Radio float64
}

func (circle *Circle) area() float64 {
	return math.Pi * (circle.Radio * circle.Radio)
}
func (circle *Circle) perimeter() float64 {
	return 3.1416 * (circle.Radio * 2)
}
