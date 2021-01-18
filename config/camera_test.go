package config

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"testing"
)

func Test_read_camera_config_from_string(t *testing.T) {
	cameraConfig := `

`
	var c camera
	err := yaml.Unmarshal([]byte(cameraConfig), c)
	assert.NoError(t, err)

	if c.hSize != 10 ||
		c.vSize != 5 ||
		c.fieldOfView != 3.312 {
		t.Errorf("Wrong camera config %v", c)
	}
}
