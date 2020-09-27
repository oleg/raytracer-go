package multid

import (
	"gray/oned"
	"image"
	"image/color"
	"image/png"
	"os"
)

/*
todo:
store color.RGBA instead of gray.Color
*/
type Canvas struct {
	Width, Height int
	Pixels        [][]oned.Color
}

func MakeCanvas(width, height int) Canvas {
	pixels := make([][]oned.Color, width)
	for i := range pixels {
		pixels[i] = make([]oned.Color, height)
	}
	return Canvas{width, height, pixels}
}

func (c Canvas) ToPNG(filename string) error {
	fo, err := os.Create(filename)
	if err != nil {
		return err
	}

	img := image.NewRGBA(image.Rect(0, 0, c.Width, c.Height))
	for i, p := range c.Pixels {
		for j, px := range p {

			img.Set(i, j, color.RGBA{ //todo (Height-j)?
				R: uint8(clamp(px.R()) * 255),
				G: uint8(clamp(px.G()) * 255),
				B: uint8(clamp(px.B()) * 255),
				A: 255})
		}
	}
	err = png.Encode(fo, img)
	if err != nil {
		return err
	}
	if err := fo.Close(); err != nil {
		return err
	}
	return nil
}

//todo refactor
func clamp(r float64) float64 {
	if r < 0 {
		return 0
	}
	if r > 1 {
		return 1
	}
	return r
}

func (c Canvas) MustToPNG(filename string) {
	err := c.ToPNG(filename)
	if err != nil {
		panic(err)
	}
}
