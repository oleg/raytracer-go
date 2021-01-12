package figure

import (
	"github.com/oleg/raytracer-go/geom"
)

type Shape interface {
	Transform() *geom.Matrix //is this method needed in the interface
	LocalIntersect(ray Ray) Inters
	LocalNormalAt(point geom.Point) geom.Vector
	Material() Material
}

func NormalAt(shape Shape, worldPoint geom.Point) geom.Vector {
	localPoint := shape.Transform().Inverse().MultiplyPoint(worldPoint)
	localNormal := shape.LocalNormalAt(localPoint)
	worldNormal := shape.Transform().Inverse().Transpose().MultiplyVector(localNormal)
	return worldNormal.Normalize()
}

func Intersect(shape Shape, worldRay Ray) Inters {
	localRay := worldRay.Transform(shape.Transform().Inverse())
	return shape.LocalIntersect(localRay)
}
