package samples

import (
	"gray/figure"
	"gray/multid"
	"gray/oned"
	"testing"
)

func Test_ball_3d_sample(t *testing.T) {
	rayOrigin := oned.Point{0, 0, -5}
	wallZ := 10.
	wallSize := 7.0
	canvasPixels := 500 //300 //100
	width := canvasPixels
	height := canvasPixels

	pixelSize := wallSize / float64(canvasPixels)
	half := wallSize / 2.
	canvas := multid.MakeCanvas(width, height)
	//white := oned.Color{1, 1, 1}

	transform := multid.IdentityMatrixF() //Matrix4x4.Shearing(1, 0, 0, 0, 0, 0) * Matrix4x4.Scaling(0.5, 1, 1)
	//material := figure.Material{Color: oned.Color{0.2, 0.8, 0.3}}
	material := figure.DefaultMaterial()
	material.Color = oned.Color{0.2, 0.8, 0.3}

	sphere := figure.MakeSphereTM(transform, material)

	lightPosition := oned.Point{-10, 10, -10}
	lightColor := oned.White
	light := figure.PointLight{lightPosition, lightColor}

	for y := 0; y < canvasPixels; y++ {
		worldY := half - pixelSize*float64(y)
		for x := 0; x < canvasPixels; x++ {

			worldX := -half + pixelSize*float64(x)
			position := oned.Point{worldX, worldY, wallZ}

			ray := figure.Ray{rayOrigin, (position.SubtractPoint(rayOrigin)).Normalize()}

			if ok, h := figure.Intersect(sphere, ray).Hit(); ok {
				point := ray.Position(h.Distance)
				normal := figure.NormalAt(h.Object, point)
				eye := ray.Direction.Negate()

				canvas.Pixels[x][y] = figure.Lighting(h.Object.Material(), h.Object, light, point, eye, normal, false)
			}
		}
	}

	canvas.MustToPNG("ball_3d_sample_test.png")
}
