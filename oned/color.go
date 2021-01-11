package oned

type Color struct {
	R, G, B float64
}

var Black = Color{0, 0, 0}
var White = Color{1, 1, 1}

//func NewColor(R, G, B float64) Color {
//	return Color{R, G, B}
//}

func (t Color) Add(o Color) Color {
	return Color{t.R + o.R, t.G + o.G, t.B + o.B}
}

func (t Color) Subtract(o Color) Color {
	return Color{t.R - o.R, t.G - o.G, t.B - o.B}
}

func (t Color) MultiplyByScalar(scalar float64) Color {
	return Color{t.R * scalar, t.G * scalar, t.B * scalar}
}

func (t Color) Multiply(o Color) Color {
	return Color{t.R * o.R, t.G * o.G, t.B * o.B}
}
