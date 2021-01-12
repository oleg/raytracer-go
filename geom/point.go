package geom

//w = 1
type Point struct {
	X, Y, Z float64
}

//todo: use link
func (t Point) AddVector(o Vector) Point {
	return Point{t.X + o.X, t.Y + o.Y, t.Z + o.Z}
}

func (t Point) SubtractVector(o Vector) Point {
	return Point{t.X - o.X, t.Y - o.Y, t.Z - o.Z}
}

func (t Point) SubtractPoint(o Point) Vector {
	return Vector{t.X - o.X, t.Y - o.Y, t.Z - o.Z}
}
