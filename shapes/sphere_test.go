package shapes

import (
	"github.com/oleg/raytracer-go/geom"
	"github.com/oleg/raytracer-go/physic"
	"github.com/stretchr/testify/assert"
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
			Ray{geom.Point{X: 0, Y: 0, Z: -5}, geom.Vector{X: 0, Y: 0, Z: 1}},
			[]float64{4, 6}},
		{"A ray intersects a sphere at a tangent",
			Ray{geom.Point{X: 0, Y: 1, Z: -5}, geom.Vector{X: 0, Y: 0, Z: 1}},
			[]float64{5, 5}},
		{"A ray misses a sphere",
			Ray{geom.Point{X: 0, Y: 2, Z: -5}, geom.Vector{X: 0, Y: 0, Z: 1}},
			[]float64{}},
		{"A ray originates inside a sphere",
			Ray{geom.Point{X: 0, Y: 0, Z: 0}, geom.Vector{X: 0, Y: 0, Z: 1}},
			[]float64{-1, 1}},
		{"A sphere is behind a ray",
			Ray{geom.Point{X: 0, Y: 0, Z: 5}, geom.Vector{X: 0, Y: 0, Z: 1}},
			[]float64{-6, -4}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := NewSphere(geom.IdentityMatrix(), physic.DefaultMaterial())

			xs := intersect(s, test.ray)

			assert.Len(t, xs, len(test.expected))
			for i, expected := range test.expected {
				assert.Equal(t, expected, xs[i].Distance)
			}
		})
	}
}

func Test_intersect_sets_object_on_intersection(t *testing.T) {
	r := Ray{geom.Point{X: 0, Y: 0, Z: -5}, geom.Vector{X: 0, Y: 0, Z: 1}}
	s := NewSphere(geom.IdentityMatrix(), physic.DefaultMaterial())

	res := intersect(s, r)

	assert.Equal(t, 2, len(res))
	assert.Equal(t, s, res[0].Object)
	assert.Equal(t, s, res[1].Object)
}

func Test_sphere_default_transformation(t *testing.T) {
	s := NewSphere(geom.IdentityMatrix(), physic.DefaultMaterial())

	r := s.Transformation()

	assert.Equal(t, geom.IdentityMatrix(), r)
}
func Test_changing_sphere_transformation(t *testing.T) {
	tr := geom.Translation(2, 3, 4)
	s := NewSphere(tr, physic.DefaultMaterial())

	r := s.Transformation()

	assert.Equal(t, tr, r)
}

func Test_intersecting_scaled_sphere_with_ray(t *testing.T) {
	r := Ray{geom.Point{X: 0, Y: 0, Z: -5}, geom.Vector{X: 0, Y: 0, Z: 1}}
	s := NewSphere(geom.Scaling(2, 2, 2), physic.DefaultMaterial())

	xs := intersect(s, r) //todo table test for intersect

	assert.Equal(t, 2, len(xs))
	assert.Equal(t, 3., xs[0].Distance)
	assert.Equal(t, 7., xs[1].Distance)
}

func Test_intersecting_translated_sphere_with_ray(t *testing.T) {
	r := Ray{geom.Point{X: 0, Y: 0, Z: -5}, geom.Vector{X: 0, Y: 0, Z: 1}}
	s := NewSphere(geom.Translation(5, 0, 0), physic.DefaultMaterial())

	xs := intersect(s, r)

	assert.Equal(t, 0, len(xs))
}

func Test_normal_on_sphere(t *testing.T) {
	sqrt3d3 := math.Sqrt(3) / 3

	tests := []struct {
		name     string
		point    geom.Point
		expected geom.Vector
	}{
		{"The normal on a sphere at a point on the x axis",
			geom.Point{X: 1, Y: 0, Z: 0}, geom.Vector{X: 1, Y: 0, Z: 0}},
		{"The normal on a sphere at a point on the y axis",
			geom.Point{X: 0, Y: 1, Z: 0}, geom.Vector{X: 0, Y: 1, Z: 0}},
		{"The normal on a sphere at a point on the z axis",
			geom.Point{X: 0, Y: 0, Z: 1}, geom.Vector{X: 0, Y: 0, Z: 1}},
		{"The normal on a sphere at a non axial point",
			geom.Point{X: sqrt3d3, Y: sqrt3d3, Z: sqrt3d3}, geom.Vector{X: sqrt3d3, Y: sqrt3d3, Z: sqrt3d3}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := NewSphere(geom.IdentityMatrix(), physic.DefaultMaterial())

			r := normalAt(s, test.point)

			assert.Equal(t, test.expected, r)
		})
	}
}

func Test_normal_is_normalized_vector(t *testing.T) {
	sqrt3d3 := math.Sqrt(3) / 3
	s := NewSphere(geom.IdentityMatrix(), physic.DefaultMaterial())

	r := normalAt(s, geom.Point{X: sqrt3d3, Y: sqrt3d3, Z: sqrt3d3})

	assert.Equal(t, r.Normalize(), r)
}

func Test_computing_normal_on_translated_sphere(t *testing.T) {
	s := NewSphere(geom.Translation(0, 1, 0), physic.DefaultMaterial())

	n := normalAt(s, geom.Point{X: 0, Y: 1.70711, Z: -0.70711})

	geom.AssertVectorEqualInDelta(t, geom.Vector{X: 0, Y: 0.70711, Z: -0.70711}, n)
}

func Test_computing_normal_on_transformed_sphere(t *testing.T) {
	s := NewSphere(geom.Scaling(1, 0.5, 1).Multiply(geom.RotationZ(math.Pi/5)), physic.DefaultMaterial())

	n := normalAt(s, geom.Point{X: 0, Y: math.Sqrt2 / 2, Z: -math.Sqrt2 / 2})

	geom.AssertVectorEqualInDelta(t, geom.Vector{X: 0, Y: 0.97014, Z: -0.24254}, n)
}

func Test_sphere_has_default_material(t *testing.T) {
	s := NewSphere(geom.IdentityMatrix(), physic.DefaultMaterial())

	assert.Equal(t, physic.DefaultMaterial(), s.Material())
}

func Test_sphere_may_be_assigned_material(t *testing.T) {
	m := &physic.Material{Ambient: 1}
	s := NewSphere(geom.IdentityMatrix(), m)

	assert.Equal(t, m, s.Material())
}

func Test_helper_for_producing_sphere_with_glassy_material(t *testing.T) {
	s := NewGlassSphere()

	assert.Equal(t, geom.IdentityMatrix(), s.Transformation())
	assert.Equal(t, 1.0, s.Material().Transparency)
	assert.Equal(t, 1.5, s.Material().RefractiveIndex)
}

//todo updates tests, move ray transformation
func intersect(shape Shape, ray Ray) Intersections {
	return shape.Intersect(ray.ToLocal(shape))
}

//todo updates tests, remove transformation and normalization
func normalAt(shape Shape, worldPoint geom.Point) geom.Vector {
	localPoint := shape.Transformation().Inverse().MultiplyPoint(worldPoint)
	localNormal := shape.NormalAt(localPoint)
	worldNormal := shape.Transformation().Inverse().Transpose().MultiplyVector(localNormal)
	return worldNormal.Normalize()
}
