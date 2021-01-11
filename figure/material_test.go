package figure

import (
	"github.com/oleg/raytracer-go/oned"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_default_material(t *testing.T) {
	m := DefaultMaterial()

	assert.Equal(t, oned.White, m.Color)
	assert.Equal(t, 0.1, m.Ambient)
	assert.Equal(t, 0.9, m.Diffuse)
	assert.Equal(t, 0.9, m.Specular)
	assert.Equal(t, 200.0, m.Shininess)
}

func Test_lighting(t *testing.T) {
	tests := []struct {
		name     string
		eyev     oned.Vector
		normalv  oned.Vector
		light    PointLight
		expected oned.Color
	}{
		{"Lighting with the eye between the light and the surface",
			oned.Vector{X: 0, Y: 0, Z: -1},
			oned.Vector{X: 0, Y: 0, Z: -1},
			PointLight{oned.Point{X: 0, Y: 0, Z: -10}, oned.White},
			oned.Color{R: 1.9, G: 1.9, B: 1.9}},
		{"Lighting with the eye between light and surface, eye offset 45°",
			oned.Vector{X: 0, Y: math.Sqrt2 / 2, Z: -math.Sqrt2 / 2},
			oned.Vector{X: 0, Y: 0, Z: -1},
			PointLight{oned.Point{X: 0, Y: 0, Z: -10}, oned.White},
			oned.White},
		{"Lighting with eye opposite surface, light offset 45°",
			oned.Vector{X: 0, Y: 0, Z: -1},
			oned.Vector{X: 0, Y: 0, Z: -1},
			PointLight{oned.Point{X: 0, Y: 10, Z: -10}, oned.White},
			oned.Color{R: 0.7364, G: 0.7364, B: 0.7364}},
		{"Lighting with eye in the path of the reflection vector",
			oned.Vector{X: 0, Y: -math.Sqrt2 / 2, Z: -math.Sqrt2 / 2},
			oned.Vector{X: 0, Y: 0, Z: -1},
			PointLight{oned.Point{X: 0, Y: 10, Z: -10}, oned.White},
			oned.Color{R: 1.6364, G: 1.6364, B: 1.6364}},
		{"Lighting with the light behind the surface",
			oned.Vector{X: 0, Y: 0, Z: -1},
			oned.Vector{X: 0, Y: 0, Z: -1},
			PointLight{oned.Point{X: 0, Y: 0, Z: 10}, oned.White},
			oned.Color{R: 0.1, G: 0.1, B: 0.1}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			color := Lighting(DefaultMaterial(), MakeSphere(), test.light, oned.Point{}, test.eyev, test.normalv, false)

			oned.AssertColorEqualInDelta(t, test.expected, color)
		})
	}
}

func Test_lighting_with_surface_in_shadow(t *testing.T) {
	m := DefaultMaterial()
	eyeV := oned.Vector{X: 0, Y: 0, Z: -1}
	normalV := oned.Vector{X: 0, Y: 0, Z: -1}
	light := PointLight{oned.Point{X: 0, Y: 0, Z: -10}, oned.White}

	r := Lighting(m, MakeSphere(), light, oned.Point{}, eyeV, normalV, true)

	assert.Equal(t, oned.Color{R: 0.1, G: 0.1, B: 0.1}, r)
}

func Test_shadow(t *testing.T) {
	tests := []struct {
		name     string
		point    oned.Point
		expected bool
	}{
		{"There is no shadow when nothing is collinear with point and light",
			oned.Point{X: 0, Y: 10, Z: 0}, false},
		{"The shadow when an object is between the point and the light",
			oned.Point{X: 10, Y: -10, Z: 10}, true},
		{"There is no shadow when an object is behind the light",
			oned.Point{X: -20, Y: 20, Z: -20}, false},
		{"There is no shadow when an object is behind the point",
			oned.Point{X: -2, Y: 2, Z: -2}, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := defaultWorld()

			r := w.IsShadowed(test.point)

			assert.Equal(t, test.expected, r)
		})
	}
}

func Test_Lighting_with_pattern_applied(t *testing.T) {
	m := MakeMaterialBuilder().
		SetAmbient(1).
		SetDiffuse(0).
		SetSpecular(0).
		SetPattern(MakeStripePattern(oned.White, oned.Black)).
		Build()

	eyeV := oned.Vector{X: 0, Y: 0, Z: -1}
	normalV := oned.Vector{X: 0, Y: 0, Z: -1}
	light := PointLight{oned.Point{X: 0, Y: 0, Z: -10}, oned.White}
	c1 := Lighting(m, MakeSphere(), light, oned.Point{X: 0.9, Y: 0, Z: 0}, eyeV, normalV, false)
	c2 := Lighting(m, MakeSphere(), light, oned.Point{X: 1.1, Y: 0, Z: 0}, eyeV, normalV, false)

	assert.Equal(t, oned.White, c1)
	assert.Equal(t, oned.Black, c2)
}
