package geom

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

func (t Color) RGBA() (r, g, b, a uint32) {
	r = uint32(clamp(t.R)*255) << 8
	g = uint32(clamp(t.G)*255) << 8
	b = uint32(clamp(t.B)*255) << 8
	a = uint32(255) << 8
	return
}

func clamp(r float64) float64 {
	if r < 0 {
		return 0
	}
	if r > 1 {
		return 1
	}
	return r
}
