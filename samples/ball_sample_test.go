package samples

import (
	"github.com/oleg/raytracer-go/figure"
	"github.com/oleg/raytracer-go/multid"
	"github.com/oleg/raytracer-go/oned"
	"os"
	"testing"
)

//todo do not use example in the name
func Test_ball_sample(t *testing.T) {
	rayOrigin := oned.Point{X: 0, Y: 0, Z: -5}
	wallSize := 7.
	canvasPixels := 100
	pixelSize := wallSize / float64(canvasPixels)
	half := wallSize / 2.
	canvas := multid.NewCanvas(canvasPixels, canvasPixels)
	red := oned.Color{R: 1, G: 0, B: 0}
	transform := multid.Shearing(1, 0, 0, 0, 0, 0).Multiply(multid.Scaling(0.5, 1, 1))
	sphere := figure.MakeSphereT(transform)

	for y := 0; y < canvasPixels; y++ {
		worldY := half - pixelSize*float64(y)
		for x := 0; x < canvasPixels; x++ {
			worldX := -half + pixelSize*float64(x)
			position := oned.Point{X: worldX, Y: worldY, Z: 10}
			ray := figure.Ray{
				Origin: rayOrigin,
				Direction: position.SubtractPoint(rayOrigin).Normalize(),
			}
			if hit, _ := figure.Intersect(sphere, ray).Hit(); hit {
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
