package config

import (
	"github.com/oleg/raytracer-go/geom"
	"github.com/oleg/raytracer-go/physic"
	"github.com/oleg/raytracer-go/scene"
	"github.com/oleg/raytracer-go/shapes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_read_world_from_config(t *testing.T) {
	sceneConfig := `
camera:
  size:
    horizontal: 1000
    vertical: 500
    field-of-view: 1.2
  sight:
    from: { x: 1, y: 2, z: 3 }
    to: { x: 30, y: 20, z: 10 }
    up: { x: 5, y: 3, z: 4 }
world:
  light:
    position: { x: 10, y: 10, z: -10 }
    intensity: { r: 0.8, g: 0.9, b: 1 }
  objects:
    - type: Plane
      transform: [
        [ 1, 0, 0, 0 ],
        [ 0, 0.00000000000000006, 1, 0 ],
        [ 0, -1, 0.00000000000000006, 4 ],
        [ 0, 0, 0, 1 ] ]
      material:
        color: { r: 0.2, g: 0.3, b: 0.4 }
        ambient: 0.5
        diffuse: 0.6
        specular: 0.7
        reflective: 0.8
        transparency: 0.9
        refractive-index: 3
        shininess: 198
        pattern:
          type: Checkers
          color-a: { r: 0.3, g: 0.4, b: 0.5 }
          color-b: { r: 0.6, g: 0.7, b: 0.8 }
          transform: [
            [ 0.1, 0, 0, 1 ],
            [ 0, 0.2, 0, 2 ],
            [ -1, 0, 0.3, 3 ],
            [ 0, 0, 0, 0.4 ] ]
    - type: Sphere
      transform: [
        [ 1, 0, 0, -2.4 ],
        [ 0, 1, 0, 1 ],
        [ 0, 0, 1, 0.2 ],
        [ 0, 0, 0, 1 ] ]
      material:
        color: { r: 0.2, g: 0.3, b: 0.4 }
        ambient: 0.5
        diffuse: 0.6
        specular: 0.7
        reflective: 0.8
        transparency: 0.9
        refractive-index: 3
        shininess: 198
        pattern:
          type: Checkers
          color-a: { r: 0.3, g: 0.4, b: 0.5 }
          color-b: { r: 0.6, g: 0.7, b: 0.8 }
          transform: [
            [ 0.1, 0, 0, 1 ],
            [ 0, 0.2, 0, 2 ],
            [ -1, 0, 0.3, 3 ],
            [ 0, 0, 0, 0.4 ] ]
`
	camera, world, err := ReadScene(strings.NewReader(sceneConfig))

	assert.NoError(t, err)

	assert.EqualValues(t,
		&scene.Camera{
			HSize:       1000,
			VSize:       500,
			FieldOfView: 1.2,
			HalfWidth:   0.6841368083416923,
			HalfHeight:  0.34206840417084616,
			PixelSize:   0.0013682736166833846,
			Transform: geom.NewMatrix([4][4]float64{
				{0.20700261440939072, -0.32876885817962054, -0.012176624377022993, 0.48706497508091934},
				{-0.05976042483504099, -0.05172247295664368, 0.3805795476336822, -0.9785332721527182},
				{-0.8323167879989226, -0.5166104201372623, -0.200904052275602, 2.4682497851002534},
				{0, 0, 0, 1},
			},
			),
		},
		camera)

	assert.EqualValues(t,
		scene.World{
			Light: scene.PointLight{
				Position:  geom.Point{X: 10, Y: 10, Z: -10},
				Intensity: geom.Color{R: 0.8, G: 0.9, B: 1},
			},
			Objects: []shapes.Shape{
				shapes.NewPlane(
					geom.NewMatrix([4][4]float64{
						{1, 0, 0, 0},
						{0, 0.00000000000000006, 1, 0},
						{0, -1, 0.00000000000000006, 4},
						{0, 0, 0, 1},
					}),
					&physic.Material{
						Color:           geom.Color{R: 0.2, G: 0.3, B: 0.4},
						Ambient:         0.5,
						Diffuse:         0.6,
						Specular:        0.7,
						Reflective:      0.8,
						Transparency:    0.9,
						RefractiveIndex: 3,
						Shininess:       198,
						Pattern: physic.CheckersPattern{
							A: geom.Color{R: 0.3, G: 0.4, B: 0.5},
							B: geom.Color{R: 0.6, G: 0.7, B: 0.8},
							Transformable: physic.Transformable{
								Rule: geom.NewMatrix([4][4]float64{
									{0.1, 0, 0, 1},
									{0, 0.2, 0, 2},
									{-1, 0, 0.3, 3},
									{0, 0, 0, 0.4}}),
							},
						},
					}),
				shapes.NewSphere(
					geom.NewMatrix([4][4]float64{
						{1, 0, 0, -2.4},
						{0, 1, 0, 1},
						{0, 0, 1, 0.2},
						{0, 0, 0, 1},
					}),
					&physic.Material{
						Color:           geom.Color{R: 0.2, G: 0.3, B: 0.4},
						Ambient:         0.5,
						Diffuse:         0.6,
						Specular:        0.7,
						Reflective:      0.8,
						Transparency:    0.9,
						RefractiveIndex: 3,
						Shininess:       198,
						Pattern: physic.CheckersPattern{
							A: geom.Color{R: 0.3, G: 0.4, B: 0.5},
							B: geom.Color{R: 0.6, G: 0.7, B: 0.8},
							Transformable: physic.Transformable{
								Rule: geom.NewMatrix([4][4]float64{
									{0.1, 0, 0, 1},
									{0, 0.2, 0, 2},
									{-1, 0, 0.3, 3},
									{0, 0, 0, 0.4}},
								),
							},
						},
					},
				),
			},
		},
		world)
}
