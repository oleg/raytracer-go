package shapes

import (
	"github.com/oleg/raytracer-go/geom"
	"github.com/oleg/raytracer-go/physic"
	"math"
)

type Plane struct {
	physic.PhysicalObject
}

func NewPlane(transform *geom.Matrix, material *physic.Material) Plane {
	return Plane{physic.NewPhysicalObject(transform, material)}
}

func (p Plane) Intersect(ray Ray) Inters {
	if math.Abs(ray.Direction.Y) < geom.Delta {
		return nil //is it ok or Inters{}?
	}
	t := -ray.Origin.Y / ray.Direction.Y
	return Inters{Inter{t, p}}
}

func (p Plane) NormalAt(geom.Point) geom.Vector {
	return geom.Vector{X: 0, Y: 1, Z: 0}
}
