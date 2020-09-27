package figure

import (
	"github.com/stretchr/testify/assert"
	"gray/multid"
	"gray/oned"
	"math"
	"testing"
)

func Test_ray_intersects_sphere_at_two_points(t *testing.T) {
	tests := []struct {
		name     string
		ray      Ray
		expected []float64
	}{
		{"A ray intersects a sphere at two points",
			Ray{oned.Point{0, 0, -5}, oned.Vector{0, 0, 1}},
			[]float64{4, 6}},
		{"A ray intersects a sphere at a tangent",
			Ray{oned.Point{0, 1, -5}, oned.Vector{0, 0, 1}},
			[]float64{5, 5}},
		{"A ray misses a sphere",
			Ray{oned.Point{0, 2, -5}, oned.Vector{0, 0, 1}},
			[]float64{}},
		{"A ray originates inside a sphere",
			Ray{oned.Point{0, 0, 0}, oned.Vector{0, 0, 1}},
			[]float64{-1, 1}},
		{"A sphere is behind a ray",
			Ray{oned.Point{0, 0, 5}, oned.Vector{0, 0, 1}},
			[]float64{-6, -4}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := MakeSphere()

			xs := Intersect(s, test.ray)

			assert.Len(t, xs, len(test.expected))
			for i, expected := range test.expected {
				assert.Equal(t, expected, xs[i].Distance)
			}
		})
	}
}

func Test_intersect_sets_object_on_intersection(t *testing.T) {
	r := Ray{oned.Point{0, 0, -5}, oned.Vector{0, 0, 1}}
	s := MakeSphere()

	res := Intersect(s, r)

	assert.Equal(t, 2, len(res))
	assert.Equal(t, s, res[0].Object)
	assert.Equal(t, s, res[1].Object)
}

func Test_sphere_default_transformation(t *testing.T) {
	s := MakeSphere()

	r := s.Transform()

	assert.Equal(t, multid.IdentityMatrix, r)
}
func Test_changing_sphere_transformation(t *testing.T) {
	tr := multid.Translation(2, 3, 4)
	s := MakeSphereT(tr)

	r := s.Transform()

	assert.Equal(t, tr, r)
}

func Test_intersecting_scaled_sphere_with_ray(t *testing.T) {
	r := Ray{oned.Point{0, 0, -5}, oned.Vector{0, 0, 1}}
	s := MakeSphereT(multid.Scaling(2, 2, 2))

	xs := Intersect(s, r) //todo table test for intersect

	assert.Equal(t, 2, len(xs))
	assert.Equal(t, 3., xs[0].Distance)
	assert.Equal(t, 7., xs[1].Distance)
}
func Test_intersecting_translated_sphere_with_ray(t *testing.T) {
	r := Ray{oned.Point{0, 0, -5}, oned.Vector{0, 0, 1}}
	s := MakeSphereT(multid.Translation(5, 0, 0))

	xs := Intersect(s, r)

	assert.Equal(t, 0, len(xs))
}

func Test_normal_on_sphere(t *testing.T) {
	sqrt3d3 := math.Sqrt(3) / 3

	tests := []struct {
		name     string
		point    oned.Point
		expected oned.Vector
	}{
		{"The normal on a sphere at a point on the x axis",
			oned.Point{1, 0, 0}, oned.Vector{1, 0, 0}},
		{"The normal on a sphere at a point on the y axis",
			oned.Point{0, 1, 0}, oned.Vector{0, 1, 0}},
		{"The normal on a sphere at a point on the z axis",
			oned.Point{0, 0, 1}, oned.Vector{0, 0, 1}},
		{"The normal on a sphere at a non axial point",
			oned.Point{sqrt3d3, sqrt3d3, sqrt3d3}, oned.Vector{sqrt3d3, sqrt3d3, sqrt3d3}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := MakeSphere()

			r := NormalAt(s, test.point)

			assert.Equal(t, test.expected, r)
		})
	}
}

func Test_normal_is_normalized_vector(t *testing.T) {
	sqrt3d3 := math.Sqrt(3) / 3
	s := MakeSphere()

	r := NormalAt(s, oned.Point{sqrt3d3, sqrt3d3, sqrt3d3})

	assert.Equal(t, r.Normalize(), r)
}

func Test_computing_normal_on_translated_sphere(t *testing.T) {
	s := MakeSphereT(multid.Translation(0, 1, 0))

	n := NormalAt(s, oned.Point{0, 1.70711, -0.70711})

	oned.AssertVectorEqualInDelta(t, oned.Vector{0, 0.70711, -0.70711}, n)
}

func Test_computing_normal_on_transformed_sphere(t *testing.T) {
	s := MakeSphereT(multid.Scaling(1, 0.5, 1).Multiply(multid.RotationZ(math.Pi / 5)))

	n := NormalAt(s, oned.Point{0, math.Sqrt2 / 2, -math.Sqrt2 / 2})

	oned.AssertVectorEqualInDelta(t, oned.Vector{0, 0.97014, -0.24254}, n)
}

func Test_sphere_has_default_material(t *testing.T) {
	s := MakeSphere()

	assert.Equal(t, DefaultMaterial(), s.Material())
}

func Test_sphere_may_be_assigned_material(t *testing.T) {
	m := Material{Ambient: 1}
	s := MakeSphereM(m)

	assert.Equal(t, m, s.Material())
}

func Test_helper_for_producing_sphere_with_glassy_material(t *testing.T) {
	s := MakeGlassSphere()

	assert.Equal(t, multid.IdentityMatrix, s.Transform())
	assert.Equal(t, 1.0, s.Material().Transparency)
	assert.Equal(t, 1.5, s.Material().RefractiveIndex)
}

