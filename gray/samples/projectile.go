package samples

import "gray/oned"

type projectile struct {
	position oned.Point
	velocity oned.Vector
}
type environment struct {
	gravity oned.Vector
	wind    oned.Vector
}

//func (p projectile) move() (position Point) {
//	return p.position.AddVector(p.velocity)
//}

func (proj projectile) tick(env environment) projectile {
	position := proj.position.AddVector(proj.velocity)
	velocity := proj.velocity.AddVector(env.gravity).AddVector(env.wind)
	return projectile{position, velocity}
}
