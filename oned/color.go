package oned

var Black = Color{0, 0, 0}
var White = Color{1, 1, 1}

type Color struct {
	R, G, B float64
}

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
