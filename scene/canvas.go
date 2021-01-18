package scene

import (
	"github.com/oleg/raytracer-go/geom"
	"image"
	"image/color"
)

//todo move encoding to separate package?
//todo store color.RGBA instead of geom.Color?
type Canvas struct {
	Width, Height int
	Pixels        [][]geom.Color
}

func NewCanvas(width, height int) *Canvas {
	pixels := make([][]geom.Color, width)
	for i := range pixels {
		pixels[i] = make([]geom.Color, height)
	}
	return &Canvas{width, height, pixels}
}

func (c *Canvas) ColorModel() color.Model {
	return color.RGBAModel
}

func (c *Canvas) Bounds() image.Rectangle {
	return image.Rect(0, 0, c.Width, c.Height)
}

func (c *Canvas) At(x, y int) color.Color {
	return c.Pixels[x][y]
}
