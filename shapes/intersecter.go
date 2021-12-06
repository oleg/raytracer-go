package shapes

import (
	"github.com/oleg/raytracer-go/geom"
	"github.com/oleg/raytracer-go/physic"
)

type Intersecter interface {
	Intersect(ray Ray) Intersections //todo fix this
}

type Ray struct {
	Origin    geom.Point
	Direction geom.Vector
}

func (ray Ray) Position(distance float64) geom.Point {
	return ray.Origin.AddVector(ray.Direction.MultiplyScalar(distance))
}

func (ray Ray) ToLocal(shape physic.TransformationProvider) Ray {
	m := shape.Transformation().Inverse()
	return Ray{
		m.MultiplyPoint(ray.Origin),
		m.MultiplyVector(ray.Direction),
	}
}

type Intersections []Intersection //todo rename to Hits

type Intersection struct { //todo rename to Hit?
	Distance float64
	Object   Shape //todo try to remove shape from this struct
}
