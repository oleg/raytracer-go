package samples

import (
	"github.com/oleg/raytracer-go/multid"
	"github.com/oleg/raytracer-go/oned"
	"os"
	"testing"
)

func Test(t *testing.T) {
	//todo: do I want to export this methods?
	start := oned.Point{X: 0, Y: 1, Z: 0}
	velocity := oned.Vector{X: 1, Y: 1.8, Z: 0}.Normalize().MultiplyScalar(11.25)
	p := projectile{start, velocity}

	gravity := oned.Vector{X: 0, Y: -0.1, Z: 0}
	wind := oned.Vector{X: -0.01, Y: 0, Z: 0}
	e := environment{gravity, wind}

	width := 900
	height := 500
	c := multid.NewCanvas(width, height)

	for p.position.X >= 0 && p.position.Y > 0 {
		x := int(p.position.X)
		y := int(p.position.Y)
		c.Pixels[x][height-y] = oned.Color{R: 1, G: 0, B: 0}
		p = p.tick(e)
	}

	outFile := "canvas_sample_test.png"
	c.MustToPNG(outFile)

	if AssertFilesEqual(t, "testdata/"+outFile, outFile) {
		_ = os.Remove(outFile)
	}
}
