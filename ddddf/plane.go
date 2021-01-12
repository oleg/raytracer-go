package ddddf

import (
	"github.com/oleg/raytracer-go/asdf"
	"github.com/oleg/raytracer-go/geom"
	"math"
)

type Plane struct {
	ShapePhysics
}

func MakePlane() Plane {
	return Plane{ShapePhysics{geom.IdentityMatrix(), asdf.DefaultMaterial()}}
}
func NewPlane(transform *geom.Matrix, material *asdf.Material) Plane {
	return Plane{ShapePhysics{transform, material}}
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
