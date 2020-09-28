package multid

import (
	"github.com/oleg/graytracer/oned"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_create_matrix(t *testing.T) {
	m := NewMatrix(
		`| 1    | 2    | 3    | 4    |  
		 | 5.5  | 6.5  | 7.5  | 8.5  |
		 | 9    | 10   | 11   | 12   |
		 | 13.5 | 14.5 | 15.5 | 16.5 |`)

	assert.Equal(t, 1., m.data[0][0])
	assert.Equal(t, 4., m.data[0][3])
	assert.Equal(t, 5.5, m.data[1][0])
	assert.Equal(t, 7.5, m.data[1][2])
	assert.Equal(t, 13.5, m.data[3][0])
	assert.Equal(t, 16.5, m.data[3][3])
}

func Test_matrices_equal(t *testing.T) {
	m1 := NewMatrix(
		`| 1 | 2 | 3 | 4 |
 		 | 5 | 6 | 7 | 8 |
	     | 9 | 8 | 7 | 6 |
	     | 5 | 4 | 3 | 2 |`)
	m2 := NewMatrix(
		`| 1 | 2 | 3 | 4 |
	     | 5 | 6 | 7 | 8 |
	     | 9 | 8 | 7 | 6 |
	     | 5 | 4 | 3 | 2 |`)

	AssertMatrixEqualInDelta(t, m1, m2)
}

func Test_matrices_not_equal(t *testing.T) {
	m1 := NewMatrix(
		`| 1 | 2 | 3 | 4 |
	     | 5 | 6 | 7 | 8 |
	     | 9 | 8 | 7 | 6 |
	     | 5 | 4 | 3 | 2 |`)
	m2 := NewMatrix(
		` | 2 | 3 | 4 | 5 |
	     | 6 | 7 | 8 | 9 |
	     | 8 | 7 | 6 | 5 |
	     | 4 | 3 | 2 | 1 |`)

	assert.NotEqual(t, m1, m2)
}

func Test_multiply_matrices(t *testing.T) {
	m1 := NewMatrix(
		`| 1 | 2 | 3 | 4 |
	     | 5 | 6 | 7 | 8 |
	     | 9 | 8 | 7 | 6 |
	     | 5 | 4 | 3 | 2 |`)
	m2 := NewMatrix(
		`| -2 | 1 | 2 |  3 |
	     |  3 | 2 | 1 | -1 |
	     |  4 | 3 | 6 |  5 |
	     |  1 | 2 | 7 |  8 |`)

	result := m1.Multiply(m2)

	expected := NewMatrix(
		`| 20|  22 |  50 |  48 |
	     | 44|  54 | 114 | 108 |
	     | 40|  58 | 110 | 102 |
	     | 16|  26 |  46 |  42 |`)

	AssertMatrixEqualInDelta(t, expected, result)
}

func Test_multiply_matrix_and_point(t *testing.T) {
	m := NewMatrix(
		`| 1 | 2 | 3 | 4 |
	     | 2 | 4 | 4 | 2 |
	     | 8 | 6 | 4 | 1 |
	     | 0 | 0 | 0 | 1 |`)
	p := oned.Point{1, 2, 3}
	result := m.MultiplyPoint(p)

	expected := oned.Point{18, 24, 33}

	assert.Equal(t, expected, result)
}

func Test_multiply_matrix_and_vector(t *testing.T) {
	m := NewMatrix(
		`| 1 | 2 | 3 | 4 |
	     | 2 | 4 | 4 | 2 |
	     | 8 | 6 | 4 | 1 |
	     | 0 | 0 | 0 | 1 |`)
	v := oned.Vector{1, 2, 3}

	result := m.MultiplyVector(v)

	expected := oned.Vector{14, 22, 32}
	assert.Equal(t, expected, result)
}

func Test_multiply_matrix_by_identity_matrix(t *testing.T) {
	m := NewMatrix(
		`| 0 | 1 |  2 |  4 |
		 | 1 | 2 |  4 |  8 |
		 | 2 | 4 |  8 | 16 |
		 | 4 | 8 | 16 | 32 |`)

	r := m.Multiply(IdentityMatrix)

	AssertMatrixEqualInDelta(t, m, r)
}

func Test_multiply_identity_matrix_by_point(t *testing.T) {
	p := oned.Point{1, 2, 3}

	r := IdentityMatrix.MultiplyPoint(p)

	assert.Equal(t, p, r)
}

func Test_transpose_matrix(t *testing.T) {
	m := NewMatrix(
		`| 0 | 9 | 3 | 0 |
		 | 9 | 8 | 0 | 8 |
		 | 1 | 8 | 5 | 3 |
		 | 0 | 0 | 5 | 8 |`)

	r := m.Transpose()
	expected := NewMatrix(
		`| 0 | 9 | 1 | 0 |
		 | 9 | 8 | 8 | 0 |
		 | 3 | 0 | 5 | 5 |
		 | 0 | 8 | 3 | 8 |`)

	AssertMatrixEqualInDelta(t, expected, r)
}

func Test_transpose_does_not_change_original(t *testing.T) {
	data :=
		`| 0 | 9 | 3 | 0 |
		 | 9 | 8 | 0 | 8 |
		 | 1 | 8 | 5 | 3 |
		 | 0 | 0 | 5 | 8 |`
	m := NewMatrix(data)

	m.Transpose()

	expected := NewMatrix(data)
	AssertMatrixEqualInDelta(t, expected, m)
}

func Test_calculating_inverse_of_matrix(t *testing.T) {
	m := NewMatrix(
		`| -5 |  2 |  6 | -8 |
		 |  1 | -5 |  1 |  8 |
		 |  7 |  7 | -6 | -7 |
		 |  1 | -3 |  7 |  4 |`)

	r := m.Inverse()

	expected := NewMatrix(
		`|  0.21805 |  0.45113 |  0.24060 | -0.04511 |
		 | -0.80827 | -1.45677 | -0.44361 |  0.52068 |
		 | -0.07895 | -0.22368 | -0.05263 |  0.19737 |
		 | -0.52256 | -0.81391 | -0.30075 |  0.30639 |`)
	AssertMatrixEqualInDelta(t, expected, r)
}

func Test_calculating_inverse_of_matrix_case_2(t *testing.T) {
	m := NewMatrix(
		`|  8 | -5 |  9 |  2 |
		 |  7 |  5 |  6 |  1 |
		 | -6 |  0 |  9 |  6 |
		 | -3 |  0 | -9 | -4 |`)

	r := m.Inverse()

	expected := NewMatrix(
		`| -0.15385 | -0.15385 | -0.28205 | -0.53846 |
		 | -0.07692 |  0.12308 |  0.02564 |  0.03077 |
		 |  0.35897 |  0.35897 |  0.43590 |  0.92308 |
		 | -0.69231 | -0.69231 | -0.76923 | -1.92308 |`)
	AssertMatrixEqualInDelta(t, expected, r)
}

func Test_calculating_inverse_of_matrix_case_3(t *testing.T) {
	m := NewMatrix(
		`|  9 |  3 |  0 |  9 |
		 | -5 | -2 | -6 | -3 |
		 | -4 |  9 |  6 |  4 |
		 | -7 |  6 |  6 |  2 |`)

	r := m.Inverse()

	expected := NewMatrix(
		`| -0.04074 | -0.07778 |  0.14444 | -0.22222 |
		 | -0.07778 |  0.03333 |  0.36667 | -0.33333 |
		 | -0.02901 | -0.14630 | -0.10926 |  0.12963 |
		 |  0.17778 |  0.06667 | -0.26667 |  0.33333 |`)
	AssertMatrixEqualInDelta(t, expected, r)
}

func Test_multiplying_product_by_its_inverse(t *testing.T) {
	ma := NewMatrix(
		`|  3 | -9 |  7 |  3 |
		 |  3 | -8 |  2 | -9 |
		 | -4 |  4 |  4 |  1 |
		 | -6 |  5 | -1 |  1 |`)
	mb := NewMatrix(
		`|  8 |  2 |  2 |  2 |
		 |  3 | -1 |  7 |  0 |
		 |  7 |  0 |  5 |  4 |
		 |  6 | -2 |  0 |  5 |`)

	r := ma.Multiply(mb).Multiply(mb.Inverse())

	AssertMatrixEqualInDelta(t, ma, r)
}
