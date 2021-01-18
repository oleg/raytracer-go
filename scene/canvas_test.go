package scene

import (
	"bytes"
	"github.com/oleg/raytracer-go/geom"
	"github.com/stretchr/testify/assert"
	"image"
	"image/png"
	"testing"
)

func Test_canvas(t *testing.T) {
	c := NewCanvas(10, 20)

	assert.Equal(t, 10, c.Width)
	assert.Equal(t, 20, c.Height)
	for _, row := range c.Pixels {
		for _, c := range row {
			assert.Equal(t, geom.Black, c)
		}
	}
}

func Test_write_pixel(t *testing.T) {
	c := NewCanvas(10, 20)
	red := geom.Color{R: 1, G: 0, B: 0}

	c.Pixels[2][3] = red

	assert.Equal(t, red, c.Pixels[2][3])
}

func Test_canvas_to_png(t *testing.T) {
	c := NewCanvas(5, 3)
	c.Pixels[0][0] = geom.Color{R: 1, G: 0, B: 0}
	c.Pixels[0][1] = geom.Color{R: 1, G: 0, B: 0}
	c.Pixels[0][2] = geom.Color{R: 1, G: 0, B: 0}
	err := png.Encode(new(bytes.Buffer), c)

	assert.Nil(t, err)
}

func Test_canvas_implements_image_interface(t *testing.T) {
	var _ image.Image = NewCanvas(10, 10)
}
