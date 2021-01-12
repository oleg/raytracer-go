package samples

import (
	"github.com/oleg/raytracer-go/figure"
	"github.com/oleg/raytracer-go/geom"
	"os"
	"testing"
)

func Test_ball_3d_sample(t *testing.T) {
	rayOrigin := geom.Point{X: 0, Y: 0, Z: -5}
	wallZ := 10.
	wallSize := 7.0
	canvasPixels := 500 //300 //100
	width := canvasPixels
	height := canvasPixels

	pixelSize := wallSize / float64(canvasPixels)
	half := wallSize / 2.
	canvas := figure.NewCanvas(width, height)
	//white := oned.Color{1, 1, 1}

	transform := geom.IdentityMatrix() //Matrix4x4.Shearing(1, 0, 0, 0, 0, 0) * Matrix4x4.Scaling(0.5, 1, 1)
	//material := figure.Material{Color: oned.Color{0.2, 0.8, 0.3}}
	material := figure.DefaultMaterial()
	material.Color = geom.Color{R: 0.2, G: 0.8, B: 0.3}

	sphere := figure.MakeSphereTM(transform, material)

	lightPosition := geom.Point{X: -10, Y: 10, Z: -10}
	lightColor := geom.White
	light := figure.PointLight{Position: lightPosition, Intensity: lightColor}

	for y := 0; y < canvasPixels; y++ {
		worldY := half - pixelSize*float64(y)
		for x := 0; x < canvasPixels; x++ {

			worldX := -half + pixelSize*float64(x)
			position := geom.Point{X: worldX, Y: worldY, Z: wallZ}

			ray := figure.Ray{
				Origin:    rayOrigin,
				Direction: (position.SubtractPoint(rayOrigin)).Normalize(),
			}

			if ok, h := figure.Intersect(sphere, ray).Hit(); ok {
				point := ray.Position(h.Distance)
				normal := figure.NormalAt(h.Object, point)
				eye := ray.Direction.Negate()

				canvas.Pixels[x][y] = figure.Lighting(h.Object.Material(), h.Object, light, point, eye, normal, false)
			}
		}
	}

	outFile := "ball_3d_sample_test.png"
	canvas.MustToPNG(outFile)

	if AssertFilesEqual(t, "testdata/"+outFile, outFile) {
		_ = os.Remove(outFile)
	}
}
