package asdf

import (
	"github.com/oleg/raytracer-go/geom"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_creating_stripe_pattern(t *testing.T) {
	pattern := MakeStripePattern(geom.White, geom.Black)

	assert.Equal(t, geom.White, pattern.A)
	assert.Equal(t, geom.Black, pattern.B)
}

func Test_stripe_pattern_is_constant_in_y(t *testing.T) {
	pattern := MakeStripePattern(geom.White, geom.Black)

	assert.Equal(t, geom.White, pattern.PatternAt(geom.Point{X: 0, Y: 0, Z: 0}))
	assert.Equal(t, geom.White, pattern.PatternAt(geom.Point{X: 0, Y: 1, Z: 0}))
	assert.Equal(t, geom.White, pattern.PatternAt(geom.Point{X: 0, Y: 2, Z: 0}))
}

func Test_stripe_pattern_is_constant_in_z(t *testing.T) {
	pattern := MakeStripePattern(geom.White, geom.Black)

	assert.Equal(t, geom.White, pattern.PatternAt(geom.Point{X: 0, Y: 0, Z: 0}))
	assert.Equal(t, geom.White, pattern.PatternAt(geom.Point{X: 0, Y: 0, Z: 1}))
	assert.Equal(t, geom.White, pattern.PatternAt(geom.Point{X: 0, Y: 0, Z: 2}))
}

func Test_stripe_pattern_alternates_in_x(t *testing.T) {
	pattern := MakeStripePattern(geom.White, geom.Black)

	assert.Equal(t, geom.White, pattern.PatternAt(geom.Point{X: 0, Y: 0, Z: 0}))
	assert.Equal(t, geom.White, pattern.PatternAt(geom.Point{X: 0.9, Y: 0, Z: 0}))

	assert.Equal(t, geom.Black, pattern.PatternAt(geom.Point{X: 1, Y: 0, Z: 0}))
	assert.Equal(t, geom.Black, pattern.PatternAt(geom.Point{X: -0.1, Y: 0, Z: 0}))

	assert.Equal(t, geom.Black, pattern.PatternAt(geom.Point{X: -1, Y: 0, Z: 0}))
	assert.Equal(t, geom.White, pattern.PatternAt(geom.Point{X: -1.1, Y: 0, Z: 0}))
}
//todo:oleg uncomment when no deps on figure package
//func Test_stripes_with_object_transformation(t *testing.T) {
//	object := figure.NewSphere(geom.Scaling(2, 2, 2), DefaultMaterial())
//	pattern := MakeStripePattern(geom.White, geom.Black)
//
//	c := PatternAtShape(pattern, object, geom.Point{X: 1.5, Y: 0, Z: 0})
//
//	assert.Equal(t, geom.White, c)
//}
//
//func Test_stripes_with_pattern_transformation(t *testing.T) {
//	object := figure.NewSphere(geom.IdentityMatrix(), DefaultMaterial())
//	pattern := MakeStripePatternT(geom.White, geom.Black, geom.Scaling(2, 2, 2))
//
//	c := PatternAtShape(pattern, object, geom.Point{X: 1.5, Y: 0, Z: 0})
//
//	assert.Equal(t, geom.White, c)
//}
//
//func Test_stripes_with_both_object_and_pattern_transformation(t *testing.T) {
//	object := figure.NewSphere(geom.Scaling(2, 2, 2), DefaultMaterial())
//	pattern := MakeStripePatternT(geom.White, geom.Black, geom.Translation(0.5, 0, 0))
//
//	c := PatternAtShape(pattern, object, geom.Point{X: 2.5, Y: 0, Z: 0})
//
//	assert.Equal(t, geom.White, c)
//}

func Test_gradient_linearly_interpolates_between_colors(t *testing.T) {
	pattern := MakeGradientPattern(geom.White, geom.Black)

	assert.Equal(t, geom.Color{R: 1, G: 1, B: 1}, pattern.PatternAt(geom.Point{X: 0, Y: 0, Z: 0}))
	assert.Equal(t, geom.Color{R: 0.75, G: 0.75, B: 0.75}, pattern.PatternAt(geom.Point{X: 0.25, Y: 0, Z: 0}))
	assert.Equal(t, geom.Color{R: 0.5, G: 0.5, B: 0.5}, pattern.PatternAt(geom.Point{X: 0.5, Y: 0, Z: 0}))
	assert.Equal(t, geom.Color{R: 0.25, G: 0.25, B: 0.25}, pattern.PatternAt(geom.Point{X: 0.75, Y: 0, Z: 0}))
}

func Test_ring_should_extend_in_both_x_and_z(t *testing.T) {
	pattern := MakeRingPattern(geom.White, geom.Black)

	assert.Equal(t, geom.White, pattern.PatternAt(geom.Point{X: 0, Y: 0, Z: 0}))
	assert.Equal(t, geom.Black, pattern.PatternAt(geom.Point{X: 1, Y: 0, Z: 0}))
	assert.Equal(t, geom.Black, pattern.PatternAt(geom.Point{X: 0, Y: 0, Z: 1}))
	assert.Equal(t, geom.Black, pattern.PatternAt(geom.Point{X: 0.708, Y: 0, Z: 0.708}))
}

func Test_checkers_should_repeat_in_x_y_z(t *testing.T) {
	pattern := MakeCheckersPattern(geom.White, geom.Black)

	assert.Equal(t, geom.White, pattern.PatternAt(geom.Point{X: 0, Y: 0, Z: 0}))
	assert.Equal(t, geom.White, pattern.PatternAt(geom.Point{X: 0.99, Y: 0, Z: 0}))
	assert.Equal(t, geom.Black, pattern.PatternAt(geom.Point{X: 1.01, Y: 0, Z: 0}))

	assert.Equal(t, geom.White, pattern.PatternAt(geom.Point{X: 0, Y: 0, Z: 0}))
	assert.Equal(t, geom.White, pattern.PatternAt(geom.Point{X: 0, Y: 0.99, Z: 0}))
	assert.Equal(t, geom.Black, pattern.PatternAt(geom.Point{X: 0, Y: 1.01, Z: 0}))

	assert.Equal(t, geom.White, pattern.PatternAt(geom.Point{X: 0, Y: 0, Z: 0}))
	assert.Equal(t, geom.White, pattern.PatternAt(geom.Point{X: 0, Y: 0, Z: 0.99}))
	assert.Equal(t, geom.Black, pattern.PatternAt(geom.Point{X: 0, Y: 0, Z: 1.01}))
}
