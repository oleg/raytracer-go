package multid

import (
	"github.com/stretchr/testify/assert"
	"gray/oned"
	"testing"
)

func Test_canvas(t *testing.T) {
	c := MakeCanvas(10, 20)

	assert.Equal(t, 10, c.Width)
	assert.Equal(t, 20, c.Height)
	for _, row := range c.Pixels {
		for _, c := range row {
			assert.Equal(t, oned.Black, c)
		}
	}
}

func Test_write_pixel(t *testing.T) {
	c := MakeCanvas(10, 20)
	red := oned.Color{1, 0, 0}

	c.Pixels[2][3] = red

	assert.Equal(t, red, c.Pixels[2][3])
}

func Test_canvas_to_png(t *testing.T) {
	c := MakeCanvas(5, 3)
	c.Pixels[0][0] = oned.Color{1, 0, 0}
	c.Pixels[0][1] = oned.Color{1, 0, 0}
	c.Pixels[0][2] = oned.Color{1, 0, 0}
	err := c.ToPNG("canvas_test.png")

	assert.Nil(t, err)
}
