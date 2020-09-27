package samples

import (
	"gray/figure"
	"gray/multid"
	"gray/oned"
	"math"
	"testing"
)

func Test_refraction_sample(t *testing.T) {
	floor := figure.MakePlaneTM(
		multid.IdentityMatrix,
		figure.MakeMaterialBuilder().
			SetReflective(0.7).
			SetTransparency(0.2).
			SetRefractiveIndex(1.3).
			SetPattern(figure.MakeCheckersPatternT(
				oned.Black,
				oned.White,
				multid.IdentityMatrix)).
			Build())

	back := figure.MakePlaneTM(
		multid.Translation(0, 0, 4).
			Multiply(multid.RotationX(-math.Pi/2)),
		figure.MakeMaterialBuilder().
			SetReflective(0.3).
			SetTransparency(0.1).
			SetRefractiveIndex(2).
			SetPattern(figure.MakeCheckersPatternT(
				oned.Black,
				oned.White,
				multid.IdentityMatrix)).
			Build())
	left := figure.MakeSphereTM(
		multid.Translation(-2.4, 1, 0.2),
		figure.MakeMaterialBuilder().
			//SetSpecular(1).
			SetTransparency(0.3).
			SetReflective(0.3).
			SetRefractiveIndex(1).
			SetAmbient(0.2).
			SetColor(oned.White).
			Build())

	middle := figure.MakeSphereTM(
		multid.Translation(-0.1, 1, 0.2),
		figure.MakeMaterialBuilder().
			SetTransparency(0.5).
			SetReflective(0.3).
			SetRefractiveIndex(1.2).
			SetColor(oned.Color{0, 0, 0.4}).
			Build())

	right := figure.MakeSphereTM(
		multid.Translation(2.2, 1, 0.2),
		figure.MakeMaterialBuilder().
			SetTransparency(0.7).
			SetReflective(0.3).
			SetRefractiveIndex(1.5).
			SetColor(oned.Color{0.4, 0, 0}).
			Build())

	light := figure.PointLight{oned.Point{10, 10, -10}, oned.White}
	world := figure.World{light, []figure.Shape{floor, back, left, middle, right}}
	camera := figure.MakeCamera(500, 250, math.Pi/3,
		figure.ViewTransform(oned.Point{0, 3, -6}, oned.Point{0, 1, 0}, oned.Vector{0, 1, 0}))

	canvas := camera.Render(world)

	canvas.MustToPNG("refraction_sample_test.png")
}
