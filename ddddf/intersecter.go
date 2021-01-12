package ddddf

import (
	"github.com/oleg/raytracer-go/geom"
	"github.com/oleg/raytracer-go/mat"
)

type Intersecter interface {
	Intersect(ray Ray) Inters //todo fix this
}

type Ray struct {
	Origin    geom.Point
	Direction geom.Vector
}

func (ray Ray) Position(distance float64) geom.Point {
	return ray.Origin.AddVector(ray.Direction.MultiplyScalar(distance))
}

func (ray Ray) ToLocal(shape mat.HasTransformation) Ray {
	m := shape.Transformation().Inverse()
	return Ray{
		m.MultiplyPoint(ray.Origin),
		m.MultiplyVector(ray.Direction),
	}
}

type Inters []Inter //todo rename

type Inter struct { //todo rename
	Distance float64
	Object   Shape //todo try to remove shape from this struct
}
