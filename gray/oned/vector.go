package oned

import "math"

type Vector Tuple //w = 0
//todo: use link
func (t Vector) AddVector(o Vector) Vector {
	return Vector(Tuple(t).add(Tuple(o)))
}

func (t Vector) SubtractVector(o Vector) Vector {
	return Vector(Tuple(t).subtract(Tuple(o)))
}

func (t Vector) Negate() Vector {
	return Vector(Tuple(t).negate())
}

func (t Vector) Dot(o Vector) float64 {
	return Tuple(t).dot(Tuple(o))
}

func (t Vector) Cross(o Vector) Vector {
	return Vector(Tuple(t).cross(Tuple(o)))
}

func (t Vector) MultiplyScalar(scalar float64) Vector {
	return Vector(Tuple(t).multiplyScalar(scalar))
}

func (t Vector) Magnitude() float64 { //todo: move to Tuple?
	return math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z)
}

func (t Vector) Normalize() Vector { //todo: move to Tuple?
	magnitude := t.Magnitude()
	return Vector{t.X / magnitude, t.Y / magnitude, t.Z / magnitude}
}

func (t Vector) Reflect(normal Vector) Vector {
	dot := t.Dot(normal)
	temp := normal.MultiplyScalar(2).MultiplyScalar(dot)
	return t.SubtractVector(temp)
}
