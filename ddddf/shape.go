package ddddf

import (
	"github.com/oleg/raytracer-go/geom"
	"github.com/oleg/raytracer-go/mat"
)

type NormalFinder interface {
	NormalAt(point geom.Point) geom.Vector
}



type Shape interface {
	mat.HasTransformation
	mat.HasMaterial
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
	//todo move to World.Intersect?
	localRay := worldRay.ToLocal(shape)
	return shape.Intersect(localRay)
}
