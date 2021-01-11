package figure

import (
	"github.com/oleg/raytracer-go/multid"
	"github.com/oleg/raytracer-go/oned"
	"math"
)

type Plane struct {
	transform *multid.Matrix4 //todo test
	material  Material        //todo test
}

func MakePlane() Plane {
	return Plane{multid.IdentityMatrix(), DefaultMaterial()}
}
func MakePlaneTM(transform *multid.Matrix4, material Material) Plane {
	return Plane{transform, material}
}

func (p Plane) LocalIntersect(ray Ray) Inters {
	if math.Abs(ray.Direction.Y) < oned.Delta {
		return nil //is it ok or Inters{}?
	}
	t := -ray.Origin.Y / ray.Direction.Y
	return Inters{Inter{t, p}}
}

func (p Plane) LocalNormalAt(oned.Point) oned.Vector {
	return oned.Vector{X: 0, Y: 1, Z: 0}
}

func (p Plane) Transform() *multid.Matrix4 {
	return p.transform
}
func (p Plane) Material() Material {
	return p.material
}
