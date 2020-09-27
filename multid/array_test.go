package multid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_multiply_arrays_empty(t *testing.T) {
	var a [][]float64 = nil
	var b [][]float64 = nil

	c := multiply(a, b)

	assert.Nil(t, c)
}

func Test_multiply_arrays_single(t *testing.T) {
	a := [][]float64{{2}}
	b := [][]float64{{3}}

	c := multiply(a, b)

	expected := [][]float64{{6}}

	assert.Equal(t, expected, c)
}

func Test_multiply_arrays_two_three_on_three_five(t *testing.T) {
	a := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}
	b := [][]float64{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 0},
		{3, 4, 5, 6, 7},
	}

	c := multiply(a, b)

	expected := [][]float64{
		{1*1 + 2*6 + 3*3, 1*2 + 2*7 + 3*4, 1*3 + 2*8 + 3*5, 1*4 + 2*9 + 3*6, 1*5 + 2*0 + 3*7},
		{4*1 + 5*6 + 6*3, 4*2 + 5*7 + 6*4, 4*3 + 5*8 + 6*5, 4*4 + 5*9 + 6*6, 4*5 + 5*0 + 6*7},
	}
	assert.Equal(t, expected, c)
}
