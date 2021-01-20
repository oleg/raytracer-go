package config

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
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
