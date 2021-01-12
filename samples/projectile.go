package samples

import "github.com/oleg/raytracer-go/geom"

type projectile struct {
	position geom.Point
	velocity geom.Vector
}
type environment struct {
	gravity geom.Vector
	wind    geom.Vector
}

//func (p projectile) move() (position Point) {
//	return p.position.AddVector(p.velocity)
//}

func (proj projectile) tick(env environment) projectile {
	position := proj.position.AddVector(proj.velocity)
	velocity := proj.velocity.AddVector(env.gravity).AddVector(env.wind)
	return projectile{position, velocity}
}
