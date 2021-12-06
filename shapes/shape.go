package shapes

import (
	"github.com/oleg/raytracer-go/geom"
	"github.com/oleg/raytracer-go/physic"
)

type NormalFinder interface {
	NormalAt(point geom.Point) geom.Vector
}

type Shape interface {
	physic.TransformationProvider
	physic.MaterialProvider
	Intersecter
	NormalFinder
}
