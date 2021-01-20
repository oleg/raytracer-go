package config

import (
	"fmt"
	"github.com/oleg/raytracer-go/geom"
	"github.com/oleg/raytracer-go/scene"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"strings"
	"testing"
)

func Test_read_camera_config_from_string(t *testing.T) {
	cameraConfig := `size:
  horizontal: 10
  vertical: 5
  field-of-view: 3.312
sight:
  from: { x: 1, y: 2, z: 3 }
  to: { x: 30, y: 20, z: 10 }
  up: { x: 5, y: 3, z: 4 } 
`

	var c Camera
	err := yaml.Unmarshal([]byte(cameraConfig), &c)
	fmt.Printf("%v\n", c)
	assert.NoError(t, err)
	assert.EqualValues(t,
		Size{
			Horizontal:  10,
			Vertical:    5,
			FieldOfView: 3.312,
		},
		c.Size)

	assert.EqualValues(t,
		Sight{
			From: Tuple{
				X: 1,
				Y: 2,
				Z: 3,
			},
			To: Tuple{
				X: 30,
				Y: 20,
				Z: 10,
			},
			Up: Tuple{
				X: 5,
				Y: 3,
				Z: 4,
			},
		},
		c.Sight)
}

func Test_create_camera(t *testing.T) {
	cameraConfig := `size:
  horizontal: 1000
  vertical: 500
  field-of-view: 1.2
sight:
  from: { x: 1, y: 2, z: 3 }
  to: { x: 30, y: 20, z: 10 }
  up: { x: 5, y: 3, z: 4 } 
`
	camera, err := ReadCamera(strings.NewReader(cameraConfig))

	assert.NoError(t, err)
	assert.EqualValues(t,
		&scene.Camera{
			HSize:       1000,
			VSize:       500,
			FieldOfView: 1.2,
			HalfWidth:   0.6841368083416923,
			HalfHeight:  0.34206840417084616,
			PixelSize:   0.0013682736166833846,
			Transform: &geom.Matrix{
				Data: [4][4]float64{
					{0.20700261440939072, -0.32876885817962054, -0.012176624377022993, 0.48706497508091934},
					{-0.05976042483504099, -0.05172247295664368, 0.3805795476336822, -0.9785332721527182},
					{-0.8323167879989226, -0.5166104201372623, -0.200904052275602, 2.4682497851002534},
					{0, 0, 0, 1},
				},
			},
		},
		camera)
}
