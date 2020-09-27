package samples

import (
	"gray/figure"
	"gray/multid"
	"gray/oned"
	"testing"
)

//todo do not use example in the name
func Test_ball_sample(t *testing.T) {
	rayOrigin := oned.Point{0, 0, -5}
	wallSize := 7.
	canvasPixels := 100
	pixelSize := wallSize / float64(canvasPixels)
	half := wallSize / 2.
	canvas := multid.MakeCanvas(canvasPixels, canvasPixels)
	red := oned.Color{1, 0, 0}
	transform := multid.Shearing(1, 0, 0, 0, 0, 0).Multiply(multid.Scaling(0.5, 1, 1))
	sphere := figure.MakeSphereT(transform)

	for y := 0; y < canvasPixels; y++ {
		worldY := half - pixelSize*float64(y)
		for x := 0; x < canvasPixels; x++ {
			worldX := -half + pixelSize*float64(x)
			position := oned.Point{worldX, worldY, 10}
			ray := figure.Ray{rayOrigin, position.SubtractPoint(rayOrigin).Normalize()}
			if hit, _ := figure.Intersect(sphere, ray).Hit(); hit {
				canvas.Pixels[x][y] = red
			}
		}
	}

	canvas.MustToPNG("ball_sample_test.png")
}
