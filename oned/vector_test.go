package oned

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_convert_tuple_to_vector(t *testing.T) {
	vector := Vector{1.1, 2.1, 3.1}
	tuple := Tuple{1.1, 2.1, 3.1}
	assert.Equal(t, vector, Vector(tuple))
}

func Test_addVector_gives_vector(t *testing.T) {
	v1 := Vector{1, 2, 3}
	v2 := Vector{2, 3, 4}

	vector := v1.AddVector(v2)

	assert.Equal(t, Vector{3, 5, 7}, vector)
}

func Test_subtractVector_gives_vector(t *testing.T) {
	v1 := Vector{5, 2, 2}
	v2 := Vector{5, 6, 1}

	vector := v1.SubtractVector(v2)

	assert.Equal(t, Vector{0, -4, 1}, vector)
}

func Test_subtruct_zero_vector_negates_vector(t *testing.T) {
	zv := Vector{0, 0, 0}
	v1 := Vector{1, -2, 3}

	vector := zv.SubtractVector(v1)

	assert.Equal(t, Vector{-1, 2, -3}, vector)
}

func Test_negate_negates_all_points(t *testing.T) {
	v := Vector{1, -2, 3}

	vector := v.Negate()

	assert.Equal(t, Vector{-1, 2, -3}, vector)
}

func Test_magnitude_of_1_0_0(t *testing.T) {
	v := Vector{1, 0, 0}

	result := v.Magnitude()

	assert.Equal(t, 1.0, result)
}

func Test_magnitude_of_0_1_0(t *testing.T) {
	v := Vector{0, 1, 0}

	result := v.Magnitude()

	assert.Equal(t, 1.0, result)
}

func Test_magnitude_of_0_0_1(t *testing.T) {
	v := Vector{0, 0, 1}

	result := v.Magnitude()

	assert.Equal(t, 1.0, result)
}

func Test_magnitude_of_1_2_3(t *testing.T) {
	v := Vector{1, 2, 3}

	result := v.Magnitude()

	assert.Equal(t, math.Sqrt(14), result)
}

func Test_magnitude_of_m1_m2_m3(t *testing.T) {
	v := Vector{-1, -2, -3}

	result := v.Magnitude()

	assert.Equal(t, math.Sqrt(14), result)
}

func Test_normalizing_vector_4_0_0(t *testing.T) {
	v := Vector{4, 0, 0}

	result := v.Normalize()

	assert.Equal(t, Vector{1, 0, 0}, result)
}

func Test_normalizing_vector_1_2_3(t *testing.T) {
	v := Vector{1, 2, 3}

	result := v.Normalize()

	AssertVectorEqualInDelta(t, Vector{0.26726, 0.53452, 0.80178}, result)
}

func Test_magnitude_of_normalized_vector(t *testing.T) {
	v := Vector{1, 2, 3}

	result := v.Normalize().Magnitude()

	assert.Equal(t, 1.0, result)
}

func Test_dot_product_vector(t *testing.T) {
	v1 := Vector{1, 2, 3}
	v2 := Vector{2, 3, 4}

	result := v1.Dot(v2)

	assert.Equal(t, 20.0, result)
}

func Test_cross_product_vector(t *testing.T) {
	v1 := Vector{1, 2, 3}
	v2 := Vector{2, 3, 4}

	assert.Equal(t, Vector{-1, 2, -1}, v1.Cross(v2))
	assert.Equal(t, Vector{1, -2, 1}, v2.Cross(v1))
}

func Test_reflecting_vector_approaching_at_45_grad(t *testing.T) {
	v := Vector{1, -1, 0}
	n := Vector{0, 1, 0}

	r := v.Reflect(n)

	assert.Equal(t, Vector{1, 1, 0}, r)
}

func Test_reflecting_vector_off_slanted_surface(t *testing.T) {
	v := Vector{0, -1, 0}
	n := Vector{math.Sqrt2 / 2, math.Sqrt2 / 2, 0}

	r := v.Reflect(n)

	AssertVectorEqualInDelta(t, Vector{1, 0, 0}, r)
}
