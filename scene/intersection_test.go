package scene

import (
	"github.com/oleg/raytracer-go/geom"
	"github.com/oleg/raytracer-go/physic"
	"github.com/oleg/raytracer-go/shapes"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_intersection_encapsulates_distance_and_object(t *testing.T) {
	s := shapes.NewSphere(geom.IdentityMatrix(), physic.DefaultMaterial())

	i := shapes.Inter{3.5, s}

	assert.Equal(t, 3.5, i.Distance)
	assert.Equal(t, s, i.Object)
}

func Test_aggregating_intersections(t *testing.T) {
	s := shapes.NewSphere(geom.IdentityMatrix(), physic.DefaultMaterial())

	i1 := shapes.Inter{1, s}
	i2 := shapes.Inter{2, s}

	xs := shapes.Inters{i1, i2}

	assert.Equal(t, xs[0].Distance, 1.)
	assert.Equal(t, xs[1].Distance, 2.)
}

func Test_hit_when_all_intersections_have_positive_distance(t *testing.T) {
	s := shapes.NewSphere(geom.IdentityMatrix(), physic.DefaultMaterial())

	i1 := shapes.Inter{1, s}
	i2 := shapes.Inter{2, s}
	xs := shapes.Inters{i2, i1}

	_, i := Hit(xs)

	assert.Equal(t, i1, i)
}

func Test_hit_intersections(t *testing.T) {
	s := shapes.NewSphere(geom.IdentityMatrix(), physic.DefaultMaterial())
	tests := []struct {
		name                 string
		intersections        shapes.Inters
		expectedFound        bool
		expectedIntersection shapes.Inter
	}{
		{"all intersections have positive t",
			shapes.Inters{shapes.Inter{2, s}, shapes.Inter{1, s}}, true, shapes.Inter{1, s}},
		{"some intersections have negative t",
			shapes.Inters{shapes.Inter{1, s}, shapes.Inter{-1, s}}, true, shapes.Inter{1, s}},
		{"all intersections have negative t",
			shapes.Inters{shapes.Inter{-1, s}, shapes.Inter{-2, s}}, false, shapes.Inter{}},
		{"is always the lowest non negative intersection",
			shapes.Inters{shapes.Inter{5, s}, shapes.Inter{7, s}, shapes.Inter{-3, s}, shapes.Inter{2, s}}, true, shapes.Inter{2, s}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			ok, hit := Hit(test.intersections)

			assert.Equal(t, test.expectedFound, ok)
			assert.Equal(t, test.expectedIntersection, hit)
		})
	}
}

func Test_precomputing_state_of_intersection(t *testing.T) {
	r := shapes.Ray{geom.Point{X: 0, Y: 0, Z: -5}, geom.Vector{X: 0, Y: 0, Z: 1}}
	shape := shapes.NewSphere(geom.IdentityMatrix(), physic.DefaultMaterial())
	i := shapes.Inter{4, shape}

	comps := NewComputations(i, r, shapes.Inters{i})

	assert.Equal(t, i.Distance, comps.Distance)
	assert.Equal(t, i.Object, comps.Object)
	assert.Equal(t, geom.Point{X: 0, Y: 0, Z: -1}, comps.Point)
	assert.Equal(t, geom.Vector{X: 0, Y: 0, Z: -1}, comps.EyeV)
	assert.Equal(t, geom.Vector{X: 0, Y: 0, Z: -1}, comps.NormalV)
}

func Test_hit_when_intersection_occurs_on_outside(t *testing.T) {
	r := shapes.Ray{geom.Point{X: 0, Y: 0, Z: -5}, geom.Vector{X: 0, Y: 0, Z: 1}}
	shape := shapes.NewSphere(geom.IdentityMatrix(), physic.DefaultMaterial())
	i := shapes.Inter{4, shape}

	comps := NewComputations(i, r, shapes.Inters{i})

	assert.Equal(t, false, comps.Inside)
}

func Test_hit_when_intersection_occurs_on_inside(t *testing.T) {
	r := shapes.Ray{geom.Point{X: 0, Y: 0, Z: 0}, geom.Vector{X: 0, Y: 0, Z: 1}}
	shape := shapes.NewSphere(geom.IdentityMatrix(), physic.DefaultMaterial())
	i := shapes.Inter{1, shape}

	comps := NewComputations(i, r, shapes.Inters{i})

	assert.Equal(t, geom.Point{X: 0, Y: 0, Z: 1}, comps.Point)
	assert.Equal(t, geom.Vector{X: 0, Y: 0, Z: -1}, comps.EyeV)
	assert.Equal(t, geom.Vector{X: 0, Y: 0, Z: -1}, comps.NormalV)
	assert.Equal(t, true, comps.Inside)
}

func Test_precomputing_reflection_vector(t *testing.T) {
	shape := shapes.NewPlane(geom.IdentityMatrix(), physic.DefaultMaterial())
	ray := shapes.Ray{geom.Point{X: 0, Y: 1, Z: -1}, geom.Vector{X: 0, Y: -math.Sqrt2 / 2, Z: math.Sqrt2 / 2}}
	i := shapes.Inter{math.Sqrt2, shape}

	comps := NewComputations(i, ray, shapes.Inters{i})

	assert.Equal(t, geom.Vector{X: 0, Y: math.Sqrt2 / 2, Z: math.Sqrt2 / 2}, comps.ReflectV)
}

func Test_finding_n1_and_n2_at_various_intersections(t *testing.T) {
	tests := []struct {
		name  string
		index int
		n1    float64
		n2    float64
	}{
		{"case 0", 0, 1.0, 1.5},
		{"case 1", 1, 1.5, 2.0},
		{"case 2", 2, 2.0, 2.5},
		{"case 3", 3, 2.5, 2.5},
		{"case 4", 4, 2.5, 1.5},
		{"case 5", 5, 1.5, 1.0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			a := shapes.NewSphere(geom.Scaling(2, 2, 2), physic.GlassMaterialBuilder().SetRefractiveIndex(1.5).Build())
			b := shapes.NewSphere(geom.Translation(0, 0, -0.25), physic.GlassMaterialBuilder().SetRefractiveIndex(2.0).Build())
			c := shapes.NewSphere(geom.Translation(0, 0, 0.25), physic.GlassMaterialBuilder().SetRefractiveIndex(2.5).Build())

			r := shapes.Ray{geom.Point{X: 0, Y: 0, Z: -4}, geom.Vector{X: 0, Y: 0, Z: 1}}
			xs := shapes.Inters{shapes.Inter{2, a}, shapes.Inter{2.75, b}, shapes.Inter{3.25, c}, shapes.Inter{4.75, b}, shapes.Inter{5.25, c}, shapes.Inter{6, a}}

			comps := NewComputations(xs[test.index], r, xs)

			assert.Equal(t, test.n1, comps.N1)
			assert.Equal(t, test.n2, comps.N2)
		})
	}
}

func Test_under_point_is_offset_below_surface(t *testing.T) {
	ray := shapes.Ray{geom.Point{X: 0, Y: 0, Z: -5}, geom.Vector{X: 0, Y: 0, Z: 1}}
	shape := shapes.NewSphere(geom.Translation(0, 0, 1), physic.GlassMaterialBuilder().Build())
	i := shapes.Inter{5, shape}
	xs := shapes.Inters{i}

	comps := NewComputations(i, ray, xs)

	assert.Greater(t, comps.UnderPoint.Z, -geom.Delta/2)
	assert.Greater(t, comps.UnderPoint.Z, comps.Point.Z)
}

func Test_schlick_approximation_under_total_internal_reflection(t *testing.T) {
	s := shapes.NewSphere(geom.IdentityMatrix(), physic.GlassMaterialBuilder().Build())
	r := shapes.Ray{geom.Point{X: 0, Y: 0, Z: math.Sqrt2 / 2}, geom.Vector{X: 0, Y: 1, Z: 0}}
	xs := shapes.Inters{shapes.Inter{-math.Sqrt2 / 2, s}, shapes.Inter{math.Sqrt2 / 2, s}}
	comps := NewComputations(xs[1], r, xs)

	reflectance := Schlick(comps)

	assert.Equal(t, 1.0, reflectance)
}

func Test_schlick_approximation_with_perpendicular_viewing_angle(t *testing.T) {
	s := shapes.NewSphere(geom.IdentityMatrix(), physic.GlassMaterialBuilder().Build())
	r := shapes.Ray{geom.Point{X: 0, Y: 0, Z: 0}, geom.Vector{X: 0, Y: 1, Z: 0}}
	xs := shapes.Inters{shapes.Inter{-1, s}, shapes.Inter{1, s}}
	comps := NewComputations(xs[1], r, xs)

	reflectance := Schlick(comps)

	assert.InDelta(t, 0.04, reflectance, geom.Delta)
}

func Test_schlick_approximation_with_small_angle_and_n2_gt_n1(t *testing.T) {
	s := shapes.NewSphere(geom.IdentityMatrix(), physic.GlassMaterialBuilder().Build())
	r := shapes.Ray{geom.Point{X: 0, Y: 0.99, Z: -2}, geom.Vector{X: 0, Y: 0, Z: 1}}
	xs := shapes.Inters{shapes.Inter{1.8589, s}}
	comps := NewComputations(xs[0], r, xs)

	reflectance := Schlick(comps)

	assert.InDelta(t, 0.48873, reflectance, geom.Delta)
}
