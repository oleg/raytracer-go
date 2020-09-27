package oned

const Delta = 0.000009

type Tuple struct {
	X, Y, Z float64
}

//todo: use link
func (t Tuple) add(o Tuple) Tuple {
	return Tuple{t.X + o.X, t.Y + o.Y, t.Z + o.Z}
}

func (t Tuple) subtract(o Tuple) Tuple {
	return Tuple{t.X - o.X, t.Y - o.Y, t.Z - o.Z}
}

func (t Tuple) negate() Tuple {
	return Tuple{-t.X, -t.Y, -t.Z}
}

func (t Tuple) multiplyScalar(scalar float64) Tuple {
	return Tuple{t.X * scalar, t.Y * scalar, t.Z * scalar}
}

func (t Tuple) divideScalar(scalar float64) Tuple {
	return Tuple{t.X / scalar, t.Y / scalar, t.Z / scalar}
}

func (t Tuple) hadamard(o Tuple) Tuple {
	return Tuple{t.X * o.X, t.Y * o.Y, t.Z * o.Z}
}

func (t Tuple) cross(o Tuple) Tuple {
	return Tuple{
		t.Y*o.Z - t.Z*o.Y,
		t.Z*o.X - t.X*o.Z,
		t.X*o.Y - t.Y*o.X}
}

func (t Tuple) dot(o Tuple) float64 {
	return t.X*o.X + t.Y*o.Y + t.Z*o.Z
}
