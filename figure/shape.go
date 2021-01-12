package figure

import (
	"github.com/oleg/raytracer-go/geom"
)

type HasTransformation interface { //todo try to convert to an action
	Transformation() *geom.Matrix
}

type Intersecter interface {
	Intersect(ray Ray) Inters //todo fix this
}

type NormalFinder interface {
	NormalAt(point geom.Point) geom.Vector
}

type ShapePhysics struct {
	transform *geom.Matrix //todo test
	material  *Material    //todo test
}

func (p ShapePhysics) Transformation() *geom.Matrix {
	return p.transform
}
func (p ShapePhysics) Material() *Material {
	return p.material
}

type Shape interface {
	HasTransformation
	HasMaterial
	Intersecter
	NormalFinder
}

func NormalAt(shape Shape, worldPoint geom.Point) geom.Vector {
	localPoint := shape.Transformation().Inverse().MultiplyPoint(worldPoint)
	localNormal := shape.NormalAt(localPoint)
	worldNormal := shape.Transformation().Inverse().Transpose().MultiplyVector(localNormal)
	return worldNormal.Normalize()
}

func Intersect(shape Shape, worldRay Ray) Inters {
	m := shape.Transformation().Inverse()
	localRay := Ray{
		m.MultiplyPoint(worldRay.Origin),
		m.MultiplyVector(worldRay.Direction),
	}
	return shape.Intersect(localRay)
}
