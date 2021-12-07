package physic

import "github.com/oleg/raytracer-go/geom"

type TransformationProvider interface {
	Transformation() *geom.Matrix //todo add type alias?
}

var NoTransformation = Transformable{geom.IdentityMatrix()}

// Transformable implements TransformationProvider and could be added to other structs
type Transformable struct {
	Rule *geom.Matrix //todo name it rule?
}

func (t Transformable) Transformation() *geom.Matrix {
	return t.Rule
}

type MaterialProvider interface {
	Material() *Material
}

type PhysicalObject struct {
	transform *geom.Matrix //todo use transformable?
	material  *Material
}

//todo remove constructor?
func NewPhysicalObject(transform *geom.Matrix, material *Material) PhysicalObject {
	return PhysicalObject{transform, material}
}

func (p PhysicalObject) Transformation() *geom.Matrix {
	return p.transform
}
func (p PhysicalObject) Material() *Material {
	return p.material
}
