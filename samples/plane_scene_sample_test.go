package samples

import (
	"github.com/oleg/raytracer-go/figure"
	"github.com/oleg/raytracer-go/multid"
	"github.com/oleg/raytracer-go/oned"
	"math"
	"os"
	"testing"
)

func Test_plane_scene_sample(t *testing.T) {
	floor := figure.MakePlaneTM(
		multid.IdentityMatrix(),
		figure.MakeMaterialBuilder().
			SetReflective(0.1).
			SetPattern(figure.MakeCheckersPatternT(
				oned.Color{R: 0.5, G: 1, B: 0.1},
				oned.Color{R: 0.7, G: 0.3, B: 1},
				multid.Translation(1, 0, 0).
					Multiply(multid.Scaling(0.5, 0.5, 0.5)))).
			Build())

	back := figure.MakePlaneTM(
		multid.Translation(0, 0, 3).
			Multiply(multid.RotationX(-math.Pi/2)),
		figure.MakeMaterialBuilder().
			SetReflective(0.3).
			SetPattern(figure.MakeRingPatternT(
				oned.Color{R: 0.8, G: 0.9, B: 0.5},
				oned.Color{R: 0.5, G: 0.2, B: 0.3},
				multid.Translation(0, 0, 2).
					Multiply(multid.Scaling(0.2, 0.2, 0.2)))).
			Build())

	left := figure.MakeSphereTM(
		multid.Translation(-1.5, 0.33, -0.75).
			Multiply(multid.Scaling(1, 0.33, 0.33)),
		figure.MakeMaterialBuilder().
			SetPattern(figure.MakeGradientPatternT(
				oned.Color{R: 0.3, G: 1, B: 0.7},
				oned.Color{R: 0.7, G: 0.3, B: 1},
				multid.Translation(1, 0, 0).
					Multiply(multid.Scaling(2, 1, 1)))).
			SetDiffuse(0.7).
			SetSpecular(0.3).Build())

	middle := figure.MakeSphereTM(
		multid.Translation(-0.5, 1, 0.2),
		figure.MakeMaterialBuilder().
			SetDiffuse(0.7).
			SetSpecular(0.3).Build())

	right := figure.MakeSphereTM(
		multid.Translation(1.5, 0.5, -0.5).
			Multiply(multid.Scaling(0.5, 0.8, 0.5)),
		figure.MakeMaterialBuilder().
			SetPattern(figure.MakeStripePatternT(
				oned.Color{R: 0.7, G: 0.9, B: 0.8},
				oned.Color{R: 0.2, G: 0.4, B: 0.1},
				multid.RotationZ(math.Pi/4).
					Multiply(multid.Scaling(0.3, 0.3, 0.3)))).
			SetDiffuse(0.7).
			SetSpecular(0.3).Build())

	light := figure.PointLight{
		Position:  oned.Point{X: -10, Y: 10, Z: -10},
		Intensity: oned.White,
	}
	world := figure.World{
		Light:   light,
		Objects: []figure.Shape{floor, back, left, middle, right},
	}
	camera := figure.NewCamera(500, 250, math.Pi/3,
		figure.ViewTransform(oned.Point{X: 0, Y: 3, Z: -6}, oned.Point{X: 0, Y: 1, Z: 0}, oned.Vector{X: 0, Y: 1, Z: 0}))

	canvas := camera.Render(world)

	outFile := "plane_scene_sample_test.png"
	canvas.MustToPNG(outFile)

	if AssertFilesEqual(t, "testdata/"+outFile, outFile) {
		_ = os.Remove(outFile)
	}
}
