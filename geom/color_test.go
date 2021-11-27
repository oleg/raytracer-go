package geom

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_color(t *testing.T) {
	c := Color{-0.5, 0.4, 1.7}

	assert.Equal(t, -0.5, c.R)
	assert.Equal(t, 0.4, c.G)
	assert.Equal(t, 1.7, c.B)
}

func Test_adding_colors(t *testing.T) {
	c1 := Color{0.9, 0.6, 0.75}
	c2 := Color{0.7, 0.1, 0.25}

	result := c1.Add(c2)

	AssertColorEqualInDelta(t, Color{1.6, 0.7, 1.0}, result)
}

func Test_subtracting_colors(t *testing.T) {
	c1 := Color{0.9, 0.6, 0.75}
	c2 := Color{0.7, 0.1, 0.25}

	result := c1.Subtract(c2)

	AssertColorEqualInDelta(t, Color{0.2, 0.5, 0.5}, result)
}

func Test_multiplying_by_scalar(t *testing.T) {
	c1 := Color{0.2, 0.3, 0.4}

	result := c1.MultiplyByScalar(2)

	AssertColorEqualInDelta(t, Color{0.4, 0.6, 0.8}, result)
}

func Test_multiply_colors(t *testing.T) {
	c1 := Color{1, 0.2, 0.4}
	c2 := Color{0.9, 1, 0.1}

	result := c1.Multiply(c2)

	AssertColorEqualInDelta(t, Color{0.9, 0.2, 0.04}, result)
}

func Test_RGBA(t *testing.T) {
	var c color.Color = Color{1, 0.2, 0.8}

	r, g, b, a := c.RGBA()

	assert.Equal(t, uint32(0xffff), r)
	assert.Equal(t, uint32(0x3333), g)
	assert.Equal(t, uint32(0xcccc), b)
	assert.Equal(t, uint32(0xffff), a)
}
