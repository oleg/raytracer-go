package samples

import (
	"github.com/oleg/raytracer-go/physic"
	"github.com/oleg/raytracer-go/shapes"
	"github.com/oleg/raytracer-go/scene"
	"github.com/oleg/raytracer-go/geom"
	"math"
	"os"
	"testing"
)

func Test_plane_scene_sample(t *testing.T) {
	floor := shapes.NewPlane(
		geom.IdentityMatrix(),
		physic.MakeMaterialBuilder().
			SetReflective(0.1).
			SetPattern(physic.MakeCheckersPatternT(
				geom.Color{R: 0.5, G: 1, B: 0.1},
				geom.Color{R: 0.7, G: 0.3, B: 1},
				geom.Translation(1, 0, 0).
					Multiply(geom.Scaling(0.5, 0.5, 0.5)))).
			Build())

	back := shapes.NewPlane(
		geom.Translation(0, 0, 3).
			Multiply(geom.RotationX(-math.Pi/2)),
		physic.MakeMaterialBuilder().
			SetReflective(0.3).
			SetPattern(physic.MakeRingPatternT(
				geom.Color{R: 0.8, G: 0.9, B: 0.5},
				geom.Color{R: 0.5, G: 0.2, B: 0.3},
				geom.Translation(0, 0, 2).
					Multiply(geom.Scaling(0.2, 0.2, 0.2)))).
			Build())

	left := shapes.NewSphere(
		geom.Translation(-1.5, 0.33, -0.75).
			Multiply(geom.Scaling(1, 0.33, 0.33)),
		physic.MakeMaterialBuilder().
			SetPattern(physic.MakeGradientPatternT(
				geom.Color{R: 0.3, G: 1, B: 0.7},
				geom.Color{R: 0.7, G: 0.3, B: 1},
				geom.Translation(1, 0, 0).
					Multiply(geom.Scaling(2, 1, 1)))).
			SetDiffuse(0.7).
			SetSpecular(0.3).Build())

	middle := shapes.NewSphere(
		geom.Translation(-0.5, 1, 0.2),
		physic.MakeMaterialBuilder().
			SetDiffuse(0.7).
			SetSpecular(0.3).Build())

	right := shapes.NewSphere(
		geom.Translation(1.5, 0.5, -0.5).
			Multiply(geom.Scaling(0.5, 0.8, 0.5)),
		physic.MakeMaterialBuilder().
			SetPattern(physic.MakeStripePatternT(
				geom.Color{R: 0.7, G: 0.9, B: 0.8},
				geom.Color{R: 0.2, G: 0.4, B: 0.1},
				geom.RotationZ(math.Pi/4).
					Multiply(geom.Scaling(0.3, 0.3, 0.3)))).
			SetDiffuse(0.7).
			SetSpecular(0.3).Build())

	light := scene.PointLight{
		Position:  geom.Point{X: -10, Y: 10, Z: -10},
		Intensity: geom.White,
	}
	world := scene.World{
		Light:   light,
		Objects: []shapes.Shape{floor, back, left, middle, right},
	}
	camera := scene.NewCamera(500, 250, math.Pi/3,
		scene.ViewTransform(geom.Point{X: 0, Y: 3, Z: -6}, geom.Point{X: 0, Y: 1, Z: 0}, geom.Vector{X: 0, Y: 1, Z: 0}))

	canvas := camera.Render(world)

	outFile := "plane_scene_sample_test.png"
	canvas.MustToPNG(outFile)

	if AssertFilesEqual(t, "testdata/"+outFile, outFile) {
		_ = os.Remove(outFile)
	}
}
