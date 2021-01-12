package ddddf

import (
	"github.com/oleg/raytracer-go/geom"
)

type Ray struct {
	Origin    geom.Point
	Direction geom.Vector
}

func (ray Ray) Position(distance float64) geom.Point {
	return ray.Origin.AddVector(ray.Direction.MultiplyScalar(distance))
}

//func (ray Ray) Transform(m *geom.Matrix) Ray {
//	return Ray{
//		m.MultiplyPoint(ray.Origin),
//		m.MultiplyVector(ray.Direction),
//	}
//}
