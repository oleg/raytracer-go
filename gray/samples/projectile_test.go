package samples

import (
	"gray/oned"
	"testing"
)

func Test_tick(t *testing.T) {
	p := projectile{oned.Point{0, 1, 0}, oned.Vector{1, 1, 0}.Normalize()}
	e := environment{oned.Vector{0, -0.1, 0}, oned.Vector{-0.01, 0, 0}}
	for p.position.Y > 0 {
		p = p.tick(e)
		//fmt.Println(p)
	}
}
