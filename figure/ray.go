package figure

import (
	"github.com/oleg/raytracer-go/multid"
	"github.com/oleg/raytracer-go/oned"
)

type Ray struct {
	Origin    oned.Point
	Direction oned.Vector
}

func (ray Ray) Position(distance float64) oned.Point {
	return ray.Origin.AddVector(ray.Direction.MultiplyScalar(distance))
}

func (ray Ray) Transform(m *multid.Matrix4) Ray {
	return Ray{
		m.MultiplyPoint(ray.Origin),
		m.MultiplyVector(ray.Direction),
	}
}
