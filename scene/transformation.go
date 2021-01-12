package scene

import (
	"github.com/oleg/raytracer-go/geom"
)

//todo move? rename file?
func ViewTransform(from, to geom.Point, up geom.Vector) *geom.Matrix {
	forward := to.SubtractPoint(from).Normalize()
	left := forward.Cross(up.Normalize())
	trueUp := left.Cross(forward)

	orientation := geom.Matrix{
		Data: [4][4]float64{
			{left.X, left.Y, left.Z, 0},
			{trueUp.X, trueUp.Y, trueUp.Z, 0},
			{-forward.X, -forward.Y, -forward.Z, 0},
			{0, 0, 0, 1},
		},
	}
	translation := geom.Translation(-from.X, -from.Y, -from.Z)
	return orientation.Multiply(translation)
}
