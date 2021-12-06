package physic

import (
	"github.com/oleg/raytracer-go/geom"
	"math"
)

type PatternFinder interface {
	PatternAt(point geom.Point) geom.Color
}

type Pattern interface {
	PatternFinder
	TransformationProvider
}

func PatternAtShape(pattern Pattern, shape TransformationProvider, worldPoint geom.Point) geom.Color {
	objectPoint := shape.Transformation().Inverse().MultiplyPoint(worldPoint)
	patternPoint := pattern.Transformation().Inverse().MultiplyPoint(objectPoint)
	return pattern.PatternAt(patternPoint)
}

type StripePattern struct {
	A, B geom.Color
	Transformable
}

func (p StripePattern) PatternAt(point geom.Point) geom.Color {
	if math.Mod(math.Floor(point.X), 2) == 0 {
		return p.A
	}
	return p.B
}

type GradientPattern struct {
	A, B geom.Color
	Transformable
}

func (p GradientPattern) PatternAt(point geom.Point) geom.Color {
	distance := p.B.Subtract(p.A)
	fraction := point.X - math.Floor(point.X)
	return p.A.Add(distance.MultiplyByScalar(fraction))
}

type RingPattern struct {
	A, B geom.Color
	Transformable
}

func (p RingPattern) PatternAt(point geom.Point) geom.Color {
	hypot := math.Floor(math.Hypot(point.X, point.Z))
	if math.Mod(hypot, 2) == 0 {
		return p.A
	}
	return p.B
}

type CheckersPattern struct {
	A, B geom.Color
	Transformable
}

func (p CheckersPattern) PatternAt(point geom.Point) geom.Color {
	sum := math.Floor(point.X) + math.Floor(point.Y) + math.Floor(point.Z)
	if math.Mod(sum, 2) == 0 {
		return p.A
	}
	return p.B
}
