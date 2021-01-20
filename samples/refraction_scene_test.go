package samples

import (
	"bytes"
	"github.com/oleg/raytracer-go/config"
	"github.com/stretchr/testify/assert"
	"image/png"
	"os"
	"testing"
)

func Test_refraction_scene(t *testing.T) {
	configFile, err := os.Open("testdata/refraction-scene.yaml")
	assert.NoError(t, err)

	camera, world, err := config.ReadScene(configFile)
	assert.NoError(t, err)

	canvas := camera.Render(world)

	b := new(bytes.Buffer)
	err = png.Encode(b, canvas)
	assert.NoError(t, err)

	AssertBytesAreEqual(t, "testdata/refraction_scene.png", b.Bytes())
}
