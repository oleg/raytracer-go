package scene

import (
	"github.com/oleg/raytracer-go/geom"
	"image"
	"image/png"
	"os"
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

func (c *Canvas) MustToPNG(filename string) {
	err := c.ToPNG(filename)
	if err != nil {
		panic(err)
	}
}

func (c *Canvas) ToPNG(filename string) error {
	fo, err := os.Create(filename)
	if err != nil {
		return err
	}
	img := c.newImage()
	err = png.Encode(fo, img)
	if err != nil {
		return err
	}
	if err := fo.Close(); err != nil {
		return err
	}
	return nil
}

func (c *Canvas) newImage() image.Image {
	img := image.NewRGBA(image.Rect(0, 0, c.Width, c.Height))
	for i, p := range c.Pixels {
		for j, px := range p {
			img.Set(i, j, px)
		}
	}
	return img
}
