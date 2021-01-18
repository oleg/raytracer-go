package samples

import (
	"bytes"
	"github.com/oleg/raytracer-go/geom"
	"github.com/oleg/raytracer-go/scene"
	"github.com/stretchr/testify/assert"
	"image/png"
	"testing"
)

func Test(t *testing.T) {
	//todo: do I want to export this methods?
	start := geom.Point{X: 0, Y: 1, Z: 0}
	velocity := geom.Vector{X: 1, Y: 1.8, Z: 0}.Normalize().MultiplyScalar(11.25)
	p := projectile{start, velocity}

	gravity := geom.Vector{X: 0, Y: -0.1, Z: 0}
	wind := geom.Vector{X: -0.01, Y: 0, Z: 0}
	e := environment{gravity, wind}

	width := 900
	height := 500
	c := scene.NewCanvas(width, height)

	for p.position.X >= 0 && p.position.Y > 0 {
		x := int(p.position.X)
		y := int(p.position.Y)
		c.Pixels[x][height-y] = geom.Color{R: 1, G: 0, B: 0}
		p = p.tick(e)
	}

	b := new(bytes.Buffer)
	err := png.Encode(b, c)
	assert.NoError(t, err)
	AssertBytesAreEqual(t, "testdata/canvas_sample_test.png", b.Bytes())
}
