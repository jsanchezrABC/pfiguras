package figuras

type Squares struct {
	Width float64
}

func (squares *Squares) area() float64 {
	return squares.Width * squares.Width
}
func (squares *Squares) perimeter() float64 {
	return (2 * squares.Width) + (2 * squares.Width)
}
