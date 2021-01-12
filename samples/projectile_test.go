package samples

import (
	"github.com/oleg/raytracer-go/geom"
	"testing"
)

func Test_tick(t *testing.T) {
	p := projectile{geom.Point{X: 0, Y: 1, Z: 0}, geom.Vector{X: 1, Y: 1, Z: 0}.Normalize()}
	e := environment{geom.Vector{X: 0, Y: -0.1, Z: 0}, geom.Vector{X: -0.01, Y: 0, Z: 0}}
	for p.position.Y > 0 {
		p = p.tick(e)
		//fmt.Println(p)
	}
}
