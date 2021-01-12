package samples

import (
	"github.com/oleg/raytracer-go/mat"
	"github.com/oleg/raytracer-go/ddddf"
	"github.com/oleg/raytracer-go/figure"
	"github.com/oleg/raytracer-go/geom"
	"math"
	"os"
	"testing"
)

func Test_refraction_sample(t *testing.T) {
	floor := ddddf.NewPlane(
		geom.IdentityMatrix(),
		mat.MakeMaterialBuilder().
			SetReflective(0.7).
			SetTransparency(0.2).
			SetRefractiveIndex(1.3).
			SetPattern(mat.MakeCheckersPatternT(
				geom.Black,
				geom.White,
				geom.IdentityMatrix())).
			Build())

	back := ddddf.NewPlane(
		geom.Translation(0, 0, 4).
			Multiply(geom.RotationX(-math.Pi/2)),
		mat.MakeMaterialBuilder().
			SetReflective(0.3).
			SetTransparency(0.1).
			SetRefractiveIndex(2).
			SetPattern(mat.MakeCheckersPatternT(
				geom.Black,
				geom.White,
				geom.IdentityMatrix())).
			Build())
	left := ddddf.NewSphere(
		geom.Translation(-2.4, 1, 0.2),
		mat.MakeMaterialBuilder().
			//SetSpecular(1).
			SetTransparency(0.3).
			SetReflective(0.3).
			SetRefractiveIndex(1).
			SetAmbient(0.2).
			SetColor(geom.White).
			Build())

	middle := ddddf.NewSphere(
		geom.Translation(-0.1, 1, 0.2),
		mat.MakeMaterialBuilder().
			SetTransparency(0.5).
			SetReflective(0.3).
			SetRefractiveIndex(1.2).
			SetColor(geom.Color{R: 0.4, G: 0, B: 0}).
			Build())

	right := ddddf.NewSphere(
		geom.Translation(2.2, 1, 0.2),
		mat.MakeMaterialBuilder().
			SetTransparency(0.7).
			SetReflective(0.3).
			SetRefractiveIndex(1.5).
			SetColor(geom.Color{R: 0, G: 0, B: 0.4}).
			Build())

	light := figure.PointLight{
		Position:  geom.Point{X: 10, Y: 10, Z: -10},
		Intensity: geom.White,
	}
	world := figure.World{
		Light:   light,
		Objects: []ddddf.Shape{floor, back, left, middle, right},
	}
	camera := figure.NewCamera(500, 250, math.Pi/3,
		figure.ViewTransform(geom.Point{X: 0, Y: 3, Z: -6}, geom.Point{X: 0, Y: 1, Z: 0}, geom.Vector{X: 0, Y: 1, Z: 0}))

	canvas := camera.Render(world)

	outFile := "refraction_sample_test.png"
	canvas.MustToPNG(outFile)

	if AssertFilesEqual(t, "testdata/"+outFile, outFile) {
		_ = os.Remove(outFile)
	}
}
