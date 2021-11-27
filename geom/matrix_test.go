package geom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_create_matrix(t *testing.T) {
	m := NewMatrix(
		`| 1    | 2    | 3    | 4    |  
		 | 5.5  | 6.5  | 7.5  | 8.5  |
		 | 9    | 10   | 11   | 12   |
		 | 13.5 | 14.5 | 15.5 | 16.5 |`)

	assert.Equal(t, 1., m.Data[0][0])
	assert.Equal(t, 4., m.Data[0][3])
	assert.Equal(t, 5.5, m.Data[1][0])
	assert.Equal(t, 7.5, m.Data[1][2])
	assert.Equal(t, 13.5, m.Data[3][0])
	assert.Equal(t, 16.5, m.Data[3][3])
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

	assert.Equal(t, m1, m2)
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

	assert.Equal(t, expected, result)
}

func Test_multiply_matrix_and_point(t *testing.T) {
	m := NewMatrix(
		`| 1 | 2 | 3 | 4 |
	     | 2 | 4 | 4 | 2 |
	     | 8 | 6 | 4 | 1 |
	     | 0 | 0 | 0 | 1 |`)
	p := Point{X: 1, Y: 2, Z: 3}
	result := m.MultiplyPoint(p)

	expected := Point{X: 18, Y: 24, Z: 33}

	assert.Equal(t, expected, result)
}

func Test_multiply_matrix_and_vector(t *testing.T) {
	m := NewMatrix(
		`| 1 | 2 | 3 | 4 |
	     | 2 | 4 | 4 | 2 |
	     | 8 | 6 | 4 | 1 |
	     | 0 | 0 | 0 | 1 |`)
	v := Vector{X: 1, Y: 2, Z: 3}

	result := m.MultiplyVector(v)

	expected := Vector{X: 14, Y: 22, Z: 32}
	assert.Equal(t, expected, result)
}

func Test_multiply_matrix_by_identity_matrix(t *testing.T) {
	m := NewMatrix(
		`| 0 | 1 |  2 |  4 |
		 | 1 | 2 |  4 |  8 |
		 | 2 | 4 |  8 | 16 |
		 | 4 | 8 | 16 | 32 |`)

	r := m.Multiply(IdentityMatrix())

	assert.Equal(t, m, r)
}

func Test_multiply_identity_matrix_by_point(t *testing.T) {
	p := Point{X: 1, Y: 2, Z: 3}

	r := IdentityMatrix().MultiplyPoint(p)

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

	assert.Equal(t, expected, r)
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
	assert.Equal(t, expected, m)
}

func Test_transpose_identity_matrix(t *testing.T) {

	m := IdentityMatrix().Transpose()

	assert.Equal(t, IdentityMatrix().Data, m.Data)
}

func Test_calculate_determinant_of_2x2_matrix(t *testing.T) {
	m := matrix2x2{
		{1, 5},
		{-3, 2}}

	d := m.determinant()

	assert.Equal(t, 17.0, d)
}

func Test_submatrix_of_4x4_matrix_is_3x3_matrix(t *testing.T) {
	m := NewMatrix(
		`| -6 | 1 |  1 | 6 |
		 | -8 | 5 |  8 | 6 |
		 | -1 | 0 |  8 | 2 |
		 | -7 | 1 | -1 | 1 |`)

	r := m.Data.submatrix(2, 1)
	expected := &matrix3x3{
		{-6, 1, 6},
		{-8, 8, 6},
		{-7, -1, 1}}

	assert.Equal(t, expected, r)
}

func Test_submatrix_of_3x3_matrix_is_2x2_matrix(t *testing.T) {
	m := matrix3x3{
		{1, 5, 0},
		{-3, 2, 7},
		{0, 6, -3}}

	r := m.submatrix(0, 2)
	expected := &matrix2x2{
		{-3, 2},
		{-0, 6}}

	assert.Equal(t, expected, r)
}

func Test_calculating_minor_of_3x3_matrix(t *testing.T) {
	m := matrix3x3{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	}

	r := m.minor(1, 0)

	assert.Equal(t, 25.0, r)
}

func Test_Calculating_cofactor_of_3x3_matrix(t *testing.T) {
	m := matrix3x3{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	}

	tests := []struct {
		name     string
		row      int
		column   int
		expected float64
	}{
		{"cofactor 0,0", 0, 0, -12.0},
		{"cofactor 1,0", 1, 0, -25.0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, m.cofactor(test.row, test.column))
		})
	}
}

func Test_calculating_determinant_of_3x3_matrix(t *testing.T) {
	m := matrix3x3{
		{1, 2, 6},
		{-5, 8, -4},
		{2, 6, 4},
	}

	r := m.determinant()

	assert.Equal(t, -196.0, r)
}

func Test_calculating_determinant_of_4x4_matrix(t *testing.T) {
	m := NewMatrix(
		`| -2 | -8 |  3 |  5 |
		 | -3 |  1 |  7 |  3 |
		 |  1 |  2 | -9 |  6 |
		 | -6 |  7 |  7 | -9 |`)

	r := m.Data.determinant()

	assert.Equal(t, -4071.0, r)
}

func Test_invertible_matrix(t *testing.T) {
	m := NewMatrix(
		`|  6 |  4 |  4 |  4 |
		 |  5 |  5 |  7 |  6 |
		 |  4 | -9 |  3 | -7 |
		 |  9 |  1 |  7 | -6 |`)

	r := m.Data.determinant()

	assert.Equal(t, -2120.0, r)
}

func Test_non_invertible_matrix(t *testing.T) {
	m := NewMatrix(
		`| -4 |  2 | -2 | -3 |
		 |  9 |  6 |  2 |  6 |
		 |  0 | -5 |  1 | -5 |
		 |  0 |  0 |  0 |  0 |`)

	r := m.Data.determinant()

	assert.Equal(t, 0.0, r)
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
