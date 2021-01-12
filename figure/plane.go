package figure

import (
	"github.com/oleg/raytracer-go/geom"
	"math"
)

type Plane struct {
	transform *geom.Matrix //todo test
	material  *Material    //todo test
}

func MakePlane() Plane {
	return Plane{geom.IdentityMatrix(), DefaultMaterial()}
}
func MakePlaneTM(transform *geom.Matrix, material *Material) Plane {
	return Plane{transform, material}
}

func (p Plane) LocalIntersect(ray Ray) Inters {
	if math.Abs(ray.Direction.Y) < geom.Delta {
		return nil //is it ok or Inters{}?
	}
	t := -ray.Origin.Y / ray.Direction.Y
	return Inters{Inter{t, p}}
}

func (p Plane) LocalNormalAt(geom.Point) geom.Vector {
	return geom.Vector{X: 0, Y: 1, Z: 0}
}

func (p Plane) Transform() *geom.Matrix {
	return p.transform
}
func (p Plane) Material() *Material {
	return p.material
}
