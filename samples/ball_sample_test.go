package samples

import (
	"bytes"
	"github.com/oleg/raytracer-go/geom"
	"github.com/oleg/raytracer-go/physic"
	"github.com/oleg/raytracer-go/scene"
	"github.com/oleg/raytracer-go/shapes"
	"github.com/stretchr/testify/assert"
	"image/png"
	"testing"
)

func Test_ball_sample(t *testing.T) {
	rayOrigin := geom.Point{X: 0, Y: 0, Z: -5}
	wallSize := 7.
	canvasPixels := 100
	pixelSize := wallSize / float64(canvasPixels)
	half := wallSize / 2.
	canvas := scene.NewCanvas(canvasPixels, canvasPixels)
	red := geom.Color{R: 1, G: 0, B: 0}
	transform := geom.Shearing(1, 0, 0, 0, 0, 0).Multiply(geom.Scaling(0.5, 1, 1))
	sphere := shapes.NewSphere(transform, physic.DefaultMaterial())

	for y := 0; y < canvasPixels; y++ {
		worldY := half - pixelSize*float64(y)
		for x := 0; x < canvasPixels; x++ {
			worldX := -half + pixelSize*float64(x)
			position := geom.Point{X: worldX, Y: worldY, Z: 10}
			ray := shapes.Ray{
				Origin:    rayOrigin,
				Direction: position.SubtractPoint(rayOrigin).Normalize(),
			}
			if hit, _ := scene.Hit(sphere.Intersect(ray.ToLocal(sphere))); hit {
				canvas.Pixels[x][y] = red
			}
		}
	}

	b := new(bytes.Buffer)
	err := png.Encode(b, canvas)
	assert.NoError(t, err)
	AssertBytesAreEqual(t, "testdata/ball_sample_test.png", b.Bytes())

}
