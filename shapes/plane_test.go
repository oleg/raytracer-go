package shapes

import (
	"github.com/oleg/raytracer-go/geom"
	"github.com/oleg/raytracer-go/physic"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_normal_of_plane_is_constant_everywhere(t *testing.T) {
	p := NewPlane(geom.IdentityMatrix(), physic.DefaultMaterial())

	assert.Equal(t, geom.Vector{X: 0, Y: 1, Z: 0}, p.NormalAt(geom.Point{X: 0, Y: 0, Z: 0}))
	assert.Equal(t, geom.Vector{X: 0, Y: 1, Z: 0}, p.NormalAt(geom.Point{X: 10, Y: 0, Z: -10}))
	assert.Equal(t, geom.Vector{X: 0, Y: 1, Z: 0}, p.NormalAt(geom.Point{X: -5, Y: 0, Z: 150}))
}

func Test_intersect_with_ray_parallel_to_plane(t *testing.T) {
	p := NewPlane(geom.IdentityMatrix(), physic.DefaultMaterial())
	r := Ray{geom.Point{X: 0, Y: 10, Z: 0}, geom.Vector{X: 0, Y: 0, Z: 1}}

	xs := p.Intersect(r)

	assert.Empty(t, xs)
}

func Test_intersect_with_coplanar_ray(t *testing.T) {
	p := NewPlane(geom.IdentityMatrix(), physic.DefaultMaterial())
	r := Ray{geom.Point{X: 0, Y: 0, Z: 0}, geom.Vector{X: 0, Y: 0, Z: 1}}

	xs := p.Intersect(r)

	assert.Empty(t, xs)
}

func Test_ray_intersecting_plane_from_above(t *testing.T) {
	p := NewPlane(geom.IdentityMatrix(), physic.DefaultMaterial())
	r := Ray{geom.Point{X: 0, Y: 1, Z: 0}, geom.Vector{X: 0, Y: -1, Z: 0}}

	xs := p.Intersect(r)

	assert.Equal(t, 1, len(xs))
	assert.Equal(t, 1., xs[0].Distance)
	assert.Equal(t, p, xs[0].Object)
}

func Test_ray_intersecting_a_plane_from_below(t *testing.T) {
	p := NewPlane(geom.IdentityMatrix(), physic.DefaultMaterial())
	r := Ray{geom.Point{X: 0, Y: -1, Z: 0}, geom.Vector{X: 0, Y: 1, Z: 0}}

	xs := p.Intersect(r)

	assert.Equal(t, 1, len(xs))
	assert.Equal(t, 1., xs[0].Distance)
	assert.Equal(t, p, xs[0].Object)
}
