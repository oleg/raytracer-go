package multid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_transpose_identity_matrix(t *testing.T) {

	m := IdentityMatrix.Transpose()

	assert.Equal(t, IdentityMatrix, m)
}

func Test_calculate_determinant_of_2x2_matrix(t *testing.T) {
	m := [2][2]float64{
		{1, 5},
		{-3, 2}}

	d := determinant2x2(&m)

	assert.Equal(t, 17.0, d)
}

func Test_submatrix_of_4x4_matrix_is_3x3_matrix(t *testing.T) {
	m := NewMatrix(
		`| -6 | 1 |  1 | 6 |
		 | -8 | 5 |  8 | 6 |
		 | -1 | 0 |  8 | 2 |
		 | -7 | 1 | -1 | 1 |`)

	r := submatrix4x4(m, 2, 1)
	expected := [3][3]float64{
		{-6, 1, 6},
		{-8, 8, 6},
		{-7, -1, 1}}

	assert.Equal(t, expected, r)
}

func Test_submatrix_of_3x3_matrix_is_2x2_matrix(t *testing.T) {
	m := [3][3]float64{
		{1, 5, 0},
		{-3, 2, 7},
		{0, 6, -3}}

	r := submatrix3x3(&m, 0, 2)
	expected := [2][2]float64{
		{-3, 2},
		{-0, 6}}

	assert.Equal(t, expected, r)
}

func Test_calculating_minor_of_3x3_matrix(t *testing.T) {
	m3x3 := [3][3]float64{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	}

	r := minor3x3(&m3x3, 1, 0)

	assert.Equal(t, 25.0, r)
}

func Test_Calculating_cofactor_of_3x3_matrix(t *testing.T) {
	m3x3 := [3][3]float64{
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
			assert.Equal(t, test.expected, cofactor3x3(&m3x3, test.row, test.column))
		})
	}
}

func Test_calculating_determinant_of_3x3_matrix(t *testing.T) {
	m3x3 := [3][3]float64{
		{1, 2, 6},
		{-5, 8, -4},
		{2, 6, 4},
	}

	r := determinant3x3(&m3x3)

	assert.Equal(t, -196.0, r)
}

func Test_calculating_determinant_of_4x4_matrix(t *testing.T) {
	m := NewMatrix(
		`| -2 | -8 |  3 |  5 |
		 | -3 |  1 |  7 |  3 |
		 |  1 |  2 | -9 |  6 |
		 | -6 |  7 |  7 | -9 |`)

	r := determinant4x4(m)

	assert.Equal(t, -4071.0, r)
}

func Test_invertible_matrix(t *testing.T) {
	m := NewMatrix(
		`|  6 |  4 |  4 |  4 |
		 |  5 |  5 |  7 |  6 |
		 |  4 | -9 |  3 | -7 |
		 |  9 |  1 |  7 | -6 |`)

	r := determinant4x4(m)

	assert.Equal(t, -2120.0, r)
}

func Test_non_invertible_matrix(t *testing.T) {
	m := NewMatrix(
		`| -4 |  2 | -2 | -3 |
		 |  9 |  6 |  2 |  6 |
		 |  0 | -5 |  1 | -5 |
		 |  0 |  0 |  0 |  0 |`)

	r := determinant4x4(m)

	assert.Equal(t, 0.0, r)
}
