package shapes

import (
	"github.com/oleg/raytracer-go/geom"
	"github.com/oleg/raytracer-go/physic"
	"math"
)

type Sphere struct {
	physic.PhysicalObject
}

func NewSphere(transform *geom.Matrix, material *physic.Material) Sphere {
	return Sphere{physic.NewPhysicalObject(transform, material)}
}

func NewGlassSphere() Sphere {
	return Sphere{physic.NewPhysicalObject(
		geom.IdentityMatrix(),
		physic.GlassMaterialBuilder().Build(),
	)}
}

//todo or Sphere?
func (sphere Sphere) Intersect(ray Ray) Intersections {
	sphereToRay := ray.Origin.SubtractPoint(geom.Point{})
	a := ray.Direction.Dot(ray.Direction)
	b := 2 * ray.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return Intersections{}
	}

	dSqrt := math.Sqrt(discriminant)
	t1 := (-b - dSqrt) / (2 * a)
	t2 := (-b + dSqrt) / (2 * a)
	return Intersections{
		Intersection{t1, sphere},
		Intersection{t2, sphere},
	}
}

func (sphere Sphere) NormalAt(localPoint geom.Point) geom.Vector {
	return localPoint.SubtractPoint(geom.Point{})
}
