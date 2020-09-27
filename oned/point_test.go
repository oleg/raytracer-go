package oned

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_make_point_from_tuple(t *testing.T) {
	point := Point{1.1, 2.1, 3.1}
	tuple := Tuple{1.1, 2.1, 3.1}

	assert.Equal(t, point, Point(tuple))
}

func Test_addVector_gives_point(t *testing.T) {
	p1 := Point{1, 2, 3}
	v1 := Vector{2, 3, 4}

	point := p1.AddVector(v1)

	assert.Equal(t, Point{3, 5, 7}, point)
}

func Test_subtractPoint_gives_vector(t *testing.T) {
	p1 := Point{3, 2, 1}
	p2 := Point{5, 6, 7}

	vector := p1.SubtractPoint(p2)

	assert.Equal(t, Vector{-2, -4, -6}, vector)
}

func Test_subtractVector_gives_point(t *testing.T) {
	p := Point{3, 2, 1}
	v := Vector{5, 6, 7}

	point := p.SubtractVector(v)

	assert.Equal(t, Point{-2, -4, -6}, point)
}
