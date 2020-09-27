package figure

import (
	"github.com/stretchr/testify/assert"
	"gray/multid"
	"gray/oned"
	"testing"
)

func Test_creating_and_querying_a_ray(t *testing.T) {
	origin := oned.Point{1, 2, 3}
	direction := oned.Vector{4, 5, 6}

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
		{"0", 0, oned.Point{2, 3, 4}},
		{"1", 1, oned.Point{3, 3, 4}},
		{"-1", -1, oned.Point{1, 3, 4}},
		{"2.5", 2.5, oned.Point{4.5, 3, 4}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := Ray{oned.Point{2, 3, 4}, oned.Vector{1, 0, 0}}
			p := r.Position(test.distance)
			assert.Equal(t, test.expected, p)
		})
	}
}

func Test_translating_ray(t *testing.T) {
	r := Ray{oned.Point{1, 2, 3}, oned.Vector{0, 1, 0}}
	m := multid.Translation(3, 4, 5)

	r2 := r.Transform(m)

	assert.Equal(t, oned.Point{4, 6, 8}, r2.Origin)
	assert.Equal(t, oned.Vector{0, 1, 0}, r2.Direction)
}

func Test_scaling_ray(t *testing.T) {
	r := Ray{oned.Point{1, 2, 3}, oned.Vector{0, 1, 0}}
	m := multid.Scaling(2, 3, 4)

	r2 := r.Transform(m)

	assert.Equal(t, oned.Point{2, 6, 12}, r2.Origin)
	assert.Equal(t, oned.Vector{0, 3, 0}, r2.Direction)
}
