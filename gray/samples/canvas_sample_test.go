package samples

import (
	"gray/multid"
	"gray/oned"
	"testing"
)

func Test(t *testing.T) {
	//todo: do I want to export this methods?
	start := oned.Point{0, 1, 0}
	velocity := oned.Vector{1, 1.8, 0}.Normalize().MultiplyScalar(11.25)
	p := projectile{start, velocity}

	gravity := oned.Vector{0, -0.1, 0}
	wind := oned.Vector{-0.01, 0, 0}
	e := environment{gravity, wind}

	width := 900
	height := 500
	c := multid.MakeCanvas(width, height)

	for p.position.X >= 0 && p.position.Y > 0 {
		x := int(p.position.X)
		y := int(p.position.Y)
		c.Pixels[x][height-y] = oned.NewColor(1, 0, 0)
		p = p.tick(e)
	}

	c.MustToPNG("canvas_sample_test.png")
}
