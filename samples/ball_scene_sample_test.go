package samples

import (
	"github.com/oleg/raytracer-go/figure"
	"github.com/oleg/raytracer-go/multid"
	"github.com/oleg/raytracer-go/oned"
	"math"
	"os"
	"testing"
)

func Test_ball_scene_sample(t *testing.T) {
	floorMaterial := figure.MakeMaterialBuilder().
		SetColor(oned.Color{R: 1, G: 0.9, B: 0.9}).
		SetSpecular(0).
		Build()

	floor := figure.MakeSphereTM(
		multid.Scaling(10, 0.01, 10),
		floorMaterial)

	leftWall := figure.MakeSphereTM(
		multid.Translation(0, 0, 5).
			Multiply(multid.RotationY(-math.Pi/4)).
			Multiply(multid.RotationX(math.Pi/2)).
			Multiply(multid.Scaling(10, 0.01, 10)),
		floorMaterial)

	rightWall := figure.MakeSphereTM(
		multid.Translation(0, 0, 5).
			Multiply(multid.RotationY(math.Pi/4)).
			Multiply(multid.RotationX(math.Pi/2)).
			Multiply(multid.Scaling(10, 0.01, 10)),
		floorMaterial)

	middle := figure.MakeSphereTM(
		multid.Translation(-0.5, 1, 0.5),
		figure.MakeMaterialBuilder().
			SetColor(oned.Color{R: 0.1, G: 1, B: 0.5}).
			SetDiffuse(0.7).
			SetSpecular(0.3).Build())

	right := figure.MakeSphereTM(
		multid.Translation(1.5, 0.5, -0.5).
			Multiply(multid.Scaling(0.5, 0.5, 0.5)),
		figure.MakeMaterialBuilder().
			SetColor(oned.Color{R: 0.5, G: 1, B: 0.1}).
			SetDiffuse(0.7).
			SetSpecular(0.3).Build())

	left := figure.MakeSphereTM(
		multid.Translation(-1.5, 0.33, -0.75).
			Multiply(multid.Scaling(0.33, 0.33, 0.33)),
		figure.MakeMaterialBuilder().
			SetColor(oned.Color{R: 1, G: 0.8, B: 0.1}).
			SetDiffuse(0.7).
			SetSpecular(0.3).Build())

	light := figure.PointLight{
		Position:  oned.Point{X: -10, Y: 10, Z: -10},
		Intensity: oned.White,
	}
	world := figure.World{
		Light: light,
		Objects: []figure.Shape{
			floor, leftWall, rightWall, middle, right, left,
		},
	}
	camera := figure.NewCamera(500, 250, math.Pi/3,
		figure.ViewTransform(
			oned.Point{X: 0, Y: 1.5, Z: -5},
			oned.Point{X: 0, Y: 1, Z: 0},
			oned.Vector{X: 0, Y: 1, Z: 0}))

	canvas := camera.Render(world)

	outFile := "ball_scene_sample_test.png"
	canvas.MustToPNG(outFile)

	if AssertFilesEqual(t, "testdata/"+outFile, outFile) {
		_ = os.Remove(outFile)
	}
}
