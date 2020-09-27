package oned

type Color Tuple

var Black = Color{0, 0, 0}
var White = Color{1, 1, 1}

func NewColor(R, G, B float64) Color {
	return Color{R, G, B}
}

//todo: use link
func (t Color) R() float64 {
	return t.X
}

func (t Color) G() float64 {
	return t.Y
}

func (t Color) B() float64 {
	return t.Z
}

func (t Color) Add(o Color) Color {
	return Color(Tuple(t).add(Tuple(o)))
}

func (t Color) Subtract(o Color) Color {
	return Color(Tuple(t).subtract(Tuple(o)))
}

func (t Color) MultiplyByScalar(scalar float64) Color {
	return Color(Tuple(t).multiplyScalar(scalar))
}

func (t Color) Multiply(o Color) Color {
	return Color(Tuple(t).hadamard(Tuple(o)))
}
