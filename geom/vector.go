package geom

import "math"

//w = 0
type Vector struct {
	X, Y, Z float64
}

//todo: use link?
func (t Vector) AddVector(o Vector) Vector {
	return Vector{t.X + o.X, t.Y + o.Y, t.Z + o.Z}
}

func (t Vector) SubtractVector(o Vector) Vector {
	return Vector{t.X - o.X, t.Y - o.Y, t.Z - o.Z}
}

func (t Vector) Negate() Vector {
	return Vector{-t.X, -t.Y, -t.Z}
}

func (t Vector) Dot(o Vector) float64 {
	return t.X*o.X + t.Y*o.Y + t.Z*o.Z
}

func (t Vector) Cross(o Vector) Vector {
	return Vector{
		t.Y*o.Z - t.Z*o.Y,
		t.Z*o.X - t.X*o.Z,
		t.X*o.Y - t.Y*o.X,
	}
}

func (t Vector) MultiplyScalar(scalar float64) Vector {
	return Vector{t.X * scalar, t.Y * scalar, t.Z * scalar}
}

func (t Vector) Magnitude() float64 {
	return math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z)
}

func (t Vector) Normalize() Vector {
	magnitude := t.Magnitude()
	return Vector{t.X / magnitude, t.Y / magnitude, t.Z / magnitude}
}

func (t Vector) Reflect(normal Vector) Vector {
	dot := t.Dot(normal)
	temp := normal.MultiplyScalar(2).MultiplyScalar(dot)
	return t.SubtractVector(temp)
}
