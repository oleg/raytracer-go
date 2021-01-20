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

func Test_ball_scene_sample(t *testing.T) {
	floorMaterial := physic.NewMaterialBuilder().
		SetColor(geom.Color{R: 1, G: 0.9, B: 0.9}).
		SetSpecular(0).
		Build()

	floor := shapes.NewSphere(
		geom.Scaling(10, 0.01, 10),
		floorMaterial)

	leftWall := shapes.NewSphere(
		geom.Translation(0, 0, 5).
			Multiply(geom.RotationY(-math.Pi/4)).
			Multiply(geom.RotationX(math.Pi/2)).
			Multiply(geom.Scaling(10, 0.01, 10)),
		floorMaterial)

	rightWall := shapes.NewSphere(
		geom.Translation(0, 0, 5).
			Multiply(geom.RotationY(math.Pi/4)).
			Multiply(geom.RotationX(math.Pi/2)).
			Multiply(geom.Scaling(10, 0.01, 10)),
		floorMaterial)

	middle := shapes.NewSphere(
		geom.Translation(-0.5, 1, 0.5),
		physic.NewMaterialBuilder().
			SetColor(geom.Color{R: 0.1, G: 1, B: 0.5}).
			SetDiffuse(0.7).
			SetSpecular(0.3).Build())

	right := shapes.NewSphere(
		geom.Translation(1.5, 0.5, -0.5).
			Multiply(geom.Scaling(0.5, 0.5, 0.5)),
		physic.NewMaterialBuilder().
			SetColor(geom.Color{R: 0.5, G: 1, B: 0.1}).
			SetDiffuse(0.7).
			SetSpecular(0.3).Build())

	left := shapes.NewSphere(
		geom.Translation(-1.5, 0.33, -0.75).
			Multiply(geom.Scaling(0.33, 0.33, 0.33)),
		physic.NewMaterialBuilder().
			SetColor(geom.Color{R: 1, G: 0.8, B: 0.1}).
			SetDiffuse(0.7).
			SetSpecular(0.3).Build())

	light := scene.PointLight{
		Position:  geom.Point{X: -10, Y: 10, Z: -10},
		Intensity: geom.White,
	}
	world := scene.World{
		Light: light,
		Objects: []shapes.Shape{
			floor, leftWall, rightWall, middle, right, left,
		},
	}
	camera := scene.NewCamera(500, 250, math.Pi/3,
		scene.Sight{
			From: geom.Point{X: 0, Y: 1.5, Z: -5},
			To:   geom.Point{X: 0, Y: 1, Z: 0},
			Up:   geom.Vector{X: 0, Y: 1, Z: 0}})

	canvas := camera.Render(world)

	b := new(bytes.Buffer)
	err := png.Encode(b, canvas)
	assert.NoError(t, err)
	AssertBytesAreEqual(t, "testdata/ball_scene_sample_test.png", b.Bytes())

}
