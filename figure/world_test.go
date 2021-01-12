package figure

import (
	"github.com/oleg/raytracer-go/geom"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_default_world(t *testing.T) {
	light := PointLight{geom.Point{X: -10, Y: 10, Z: -10}, geom.White}

	material := DefaultMaterial()
	material.Color = geom.Color{R: 0.8, G: 1.0, B: 0.6}
	material.Diffuse = 0.7
	material.Specular = 0.2
	s1 := NewSphere(geom.IdentityMatrix(), material)

	transform := geom.Scaling(0.5, 0.5, 0.5)
	s2 := NewSphere(transform, DefaultMaterial())

	w := defaultWorld()

	assert.Equal(t, light, w.Light)
	assert.Equal(t, s1, w.Objects[0])
	assert.Equal(t, s2, w.Objects[1])
}

func Test_Intersect_world_with_ray(t *testing.T) {
	w := defaultWorld()
	r := Ray{geom.Point{X: 0, Y: 0, Z: -5}, geom.Vector{X: 0, Y: 0, Z: 1}}

	xs := w.Intersect(r)

	assert.Equal(t, 4, len(xs))
	assert.Equal(t, 4.0, xs[0].Distance)
	assert.Equal(t, 4.5, xs[1].Distance)
	assert.Equal(t, 5.5, xs[2].Distance)
	assert.Equal(t, 6.0, xs[3].Distance)
}

func Test_shading_intersection(t *testing.T) {
	w := defaultWorld()
	r := Ray{geom.Point{X: 0, Y: 0, Z: -5}, geom.Vector{X: 0, Y: 0, Z: 1}}
	shape := w.Objects[0]
	i := Inter{4, shape}
	comps := i.PrepareComputations(r, Inters{i})

	c := w.ShadeHit(comps, MaxDepth)

	geom.AssertColorEqualInDelta(t, geom.Color{R: 0.38066, G: 0.47583, B: 0.2855}, c)
}

func Test_shading_intersection_from_inside(t *testing.T) {
	w := defaultWorld()
	w.Light = PointLight{geom.Point{X: 0, Y: 0.25, Z: 0}, geom.White}
	r := Ray{geom.Point{X: 0, Y: 0, Z: 0}, geom.Vector{X: 0, Y: 0, Z: 1}}
	shape := w.Objects[1]
	i := Inter{0.5, shape}
	comps := i.PrepareComputations(r, Inters{i})

	c := w.ShadeHit(comps, MaxDepth)

	geom.AssertColorEqualInDelta(t, geom.Color{R: 0.90498, G: 0.90498, B: 0.90498}, c)
}

func Test_color_when_ray_misses(t *testing.T) {
	w := defaultWorld()
	r := Ray{geom.Point{X: 0, Y: 0, Z: -5}, geom.Vector{X: 0, Y: 1, Z: 0}}

	c := w.ColorAt(r, MaxDepth)

	geom.AssertColorEqualInDelta(t, geom.Black, c)
}

func Test_color_when_ray_hits(t *testing.T) {
	w := defaultWorld()
	r := Ray{geom.Point{X: 0, Y: 0, Z: -5}, geom.Vector{X: 0, Y: 0, Z: 1}}

	c := w.ColorAt(r, MaxDepth)

	geom.AssertColorEqualInDelta(t, geom.Color{R: 0.38066, G: 0.47583, B: 0.2855}, c)
}

func Test_color_with_intersection_behind_ray(t *testing.T) {
	w := World{pointLightSample(), []Shape{
		NewSphere(geom.IdentityMatrix(), testMaterialBuilder().SetAmbient(1).Build()),
		NewSphere(geom.Scaling(0.5, 0.5, 0.5), testMaterialBuilder().SetAmbient(1).Build())}}
	r := Ray{geom.Point{X: 0, Y: 0, Z: 0.75}, geom.Vector{X: 0, Y: 0, Z: -1}}

	c := w.ColorAt(r, MaxDepth)

	geom.AssertColorEqualInDelta(t, testMaterialBuilder().SetAmbient(1).Build().Color, c)
}

func Test_shade_hit_is_given_intersection_in_shadow(t *testing.T) {
	s1 := NewSphere(geom.IdentityMatrix(), DefaultMaterial())
	s2 := NewSphere(geom.Translation(0, 0, 10), DefaultMaterial())
	w := World{
		PointLight{geom.Point{X: 0, Y: 0, Z: -10}, geom.White},
		[]Shape{s1, s2},
	}
	r := Ray{geom.Point{X: 0, Y: 0, Z: 5}, geom.Vector{X: 0, Y: 0, Z: 1}}
	i := Inter{4, s2}
	comps := i.PrepareComputations(r, Inters{i})

	color := w.ShadeHit(comps, MaxDepth)

	assert.Equal(t, geom.Color{R: 0.1, G: 0.1, B: 0.1}, color)
}

func Test_hit_should_offset_point(t *testing.T) {
	r := Ray{geom.Point{X: 0, Y: 0, Z: -5}, geom.Vector{X: 0, Y: 0, Z: 1}}
	s := NewSphere(geom.Translation(0, 0, 1), DefaultMaterial())
	i := Inter{5, s}

	comps := i.PrepareComputations(r, Inters{i})

	assert.Less(t, comps.OverPoint.Z, -geom.Delta/2)
	assert.Less(t, comps.OverPoint.Z, comps.Point.Z)
}

func Test_reflected_color_for_non_reflective_material(t *testing.T) {
	s1 := NewSphere(geom.IdentityMatrix(), testMaterialBuilder().Build())
	s2 := NewSphere(geom.Scaling(0.5, 0.5, 0.5), testMaterialBuilder().SetAmbient(1).Build())
	w := World{pointLightSample(), []Shape{s1, s2}}
	r := Ray{geom.Point{X: 0, Y: 0, Z: 0}, geom.Vector{X: 0, Y: 0, Z: 1}}
	i := Inter{1, s2}
	comps := i.PrepareComputations(r, Inters{i})

	color := w.ReflectedColor(comps, 5)

	assert.Equal(t, geom.Color{R: 0, G: 0, B: 0}, color)
}

func Test_reflected_color_for_reflective_material(t *testing.T) {
	s1 := NewSphere(geom.IdentityMatrix(), testMaterialBuilder().Build())
	s2 := NewSphere(geom.Scaling(0.5, 0.5, 0.5), testMaterialBuilder().SetAmbient(1).Build())
	s3 := MakePlaneTM(geom.Translation(0, -1, 0), MakeMaterialBuilder().SetReflective(0.5).Build())
	w := World{pointLightSample(), []Shape{s1, s2, s3}}
	r := Ray{geom.Point{X: 0, Y: 0, Z: -3}, geom.Vector{X: 0, Y: -math.Sqrt2 / 2, Z: math.Sqrt2 / 2}}
	i := Inter{math.Sqrt2, s3}
	comps := i.PrepareComputations(r, Inters{i})

	color := w.ReflectedColor(comps, 5)

	geom.AssertColorEqualInDelta(t, geom.Color{R: 0.19033, G: 0.23791, B: 0.142749}, color)
}

func Test_shade_hit_with_reflective_material(t *testing.T) {
	s1 := NewSphere(geom.IdentityMatrix(), testMaterialBuilder().Build())
	s2 := NewSphere(geom.Scaling(0.5, 0.5, 0.5), testMaterialBuilder().SetAmbient(1).Build())
	s3 := MakePlaneTM(geom.Translation(0, -1, 0), MakeMaterialBuilder().SetReflective(0.5).Build())
	w := World{pointLightSample(), []Shape{s1, s2, s3}}
	r := Ray{geom.Point{X: 0, Y: 0, Z: -3}, geom.Vector{X: 0, Y: -math.Sqrt2 / 2, Z: math.Sqrt2 / 2}}
	i := Inter{math.Sqrt2, s3}
	comps := i.PrepareComputations(r, Inters{i})

	color := w.ShadeHit(comps, MaxDepth)

	geom.AssertColorEqualInDelta(t, geom.Color{R: 0.87675, G: 0.92434, B: 0.82918}, color)
}

func Test_color_at_with_mutually_reflective_surfaces(t *testing.T) {
	w := World{
		PointLight{geom.Point{X: 0, Y: 0, Z: 0}, geom.Color{R: 1, G: 1, B: 1}},
		[]Shape{
			MakePlaneTM(geom.Translation(0, -1, 0), MakeMaterialBuilder().SetReflective(1).Build()),
			MakePlaneTM(geom.Translation(0, 1, 0), MakeMaterialBuilder().SetReflective(1).Build())}}

	w.ColorAt(Ray{geom.Point{X: 0, Y: 0, Z: 0}, geom.Vector{X: 0, Y: 1, Z: 0}}, MaxDepth)

	//should terminate
}

func Test_reflected_color_at_maximum_recursive_depth(t *testing.T) {
	s1 := NewSphere(geom.IdentityMatrix(), testMaterialBuilder().Build())
	s2 := NewSphere(geom.Scaling(0.5, 0.5, 0.5), testMaterialBuilder().SetAmbient(1).Build())
	s3 := MakePlaneTM(geom.Translation(0, -1, 0), MakeMaterialBuilder().SetReflective(0.5).Build())
	w := World{pointLightSample(), []Shape{s1, s2, s3}}
	r := Ray{geom.Point{X: 0, Y: 0, Z: -3}, geom.Vector{X: 0, Y: -math.Sqrt2 / 2, Z: math.Sqrt2 / 2}}
	i := Inter{math.Sqrt2, s3}
	comps := i.PrepareComputations(r, Inters{i})

	color := w.ReflectedColor(comps, 0)

	geom.AssertColorEqualInDelta(t, geom.Color{R: 0, G: 0, B: 0}, color)
}

func Test_refracted_color_with_opaque_surface(t *testing.T) {
	w := defaultWorld()
	s := w.Objects[0]
	r := Ray{geom.Point{X: 0, Y: 0, Z: -5}, geom.Vector{X: 0, Y: 0, Z: 1}}
	xs := Inters{Inter{4, s}, Inter{6, s}}
	comps := xs[0].PrepareComputations(r, xs)

	c := w.RefractedColor(comps, MaxDepth)

	assert.Equal(t, geom.Color{R: 0, G: 0, B: 0}, c)
}

func Test_refracted_color_at_the_maximum_recursive_depth(t *testing.T) {
	s1 := NewSphere(geom.IdentityMatrix(), testMaterialBuilder().SetTransparency(1.0).SetRefractiveIndex(1.5).Build())
	s2 := NewSphere(geom.Scaling(0.5, 0.5, 0.5), DefaultMaterial())
	w := World{pointLightSample(), []Shape{s1, s2}}
	r := Ray{geom.Point{X: 0, Y: 0, Z: -5}, geom.Vector{X: 0, Y: 0, Z: 1}}
	xs := Inters{Inter{4, s1}, Inter{6, s1}}
	comps := xs[0].PrepareComputations(r, xs)

	c := w.RefractedColor(comps, 0)

	assert.Equal(t, geom.Color{R: 0, G: 0, B: 0}, c)
}

func Test_refracted_color_under_total_internal_reflection(t *testing.T) {
	s1 := NewSphere(geom.IdentityMatrix(), testMaterialBuilder().SetTransparency(1.0).SetRefractiveIndex(1.5).Build())
	s2 := NewSphere(geom.Scaling(0.5, 0.5, 0.5), DefaultMaterial())
	w := World{pointLightSample(), []Shape{s1, s2}}
	r := Ray{geom.Point{X: 0, Y: 0, Z: math.Sqrt2 / 2}, geom.Vector{X: 0, Y: 1, Z: 0}}
	xs := Inters{Inter{-math.Sqrt2 / 2, s1}, Inter{math.Sqrt2 / 2, s1}}
	comps := xs[1].PrepareComputations(r, xs)

	c := w.RefractedColor(comps, MaxDepth)

	assert.Equal(t, geom.Color{R: 0, G: 0, B: 0}, c)
}

func Test_refracted_color_with_refracted_ray(t *testing.T) {
	s1 := NewSphere(geom.IdentityMatrix(),
		testMaterialBuilder().
			SetAmbient(1.0).
			SetPattern(TestPattern{}).
			Build())
	s2 := NewSphere(
		geom.Scaling(0.5, 0.5, 0.5),
		testMaterialBuilder().
			SetTransparency(1.0).
			SetRefractiveIndex(1.5).
			Build())
	w := World{pointLightSample(), []Shape{s1, s2}}
	r := Ray{geom.Point{X: 0, Y: 0, Z: 0.1}, geom.Vector{X: 0, Y: 1, Z: 0}}
	xs := Inters{Inter{-0.9899, s1}, Inter{-0.4899, s2}, Inter{0.4899, s2}, Inter{0.9899, s1}}
	comps := xs[2].PrepareComputations(r, xs)

	c := w.RefractedColor(comps, MaxDepth)

	geom.AssertColorEqualInDelta(t, geom.Color{R: 0, G: 0.99888, B: 0.04721}, c)
}

func Test_shade_hit_with_transparent_material(t *testing.T) {
	s1 := NewSphere(geom.IdentityMatrix(), testMaterialBuilder().Build())
	s2 := NewSphere(geom.Scaling(0.5, 0.5, 0.5), DefaultMaterial())
	floor := MakePlaneTM(geom.Translation(0, -1, 0),
		MakeMaterialBuilder().
			SetTransparency(0.5).
			SetRefractiveIndex(1.5).
			Build())
	ball := NewSphere(geom.Translation(0, -3.5, -0.5),
		MakeMaterialBuilder().
			SetColor(geom.Color{R: 1, G: 0, B: 0}).
			SetAmbient(0.5).
			Build())
	w := World{pointLightSample(), []Shape{s1, s2, floor, ball}}
	r := Ray{geom.Point{X: 0, Y: 0, Z: -3}, geom.Vector{X: 0, Y: -math.Sqrt2 / 2, Z: math.Sqrt2 / 2}}
	xs := Inters{Inter{math.Sqrt2, floor}}
	comps := xs[0].PrepareComputations(r, xs)

	color := w.ShadeHit(comps, MaxDepth)

	geom.AssertColorEqualInDelta(t, geom.Color{R: 0.93642, G: 0.68642, B: 0.68642}, color)
}

func Test_shade_hit_with_reflective_transparent_material(t *testing.T) {
	s1 := NewSphere(geom.IdentityMatrix(), testMaterialBuilder().Build())
	s2 := NewSphere(geom.Scaling(0.5, 0.5, 0.5), DefaultMaterial())
	floor := MakePlaneTM(geom.Translation(0, -1, 0),
		MakeMaterialBuilder().
			SetReflective(0.5).
			SetTransparency(0.5).
			SetRefractiveIndex(1.5).
			Build())
	ball := NewSphere(geom.Translation(0, -3.5, -0.5),
		MakeMaterialBuilder().
			SetColor(geom.Color{R: 1, G: 0, B: 0}).
			SetAmbient(0.5).
			Build())
	w := World{pointLightSample(), []Shape{s1, s2, floor, ball}}
	r := Ray{geom.Point{X: 0, Y: 0, Z: -3}, geom.Vector{X: 0, Y: -math.Sqrt2 / 2, Z: math.Sqrt2 / 2}}
	xs := Inters{Inter{math.Sqrt2, floor}}
	comps := xs[0].PrepareComputations(r, xs)

	color := w.ShadeHit(comps, MaxDepth)

	geom.AssertColorEqualInDelta(t, geom.Color{R: 0.93391, G: 0.69643, B: 0.69243}, color)
}

//util
func defaultWorld() World {
	s1 := NewSphere(geom.IdentityMatrix(), testMaterialBuilder().Build())
	s2 := NewSphere(geom.Scaling(0.5, 0.5, 0.5), DefaultMaterial())

	return World{
		pointLightSample(),
		[]Shape{s1, s2},
	}
}

func pointLightSample() PointLight {
	return PointLight{geom.Point{X: -10, Y: 10, Z: -10}, geom.White}
}

func testMaterialBuilder() *MaterialBuilder {
	return MakeMaterialBuilder().
		SetColor(geom.Color{R: 0.8, G: 1.0, B: 0.6}).
		SetDiffuse(0.7).
		SetSpecular(0.2)
}

type TestPattern struct {
}

func (t TestPattern) PatternAt(point geom.Point) geom.Color {
	return geom.Color{R: point.X, G: point.Y, B: point.Z}
}
func (t TestPattern) Transform() *geom.Matrix {
	return geom.IdentityMatrix()
}
