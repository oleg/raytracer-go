package physic

import (
	"github.com/oleg/raytracer-go/geom"
)

//todo change types?
//todo reorder members
//todo implement MaterialProvider
type Material struct {
	Color           geom.Color
	Pattern         Pattern
	Ambient         float64
	Diffuse         float64
	Specular        float64
	Shininess       float64
	Reflective      float64
	Transparency    float64
	RefractiveIndex float64
}

//todo change api, should accept overrides, use builder?
func DefaultMaterial() *Material {
	return NewMaterialBuilder().Build()
}

func GlassMaterialBuilder() *MaterialBuilder {
	return NewMaterialBuilder().
		SetRefractiveIndex(1.5).
		SetTransparency(1.0)
}

//todo think about it
type MaterialBuilder struct {
	color           geom.Color
	pattern         Pattern
	ambient         float64
	diffuse         float64
	specular        float64
	shininess       float64
	reflective      float64
	transparency    float64
	refractiveIndex float64
}

func NewMaterialBuilder() *MaterialBuilder {
	return &MaterialBuilder{
		color:           geom.White,
		ambient:         0.1,
		diffuse:         0.9,
		specular:        0.9,
		refractiveIndex: 1.0,
		shininess:       200.0,
	}
}
func (mb *MaterialBuilder) Build() *Material {
	return &Material{
		Color:           mb.color,
		Pattern:         mb.pattern,
		Ambient:         mb.ambient,
		Diffuse:         mb.diffuse,
		Specular:        mb.specular,
		Shininess:       mb.shininess,
		Reflective:      mb.reflective,
		Transparency:    mb.transparency,
		RefractiveIndex: mb.refractiveIndex,
	}
}
func (mb *MaterialBuilder) SetColor(color geom.Color) *MaterialBuilder {
	mb.color = color
	return mb
}
func (mb *MaterialBuilder) SetAmbient(ambient float64) *MaterialBuilder {
	mb.ambient = ambient
	return mb
}
func (mb *MaterialBuilder) SetDiffuse(diffuse float64) *MaterialBuilder {
	mb.diffuse = diffuse
	return mb
}
func (mb *MaterialBuilder) SetSpecular(specular float64) *MaterialBuilder {
	mb.specular = specular
	return mb
}
func (mb *MaterialBuilder) SetShininess(shininess float64) *MaterialBuilder {
	mb.shininess = shininess
	return mb
}
func (mb *MaterialBuilder) SetPattern(pattern Pattern) *MaterialBuilder {
	mb.pattern = pattern
	return mb
}
func (mb *MaterialBuilder) SetReflective(reflective float64) *MaterialBuilder {
	mb.reflective = reflective
	return mb
}
func (mb *MaterialBuilder) SetTransparency(transparency float64) *MaterialBuilder {
	mb.transparency = transparency
	return mb
}
func (mb *MaterialBuilder) SetRefractiveIndex(refractiveIndex float64) *MaterialBuilder {
	mb.refractiveIndex = refractiveIndex
	return mb
}
