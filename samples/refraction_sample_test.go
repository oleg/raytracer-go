package samples

import (
	"bytes"
	"github.com/oleg/raytracer-go/geom"
	"github.com/oleg/raytracer-go/physic"
	"github.com/oleg/raytracer-go/scene"
	"github.com/oleg/raytracer-go/shapes"
	"github.com/stretchr/testify/assert"
	"image/png"
	"math"
	"testing"
)

func Test_refraction_sample(t *testing.T) {
	floor := shapes.NewPlane(
		geom.IdentityMatrix(),
		physic.NewMaterialBuilder().
			SetReflective(0.7).
			SetTransparency(0.2).
			SetRefractiveIndex(1.3).
			SetPattern(physic.CheckersPattern{
				A:             geom.Black,
				B:             geom.White,
				Transformable: physic.NoTransformation}).
			Build())

	back := shapes.NewPlane(
		geom.Translation(0, 0, 4).
			Multiply(geom.RotationX(-math.Pi/2)),
		physic.NewMaterialBuilder().
			SetReflective(0.3).
			SetTransparency(0.1).
			SetRefractiveIndex(2).
			SetPattern(physic.CheckersPattern{
				A:             geom.Black,
				B:             geom.White,
				Transformable: physic.NoTransformation}).
			Build())
	left := shapes.NewSphere(
		geom.Translation(-2.4, 1, 0.2),
		physic.NewMaterialBuilder().
			//SetSpecular(1).
			SetTransparency(0.3).
			SetReflective(0.3).
			SetRefractiveIndex(1).
			SetAmbient(0.2).
			SetColor(geom.White).
			Build())

	middle := shapes.NewSphere(
		geom.Translation(-0.1, 1, 0.2),
		physic.NewMaterialBuilder().
			SetTransparency(0.5).
			SetReflective(0.3).
			SetRefractiveIndex(1.2).
			SetColor(geom.Color{R: 0.4, G: 0, B: 0}).
			Build())

	right := shapes.NewSphere(
		geom.Translation(2.2, 1, 0.2),
		physic.NewMaterialBuilder().
			SetTransparency(0.7).
			SetReflective(0.3).
			SetRefractiveIndex(1.5).
			SetColor(geom.Color{R: 0, G: 0, B: 0.4}).
			Build())

	light := scene.PointLight{
		Position:  geom.Point{X: 10, Y: 10, Z: -10},
		Intensity: geom.White,
	}
	world := scene.World{
		Light:   light,
		Objects: []shapes.Shape{floor, back, left, middle, right},
	}
	camera := scene.NewCamera(500, 250, math.Pi/3,
		scene.ViewTransform(geom.Point{X: 0, Y: 3, Z: -6}, geom.Point{X: 0, Y: 1, Z: 0}, geom.Vector{X: 0, Y: 1, Z: 0}))

	canvas := camera.Render(world)

	b := new(bytes.Buffer)
	err := png.Encode(b, canvas)
	assert.NoError(t, err)
	AssertBytesAreEqual(t, "testdata/refraction_sample_test.png", b.Bytes())
}
