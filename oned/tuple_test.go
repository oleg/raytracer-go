package oned

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_tuple_properties(t *testing.T) {
	result := Tuple{1.1, 2.2, 3.3}

	assert.EqualValues(t, 1.1, result.X)
	assert.EqualValues(t, 2.2, result.Y)
	assert.EqualValues(t, 3.3, result.Z)
}

func Test_add_tuples(t *testing.T) {
	t1 := Tuple{3, -2, 5}
	t2 := Tuple{-2, 3, 1}

	result := t1.add(t2)

	assert.Equal(t, Tuple{1, 1, 6}, result)
}

func Test_subtract_tuples(t *testing.T) {
	t1 := Tuple{3, -2, 5}
	t2 := Tuple{-2, 3, 1}

	result := t1.subtract(t2)

	assert.Equal(t, Tuple{5, -5, 4}, result)
}

func Test_negate_tuples(t *testing.T) {
	t1 := Tuple{-2, 3, -1}

	result := t1.negate()

	assert.Equal(t, Tuple{2, -3, 1}, result)
}

func Test_multiply_scalar_tuples(t *testing.T) {
	t1 := Tuple{1, -2, 3}

	result := t1.multiplyScalar(3.5)

	assert.Equal(t, Tuple{3.5, -7, 10.5}, result)
}

func Test_multiply_tuple_by_fraction(t *testing.T) {
	t1 := Tuple{1, -2, 3}

	result := t1.multiplyScalar(0.5)

	assert.Equal(t, Tuple{0.5, -1, 1.5}, result)
}

func Test_divide_tuple(t *testing.T) {
	t1 := Tuple{1, -2, 3}

	result := t1.divideScalar(2)

	assert.Equal(t, Tuple{0.5, -1, 1.5}, result)
}

func Test_multiply_tuple(t *testing.T) {
	t1 := Tuple{1, -2, 3}
	t2 := Tuple{2, 3, 4}

	result := t1.hadamard(t2)

	assert.Equal(t, Tuple{2, -6, 12}, result)
}

func Test_dot_product_tuple(t *testing.T) {
	v1 := Tuple{1, 2, 3}
	v2 := Tuple{2, 3, 4}

	result := v1.dot(v2)

	assert.Equal(t, 20.0, result)
}

func Test_cross_product_tuple(t *testing.T) {
	v1 := Tuple{1, 2, 3}
	v2 := Tuple{2, 3, 4}

	assert.Equal(t, Tuple{-1, 2, -1}, v1.cross(v2))
	assert.Equal(t, Tuple{1, -2, 1}, v2.cross(v1))
}
