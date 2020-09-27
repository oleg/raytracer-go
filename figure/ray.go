package figure

import (
	"github.com/oleg/graytracer/multid"
	"github.com/oleg/graytracer/oned"
)

type Ray struct {
	Origin    oned.Point
	Direction oned.Vector
}

func (ray Ray) Position(distance float64) oned.Point {
	return ray.Origin.AddVector(ray.Direction.MultiplyScalar(distance))
}

func (ray Ray) Transform(m multid.Matrix) Ray {
	return Ray{
		m.MultiplyPoint(ray.Origin),
		m.MultiplyVector(ray.Direction),
	}
}
