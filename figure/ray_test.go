package figure

import (
	"github.com/oleg/raytracer-go/multid"
	"github.com/oleg/raytracer-go/oned"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_creating_and_querying_a_ray(t *testing.T) {
	origin := oned.Point{X: 1, Y: 2, Z: 3}
	direction := oned.Vector{X: 4, Y: 5, Z: 6}

	ray := Ray{origin, direction}

	assert.Equal(t, origin, ray.Origin)
	assert.Equal(t, direction, ray.Direction)
}

func Test_Computing_point_from_distance(t *testing.T) {

	tests := []struct {
		name     string
		distance float64
		expected oned.Point
	}{
		{"0", 0, oned.Point{X: 2, Y: 3, Z: 4}},
		{"1", 1, oned.Point{X: 3, Y: 3, Z: 4}},
		{"-1", -1, oned.Point{X: 1, Y: 3, Z: 4}},
		{"2.5", 2.5, oned.Point{X: 4.5, Y: 3, Z: 4}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := Ray{oned.Point{X: 2, Y: 3, Z: 4}, oned.Vector{X: 1, Y: 0, Z: 0}}
			p := r.Position(test.distance)
			assert.Equal(t, test.expected, p)
		})
	}
}

func Test_translating_ray(t *testing.T) {
	r := Ray{oned.Point{X: 1, Y: 2, Z: 3}, oned.Vector{X: 0, Y: 1, Z: 0}}
	m := multid.Translation(3, 4, 5)

	r2 := r.Transform(m)

	assert.Equal(t, oned.Point{X: 4, Y: 6, Z: 8}, r2.Origin)
	assert.Equal(t, oned.Vector{X: 0, Y: 1, Z: 0}, r2.Direction)
}

func Test_scaling_ray(t *testing.T) {
	r := Ray{oned.Point{X: 1, Y: 2, Z: 3}, oned.Vector{X: 0, Y: 1, Z: 0}}
	m := multid.Scaling(2, 3, 4)

	r2 := r.Transform(m)

	assert.Equal(t, oned.Point{X: 2, Y: 6, Z: 12}, r2.Origin)
	assert.Equal(t, oned.Vector{X: 0, Y: 3, Z: 0}, r2.Direction)
}
