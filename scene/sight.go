package scene

import "github.com/oleg/raytracer-go/geom"

type Sight struct {
	From, To geom.Point
	Up       geom.Vector
}

func (s Sight) Transformation() *geom.Matrix {
	forward := s.To.SubtractPoint(s.From).Normalize()
	left := forward.Cross(s.Up.Normalize())
	trueUp := left.Cross(forward)

	orientation := geom.Matrix{
		Data: [4][4]float64{
			{left.X, left.Y, left.Z, 0},
			{trueUp.X, trueUp.Y, trueUp.Z, 0},
			{-forward.X, -forward.Y, -forward.Z, 0},
			{0, 0, 0, 1},
		},
	}
	return orientation.Multiply(geom.Translation(-s.From.X, -s.From.Y, -s.From.Z))
}
