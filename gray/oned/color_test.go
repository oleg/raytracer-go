package oned

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_color(t *testing.T) {
	c := NewColor(-0.5, 0.4, 1.7)

	assert.Equal(t, -0.5, c.R())
	assert.Equal(t, 0.4, c.G())
	assert.Equal(t, 1.7, c.B())
}

func Test_adding_colors(t *testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)

	result := c1.Add(c2)

	AssertColorEqualInDelta(t, NewColor(1.6, 0.7, 1.0), result)
}

func Test_subtracting_colors(t *testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)

	result := c1.Subtract(c2)

	AssertColorEqualInDelta(t, NewColor(0.2, 0.5, 0.5), result)
}

func Test_multiplying_by_scalar(t *testing.T) {
	c1 := NewColor(0.2, 0.3, 0.4)

	result := c1.MultiplyByScalar(2)

	AssertColorEqualInDelta(t, NewColor(0.4, 0.6, 0.8), result)
}

func Test_multiply_colors(t *testing.T) {
	c1 := NewColor(1, 0.2, 0.4)
	c2 := NewColor(0.9, 1, 0.1)

	result := c1.Multiply(c2)

	AssertColorEqualInDelta(t, NewColor(0.9, 0.2, 0.04), result)
}
