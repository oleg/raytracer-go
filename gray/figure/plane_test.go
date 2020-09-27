package figure

import (
	"github.com/stretchr/testify/assert"
	"gray/oned"
	"testing"
)

func Test_normal_of_plane_is_constant_everywhere(t *testing.T) {
	p := MakePlane()

	assert.Equal(t, oned.Vector{0, 1, 0}, p.LocalNormalAt(oned.Point{0, 0, 0}))
	assert.Equal(t, oned.Vector{0, 1, 0}, p.LocalNormalAt(oned.Point{10, 0, -10}))
	assert.Equal(t, oned.Vector{0, 1, 0}, p.LocalNormalAt(oned.Point{-5, 0, 150}))
}

func Test_intersect_with_ray_parallel_to_plane(t *testing.T) {
	p := MakePlane()
	r := Ray{oned.Point{0, 10, 0}, oned.Vector{0, 0, 1}}

	xs := p.LocalIntersect(r)

	assert.Empty(t, xs)
}

func Test_intersect_with_coplanar_ray(t *testing.T) {
	p := MakePlane()
	r := Ray{oned.Point{0, 0, 0}, oned.Vector{0, 0, 1}}

	xs := p.LocalIntersect(r)

	assert.Empty(t, xs)
}

func Test_ray_intersecting_plane_from_above(t *testing.T) {
	p := MakePlane()
	r := Ray{oned.Point{0, 1, 0}, oned.Vector{0, -1, 0}}

	xs := p.LocalIntersect(r)

	assert.Equal(t, 1, len(xs))
	assert.Equal(t, 1., xs[0].Distance)
	assert.Equal(t, p, xs[0].Object)
}

func Test_ray_intersecting_a_plane_from_below(t *testing.T) {
	p := MakePlane()
	r := Ray{oned.Point{0, -1, 0}, oned.Vector{0, 1, 0}}

	xs := p.LocalIntersect(r)

	assert.Equal(t, 1, len(xs))
	assert.Equal(t, 1., xs[0].Distance)
	assert.Equal(t, p, xs[0].Object)
}
