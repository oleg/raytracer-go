package config

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"testing"
)

func Test_read_camera_config_from_string(t *testing.T) {
	cameraConfig := `
size: 
    horizontal: 10
    vertical: 5
field-of-view: 3.312
`
	var c Camera
	err := yaml.Unmarshal([]byte(cameraConfig), &c)
	assert.NoError(t, err)

	if c.Size.Horizontal != 10 ||
		c.Size.Vertical != 5 ||
		c.FieldOfView != 3.312 {
		t.Errorf("Wrong camera config %v", c)
	}
}
