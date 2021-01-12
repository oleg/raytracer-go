package samples

import (
	"github.com/oleg/raytracer-go/figure"
	"github.com/oleg/raytracer-go/geom"
	"math"
	"os"
	"testing"
)

func Test_clock_example_test(t *testing.T) {
	canvas := figure.NewCanvas(500, 500)
	radius := float64(canvas.Width * 3 / 8)

	rotationY := geom.RotationY(math.Pi / 6)

	points := make([]geom.Point, 12, 12)
	points[0] = geom.Point{X: 0, Y: 0, Z: 1}
	for i := 1; i < 12; i++ {
		points[i] = rotationY.MultiplyPoint(points[i-1])
	}

	white := geom.Color{R: 1, G: 1, B: 1}
	for _, p := range points {
		x := int(p.X*radius + 250)
		y := int(p.Z*radius + 250)
		canvas.Pixels[x][y] = white
	}

	outFile := "clock_sample_test.png"
	canvas.MustToPNG(outFile)

	if AssertFilesEqual(t, "testdata/"+outFile, outFile) {
		_ = os.Remove(outFile)
	}
}
