package figure

import (
	"gray/multid"
	"gray/oned"
	"math"
)

type Plane struct {
	transform multid.Matrix4 //todo test
	material  Material       //todo test
}

func MakePlane() Plane {
	return Plane{multid.IdentityMatrix, DefaultMaterial()}
}
func MakePlaneTM(transform multid.Matrix4, material Material) Plane {
	return Plane{transform, material}
}

func (p Plane) LocalIntersect(ray Ray) Inters {
	if math.Abs(ray.Direction.Y) < oned.Delta {
		return nil //is it ok or Inters{}?
	}
	t := -ray.Origin.Y / ray.Direction.Y
	return Inters{Inter{t, p}}
}

func (p Plane) LocalNormalAt(point oned.Point) oned.Vector {
	return oned.Vector{0, 1, 0}
}

func (p Plane) Transform() multid.Matrix4 {
	return p.transform
}
func (p Plane) Material() Material {
	return p.material
}
