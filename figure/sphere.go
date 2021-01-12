package figure

import (
	"github.com/oleg/raytracer-go/geom"
	"math"
)

type Sphere struct {
	transform *geom.Matrix
	material  Material
}

func NewSphere(transform *geom.Matrix, material Material) Sphere {
	return Sphere{transform, material}
}

func NewSphereT(transform *geom.Matrix) Sphere {
	return Sphere{transform, DefaultMaterial()}
}

func NewGlassSphere() Sphere {
	return Sphere{
		geom.IdentityMatrix(),
		GlassMaterialBuilder().Build()}
}

func (sphere Sphere) Transform() *geom.Matrix {
	return sphere.transform
}
func (sphere Sphere) Material() Material {
	return sphere.material
}

//todo or Sphere?
func (sphere Sphere) LocalIntersect(ray Ray) Inters {
	sphereToRay := ray.Origin.SubtractPoint(geom.Point{})
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

func (sphere Sphere) LocalNormalAt(localPoint geom.Point) geom.Vector {
	return localPoint.SubtractPoint(geom.Point{})
}
