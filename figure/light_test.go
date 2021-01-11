package figure

import (
	"github.com/oleg/raytracer-go/oned"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_point_light_has_position_and_intensity(t *testing.T) {
	intensity := oned.White
	position := oned.Point{X: 0, Y: 0, Z: 0}

	light := PointLight{position, intensity}

	assert.Equal(t, position, light.Position)
	assert.Equal(t, intensity, light.Intensity)
}
