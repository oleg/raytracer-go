package samples

import (
	"bytes"
	"github.com/oleg/raytracer-go/geom"
	"github.com/oleg/raytracer-go/scene"
	"github.com/stretchr/testify/assert"
	"image/png"
	"math"
	"testing"
)

func Test_clock_example_test(t *testing.T) {
	canvas := scene.NewCanvas(500, 500)
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

	b := new(bytes.Buffer)
	err := png.Encode(b, canvas)
	assert.NoError(t, err)
	AssertBytesAreEqual(t, "testdata/clock_sample_test.png", b.Bytes())
}
