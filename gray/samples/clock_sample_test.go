package samples

import (
	"gray/multid"
	"gray/oned"
	"math"
	"testing"
)

//todo add assert
func Test_clock_example_test(t *testing.T) {
	canvas := multid.MakeCanvas(500, 500)
	radius := float64(canvas.Width * 3 / 8)

	rotationY := multid.RotationY(math.Pi / 6)

	points := make([]oned.Point, 12, 12)
	points[0] = oned.Point{0, 0, 1}
	for i := 1; i < 12; i++ {
		points[i] = rotationY.MultiplyPoint(points[i-1])
	}

	white := oned.Color{1, 1, 1}
	for _, p := range points {
		x := int(p.X*radius + 250)
		y := int(p.Z*radius + 250)
		canvas.Pixels[x][y] = white
	}

	canvas.MustToPNG("clock_sample_test.png")
}
