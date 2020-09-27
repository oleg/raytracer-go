package figure

import (
	"gray/multid"
	"gray/oned"
)

type Shape interface {
	Transform() multid.Matrix4 //is this method needed in the interface
	LocalIntersect(ray Ray) Inters
	LocalNormalAt(point oned.Point) oned.Vector
	Material() Material
}

func NormalAt(shape Shape, worldPoint oned.Point) oned.Vector {
	localPoint := shape.Transform().Inverse().MultiplyPoint(worldPoint)
	localNormal := shape.LocalNormalAt(localPoint)
	worldNormal := shape.Transform().Inverse().Transpose().MultiplyVector(localNormal)
	return worldNormal.Normalize()
}

func Intersect(shape Shape, worldRay Ray) Inters {
	localRay := worldRay.Transform(shape.Transform().Inverse())
	return shape.LocalIntersect(localRay)
}
