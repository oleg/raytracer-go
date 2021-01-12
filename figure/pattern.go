package figure

import (
	"github.com/oleg/raytracer-go/geom"
	"math"
)

type PatternFinder interface {
	PatternAt(point geom.Point) geom.Color
}

type Pattern interface {
	HasTransformation
	PatternFinder
}

func PatternAtShape(pattern Pattern, shape Shape, worldPoint geom.Point) geom.Color {
	objectPoint := shape.Transformation().Inverse().MultiplyPoint(worldPoint)
	patternPoint := pattern.Transformation().Inverse().MultiplyPoint(objectPoint)
	return pattern.PatternAt(patternPoint)
}

//todo refactor: remove duplicates
type StripePattern struct {
	A, B      geom.Color
	transform *geom.Matrix
}

func MakeStripePattern(A, B geom.Color) StripePattern {
	return StripePattern{A, B, geom.IdentityMatrix()}
}
func MakeStripePatternT(A, B geom.Color, transform *geom.Matrix) StripePattern {
	return StripePattern{A, B, transform}
}

func (p StripePattern) PatternAt(point geom.Point) geom.Color {
	if math.Mod(math.Floor(point.X), 2) == 0 {
		return p.A
	}
	return p.B
}
func (p StripePattern) Transformation() *geom.Matrix {
	return p.transform
}

type GradientPattern struct {
	A, B      geom.Color
	transform *geom.Matrix
}

func MakeGradientPattern(a, b geom.Color) GradientPattern {
	return GradientPattern{a, b, geom.IdentityMatrix()}
}
func MakeGradientPatternT(a, b geom.Color, transform *geom.Matrix) GradientPattern {
	return GradientPattern{a, b, transform}
}

func (p GradientPattern) PatternAt(point geom.Point) geom.Color {
	distance := p.B.Subtract(p.A)
	fraction := point.X - math.Floor(point.X)
	return p.A.Add(distance.MultiplyByScalar(fraction))
}

func (p GradientPattern) Transformation() *geom.Matrix {
	return p.transform
}

type RingPattern struct {
	A, B      geom.Color
	transform *geom.Matrix
}

func MakeRingPattern(a, b geom.Color) RingPattern {
	return RingPattern{a, b, geom.IdentityMatrix()}
}

func MakeRingPatternT(a, b geom.Color, transform *geom.Matrix) RingPattern {
	return RingPattern{a, b, transform}
}

func (p RingPattern) PatternAt(point geom.Point) geom.Color {
	hypot := math.Floor(math.Hypot(point.X, point.Z))
	if math.Mod(hypot, 2) == 0 {
		return p.A
	}
	return p.B
}

func (p RingPattern) Transformation() *geom.Matrix {
	return p.transform
}

type CheckersPattern struct {
	A, B      geom.Color
	transform *geom.Matrix
}

func MakeCheckersPattern(a, b geom.Color) CheckersPattern {
	return CheckersPattern{a, b, geom.IdentityMatrix()}
}

func MakeCheckersPatternT(a, b geom.Color, transform *geom.Matrix) CheckersPattern {
	return CheckersPattern{a, b, transform}
}

func (p CheckersPattern) PatternAt(point geom.Point) geom.Color {
	sum := math.Floor(point.X) + math.Floor(point.Y) + math.Floor(point.Z)
	if math.Mod(sum, 2) == 0 {
		return p.A
	}
	return p.B
}

func (p CheckersPattern) Transformation() *geom.Matrix {
	return p.transform
}
