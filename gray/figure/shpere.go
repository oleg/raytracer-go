package figure

import (
	"gray/multid"
	"gray/oned"
	"math"
)

type Sphere struct {
	transform multid.Matrix4
	material  Material
}

func MakeSphere() Sphere {
	return Sphere{multid.IdentityMatrix, DefaultMaterial()}
}

func MakeSphereTM(transform multid.Matrix4, material Material) Sphere {
	return Sphere{transform, material}
}

func MakeSphereT(transform multid.Matrix4) Sphere {
	return Sphere{transform, DefaultMaterial()}
}

func MakeSphereM(material Material) Sphere {
	return Sphere{multid.IdentityMatrix, material}
}

func MakeGlassSphere() Sphere {
	return Sphere{
		multid.IdentityMatrix,
		GlassMaterialBuilder().Build()}
}

func (sphere Sphere) Transform() multid.Matrix4 {
	return sphere.transform
}
func (sphere Sphere) Material() Material {
	return sphere.material
}

//todo or Sphere?
func (sphere Sphere) LocalIntersect(ray Ray) Inters {
	sphereToRay := ray.Origin.SubtractPoint(oned.Point{})
	a := ray.Direction.Dot(ray.Direction)
	b := 2 * ray.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return Inters{}
	}

	dSqrt := math.Sqrt(discriminant)
	t1 := (-b - dSqrt) / (2 * a)
	t2 := (-b + dSqrt) / (2 * a)
	return Inters{
		Inter{t1, sphere},
		Inter{t2, sphere},
	}
}

func (sphere Sphere) LocalNormalAt(localPoint oned.Point) oned.Vector {
	return localPoint.SubtractPoint(oned.Point{})
}
