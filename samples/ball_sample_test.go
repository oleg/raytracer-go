package samples

import (
	"github.com/oleg/raytracer-go/asdf"
	"github.com/oleg/raytracer-go/ddddf"
	"github.com/oleg/raytracer-go/figure"
	"github.com/oleg/raytracer-go/geom"
	"os"
	"testing"
)

//todo do not use example in the name
func Test_ball_sample(t *testing.T) {
	rayOrigin := geom.Point{X: 0, Y: 0, Z: -5}
	wallSize := 7.
	canvasPixels := 100
	pixelSize := wallSize / float64(canvasPixels)
	half := wallSize / 2.
	canvas := figure.NewCanvas(canvasPixels, canvasPixels)
	red := geom.Color{R: 1, G: 0, B: 0}
	transform := geom.Shearing(1, 0, 0, 0, 0, 0).Multiply(geom.Scaling(0.5, 1, 1))
	sphere := ddddf.NewSphere(transform, asdf.DefaultMaterial())

	for y := 0; y < canvasPixels; y++ {
		worldY := half - pixelSize*float64(y)
		for x := 0; x < canvasPixels; x++ {
			worldX := -half + pixelSize*float64(x)
			position := geom.Point{X: worldX, Y: worldY, Z: 10}
			ray := ddddf.Ray{
				Origin:    rayOrigin,
				Direction: position.SubtractPoint(rayOrigin).Normalize(),
			}
			if hit, _ := figure.Hit(ddddf.Intersect(sphere, ray)); hit {
				canvas.Pixels[x][y] = red
			}
		}
	}

	outFile := "ball_sample_test.png"
	canvas.MustToPNG(outFile)

	if AssertFilesEqual(t, "testdata/"+outFile, outFile) {
		_ = os.Remove(outFile)
	}
}
