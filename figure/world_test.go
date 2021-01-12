package figure

import (
	"github.com/oleg/raytracer-go/asdf"
	"github.com/oleg/raytracer-go/ddddf"
	"github.com/oleg/raytracer-go/geom"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_default_world(t *testing.T) {
	light := PointLight{geom.Point{X: -10, Y: 10, Z: -10}, geom.White}

	material := asdf.DefaultMaterial()
	material.Color = geom.Color{R: 0.8, G: 1.0, B: 0.6}
	material.Diffuse = 0.7
	material.Specular = 0.2
	s1 := ddddf.NewSphere(geom.IdentityMatrix(), material)

	transform := geom.Scaling(0.5, 0.5, 0.5)
	s2 := ddddf.NewSphere(transform, asdf.DefaultMaterial())

	w := defaultWorld()

	assert.Equal(t, light, w.Light)
	assert.Equal(t, s1, w.Objects[0])
	assert.Equal(t, s2, w.Objects[1])
}

func Test_Intersect_world_with_ray(t *testing.T) {
	w := defaultWorld()
	r := ddddf.Ray{geom.Point{X: 0, Y: 0, Z: -5}, geom.Vector{X: 0, Y: 0, Z: 1}}

	xs := w.Intersect(r)

	assert.Equal(t, 4, len(xs))
	assert.Equal(t, 4.0, xs[0].Distance)
	assert.Equal(t, 4.5, xs[1].Distance)
	assert.Equal(t, 5.5, xs[2].Distance)
	assert.Equal(t, 6.0, xs[3].Distance)
}

func Test_shading_intersection(t *testing.T) {
	w := defaultWorld()
	r := ddddf.Ray{geom.Point{X: 0, Y: 0, Z: -5}, geom.Vector{X: 0, Y: 0, Z: 1}}
	shape := w.Objects[0]
	i := ddddf.Inter{4, shape}
	comps := PrepareComputations(i, r, ddddf.Inters{i})

	c := w.ShadeHit(comps, MaxDepth)

	geom.AssertColorEqualInDelta(t, geom.Color{R: 0.38066, G: 0.47583, B: 0.2855}, c)
}

func Test_shading_intersection_from_inside(t *testing.T) {
	w := defaultWorld()
	w.Light = PointLight{geom.Point{X: 0, Y: 0.25, Z: 0}, geom.White}
	r := ddddf.Ray{geom.Point{X: 0, Y: 0, Z: 0}, geom.Vector{X: 0, Y: 0, Z: 1}}
	shape := w.Objects[1]
	i := ddddf.Inter{0.5, shape}
	comps := PrepareComputations(i, r, ddddf.Inters{i})

	c := w.ShadeHit(comps, MaxDepth)

	geom.AssertColorEqualInDelta(t, geom.Color{R: 0.90498, G: 0.90498, B: 0.90498}, c)
}

func Test_color_when_ray_misses(t *testing.T) {
	w := defaultWorld()
	r := ddddf.Ray{geom.Point{X: 0, Y: 0, Z: -5}, geom.Vector{X: 0, Y: 1, Z: 0}}

	c := w.ColorAt(r, MaxDepth)

	geom.AssertColorEqualInDelta(t, geom.Black, c)
}

func Test_color_when_ray_hits(t *testing.T) {
	w := defaultWorld()
	r := ddddf.Ray{geom.Point{X: 0, Y: 0, Z: -5}, geom.Vector{X: 0, Y: 0, Z: 1}}

	c := w.ColorAt(r, MaxDepth)

	geom.AssertColorEqualInDelta(t, geom.Color{R: 0.38066, G: 0.47583, B: 0.2855}, c)
}

func Test_color_with_intersection_behind_ray(t *testing.T) {
	w := World{pointLightSample(), []ddddf.Shape{
		ddddf.NewSphere(geom.IdentityMatrix(), testMaterialBuilder().SetAmbient(1).Build()),
		ddddf.NewSphere(geom.Scaling(0.5, 0.5, 0.5), testMaterialBuilder().SetAmbient(1).Build())}}
	r := ddddf.Ray{geom.Point{X: 0, Y: 0, Z: 0.75}, geom.Vector{X: 0, Y: 0, Z: -1}}

	c := w.ColorAt(r, MaxDepth)

	geom.AssertColorEqualInDelta(t, testMaterialBuilder().SetAmbient(1).Build().Color, c)
}

func Test_shade_hit_is_given_intersection_in_shadow(t *testing.T) {
	s1 := ddddf.NewSphere(geom.IdentityMatrix(), asdf.DefaultMaterial())
	s2 := ddddf.NewSphere(geom.Translation(0, 0, 10), asdf.DefaultMaterial())
	w := World{
		PointLight{geom.Point{X: 0, Y: 0, Z: -10}, geom.White},
		[]ddddf.Shape{s1, s2},
	}
	r := ddddf.Ray{geom.Point{X: 0, Y: 0, Z: 5}, geom.Vector{X: 0, Y: 0, Z: 1}}
	i := ddddf.Inter{4, s2}
	comps := PrepareComputations(i, r, ddddf.Inters{i})

	color := w.ShadeHit(comps, MaxDepth)

	assert.Equal(t, geom.Color{R: 0.1, G: 0.1, B: 0.1}, color)
}

func Test_hit_should_offset_point(t *testing.T) {
	r := ddddf.Ray{geom.Point{X: 0, Y: 0, Z: -5}, geom.Vector{X: 0, Y: 0, Z: 1}}
	s := ddddf.NewSphere(geom.Translation(0, 0, 1), asdf.DefaultMaterial())
	i := ddddf.Inter{5, s}

	comps := PrepareComputations(i, r, ddddf.Inters{i})

	assert.Less(t, comps.OverPoint.Z, -geom.Delta/2)
	assert.Less(t, comps.OverPoint.Z, comps.Point.Z)
}

func Test_reflected_color_for_non_reflective_material(t *testing.T) {
	s1 := ddddf.NewSphere(geom.IdentityMatrix(), testMaterialBuilder().Build())
	s2 := ddddf.NewSphere(geom.Scaling(0.5, 0.5, 0.5), testMaterialBuilder().SetAmbient(1).Build())
	w := World{pointLightSample(), []ddddf.Shape{s1, s2}}
	r := ddddf.Ray{geom.Point{X: 0, Y: 0, Z: 0}, geom.Vector{X: 0, Y: 0, Z: 1}}
	i := ddddf.Inter{1, s2}
	comps := PrepareComputations(i, r, ddddf.Inters{i})

	color := w.ReflectedColor(comps, 5)

	assert.Equal(t, geom.Color{R: 0, G: 0, B: 0}, color)
}

func Test_reflected_color_for_reflective_material(t *testing.T) {
	s1 := ddddf.NewSphere(geom.IdentityMatrix(), testMaterialBuilder().Build())
	s2 := ddddf.NewSphere(geom.Scaling(0.5, 0.5, 0.5), testMaterialBuilder().SetAmbient(1).Build())
	s3 := ddddf.NewPlane(geom.Translation(0, -1, 0), asdf.MakeMaterialBuilder().SetReflective(0.5).Build())
	w := World{pointLightSample(), []ddddf.Shape{s1, s2, s3}}
	r := ddddf.Ray{geom.Point{X: 0, Y: 0, Z: -3}, geom.Vector{X: 0, Y: -math.Sqrt2 / 2, Z: math.Sqrt2 / 2}}
	i := ddddf.Inter{math.Sqrt2, s3}
	comps := PrepareComputations(i, r, ddddf.Inters{i})

	color := w.ReflectedColor(comps, 5)

	geom.AssertColorEqualInDelta(t, geom.Color{R: 0.19033, G: 0.23791, B: 0.142749}, color)
}

func Test_shade_hit_with_reflective_material(t *testing.T) {
	s1 := ddddf.NewSphere(geom.IdentityMatrix(), testMaterialBuilder().Build())
	s2 := ddddf.NewSphere(geom.Scaling(0.5, 0.5, 0.5), testMaterialBuilder().SetAmbient(1).Build())
	s3 := ddddf.NewPlane(geom.Translation(0, -1, 0), asdf.MakeMaterialBuilder().SetReflective(0.5).Build())
	w := World{pointLightSample(), []ddddf.Shape{s1, s2, s3}}
	r := ddddf.Ray{geom.Point{X: 0, Y: 0, Z: -3}, geom.Vector{X: 0, Y: -math.Sqrt2 / 2, Z: math.Sqrt2 / 2}}
	i := ddddf.Inter{math.Sqrt2, s3}
	comps := PrepareComputations(i, r, ddddf.Inters{i})

	color := w.ShadeHit(comps, MaxDepth)

	geom.AssertColorEqualInDelta(t, geom.Color{R: 0.87675, G: 0.92434, B: 0.82918}, color)
}

func Test_color_at_with_mutually_reflective_surfaces(t *testing.T) {
	w := World{
		PointLight{geom.Point{X: 0, Y: 0, Z: 0}, geom.Color{R: 1, G: 1, B: 1}},
		[]ddddf.Shape{
			ddddf.NewPlane(geom.Translation(0, -1, 0), asdf.MakeMaterialBuilder().SetReflective(1).Build()),
			ddddf.NewPlane(geom.Translation(0, 1, 0), asdf.MakeMaterialBuilder().SetReflective(1).Build())}}

	w.ColorAt(ddddf.Ray{geom.Point{X: 0, Y: 0, Z: 0}, geom.Vector{X: 0, Y: 1, Z: 0}}, MaxDepth)

	//should terminate
}

func Test_reflected_color_at_maximum_recursive_depth(t *testing.T) {
	s1 := ddddf.NewSphere(geom.IdentityMatrix(), testMaterialBuilder().Build())
	s2 := ddddf.NewSphere(geom.Scaling(0.5, 0.5, 0.5), testMaterialBuilder().SetAmbient(1).Build())
	s3 := ddddf.NewPlane(geom.Translation(0, -1, 0), asdf.MakeMaterialBuilder().SetReflective(0.5).Build())
	w := World{pointLightSample(), []ddddf.Shape{s1, s2, s3}}
	r := ddddf.Ray{geom.Point{X: 0, Y: 0, Z: -3}, geom.Vector{X: 0, Y: -math.Sqrt2 / 2, Z: math.Sqrt2 / 2}}
	i := ddddf.Inter{math.Sqrt2, s3}
	comps := PrepareComputations(i, r, ddddf.Inters{i})

	color := w.ReflectedColor(comps, 0)

	geom.AssertColorEqualInDelta(t, geom.Color{R: 0, G: 0, B: 0}, color)
}

func Test_refracted_color_with_opaque_surface(t *testing.T) {
	w := defaultWorld()
	s := w.Objects[0]
	r := ddddf.Ray{geom.Point{X: 0, Y: 0, Z: -5}, geom.Vector{X: 0, Y: 0, Z: 1}}
	xs := ddddf.Inters{ddddf.Inter{4, s}, ddddf.Inter{6, s}}
	comps := PrepareComputations(xs[0], r, xs)

	c := w.RefractedColor(comps, MaxDepth)

	assert.Equal(t, geom.Color{R: 0, G: 0, B: 0}, c)
}

func Test_refracted_color_at_the_maximum_recursive_depth(t *testing.T) {
	s1 := ddddf.NewSphere(geom.IdentityMatrix(), testMaterialBuilder().SetTransparency(1.0).SetRefractiveIndex(1.5).Build())
	s2 := ddddf.NewSphere(geom.Scaling(0.5, 0.5, 0.5), asdf.DefaultMaterial())
	w := World{pointLightSample(), []ddddf.Shape{s1, s2}}
	r := ddddf.Ray{geom.Point{X: 0, Y: 0, Z: -5}, geom.Vector{X: 0, Y: 0, Z: 1}}
	xs := ddddf.Inters{ddddf.Inter{4, s1}, ddddf.Inter{6, s1}}
	comps := PrepareComputations(xs[0], r, xs)

	c := w.RefractedColor(comps, 0)

	assert.Equal(t, geom.Color{R: 0, G: 0, B: 0}, c)
}

func Test_refracted_color_under_total_internal_reflection(t *testing.T) {
	s1 := ddddf.NewSphere(geom.IdentityMatrix(), testMaterialBuilder().SetTransparency(1.0).SetRefractiveIndex(1.5).Build())
	s2 := ddddf.NewSphere(geom.Scaling(0.5, 0.5, 0.5), asdf.DefaultMaterial())
	w := World{pointLightSample(), []ddddf.Shape{s1, s2}}
	r := ddddf.Ray{geom.Point{X: 0, Y: 0, Z: math.Sqrt2 / 2}, geom.Vector{X: 0, Y: 1, Z: 0}}
	xs := ddddf.Inters{ddddf.Inter{-math.Sqrt2 / 2, s1}, ddddf.Inter{math.Sqrt2 / 2, s1}}
	comps := PrepareComputations(xs[1], r, xs)

	c := w.RefractedColor(comps, MaxDepth)

	assert.Equal(t, geom.Color{R: 0, G: 0, B: 0}, c)
}

func Test_refracted_color_with_refracted_ray(t *testing.T) {
	s1 := ddddf.NewSphere(geom.IdentityMatrix(),
		testMaterialBuilder().
			SetAmbient(1.0).
			SetPattern(TestPattern{}).
			Build())
	s2 := ddddf.NewSphere(
		geom.Scaling(0.5, 0.5, 0.5),
		testMaterialBuilder().
			SetTransparency(1.0).
			SetRefractiveIndex(1.5).
			Build())
	w := World{pointLightSample(), []ddddf.Shape{s1, s2}}
	r := ddddf.Ray{geom.Point{X: 0, Y: 0, Z: 0.1}, geom.Vector{X: 0, Y: 1, Z: 0}}
	xs := ddddf.Inters{ddddf.Inter{-0.9899, s1}, ddddf.Inter{-0.4899, s2}, ddddf.Inter{0.4899, s2}, ddddf.Inter{0.9899, s1}}
	comps := PrepareComputations(xs[2], r, xs)

	c := w.RefractedColor(comps, MaxDepth)

	geom.AssertColorEqualInDelta(t, geom.Color{R: 0, G: 0.99888, B: 0.04721}, c)
}

func Test_shade_hit_with_transparent_material(t *testing.T) {
	s1 := ddddf.NewSphere(geom.IdentityMatrix(), testMaterialBuilder().Build())
	s2 := ddddf.NewSphere(geom.Scaling(0.5, 0.5, 0.5), asdf.DefaultMaterial())
	floor := ddddf.NewPlane(geom.Translation(0, -1, 0),
		asdf.MakeMaterialBuilder().
			SetTransparency(0.5).
			SetRefractiveIndex(1.5).
			Build())
	ball := ddddf.NewSphere(geom.Translation(0, -3.5, -0.5),
		asdf.MakeMaterialBuilder().
			SetColor(geom.Color{R: 1, G: 0, B: 0}).
			SetAmbient(0.5).
			Build())
	w := World{pointLightSample(), []ddddf.Shape{s1, s2, floor, ball}}
	r := ddddf.Ray{geom.Point{X: 0, Y: 0, Z: -3}, geom.Vector{X: 0, Y: -math.Sqrt2 / 2, Z: math.Sqrt2 / 2}}
	xs := ddddf.Inters{ddddf.Inter{math.Sqrt2, floor}}
	comps := PrepareComputations(xs[0], r, xs)

	color := w.ShadeHit(comps, MaxDepth)

	geom.AssertColorEqualInDelta(t, geom.Color{R: 0.93642, G: 0.68642, B: 0.68642}, color)
}

func Test_shade_hit_with_reflective_transparent_material(t *testing.T) {
	s1 := ddddf.NewSphere(geom.IdentityMatrix(), testMaterialBuilder().Build())
	s2 := ddddf.NewSphere(geom.Scaling(0.5, 0.5, 0.5), asdf.DefaultMaterial())
	floor := ddddf.NewPlane(geom.Translation(0, -1, 0),
		asdf.MakeMaterialBuilder().
			SetReflective(0.5).
			SetTransparency(0.5).
			SetRefractiveIndex(1.5).
			Build())
	ball := ddddf.NewSphere(geom.Translation(0, -3.5, -0.5),
		asdf.MakeMaterialBuilder().
			SetColor(geom.Color{R: 1, G: 0, B: 0}).
			SetAmbient(0.5).
			Build())
	w := World{pointLightSample(), []ddddf.Shape{s1, s2, floor, ball}}
	r := ddddf.Ray{geom.Point{X: 0, Y: 0, Z: -3}, geom.Vector{X: 0, Y: -math.Sqrt2 / 2, Z: math.Sqrt2 / 2}}
	xs := ddddf.Inters{ddddf.Inter{math.Sqrt2, floor}}
	comps := PrepareComputations(xs[0], r, xs)

	color := w.ShadeHit(comps, MaxDepth)

	geom.AssertColorEqualInDelta(t, geom.Color{R: 0.93391, G: 0.69643, B: 0.69243}, color)
}

func Test_lighting(t *testing.T) {
	tests := []struct {
		name     string
		eyev     geom.Vector
		normalv  geom.Vector
		light    PointLight
		expected geom.Color
	}{
		{"Lighting with the eye between the light and the surface",
			geom.Vector{X: 0, Y: 0, Z: -1},
			geom.Vector{X: 0, Y: 0, Z: -1},
			PointLight{geom.Point{X: 0, Y: 0, Z: -10}, geom.White},
			geom.Color{R: 1.9, G: 1.9, B: 1.9}},
		{"Lighting with the eye between light and surface, eye offset 45°",
			geom.Vector{X: 0, Y: math.Sqrt2 / 2, Z: -math.Sqrt2 / 2},
			geom.Vector{X: 0, Y: 0, Z: -1},
			PointLight{geom.Point{X: 0, Y: 0, Z: -10}, geom.White},
			geom.White},
		{"Lighting with eye opposite surface, light offset 45°",
			geom.Vector{X: 0, Y: 0, Z: -1},
			geom.Vector{X: 0, Y: 0, Z: -1},
			PointLight{geom.Point{X: 0, Y: 10, Z: -10}, geom.White},
			geom.Color{R: 0.7364, G: 0.7364, B: 0.7364}},
		{"Lighting with eye in the path of the reflection vector",
			geom.Vector{X: 0, Y: -math.Sqrt2 / 2, Z: -math.Sqrt2 / 2},
			geom.Vector{X: 0, Y: 0, Z: -1},
			PointLight{geom.Point{X: 0, Y: 10, Z: -10}, geom.White},
			geom.Color{R: 1.6364, G: 1.6364, B: 1.6364}},
		{"Lighting with the light behind the surface",
			geom.Vector{X: 0, Y: 0, Z: -1},
			geom.Vector{X: 0, Y: 0, Z: -1},
			PointLight{geom.Point{X: 0, Y: 0, Z: 10}, geom.White},
			geom.Color{R: 0.1, G: 0.1, B: 0.1}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			color := Lighting(asdf.DefaultMaterial(), ddddf.NewSphere(geom.IdentityMatrix(), asdf.DefaultMaterial()), test.light, geom.Point{}, test.eyev, test.normalv, false)

			geom.AssertColorEqualInDelta(t, test.expected, color)
		})
	}
}

func Test_lighting_with_surface_in_shadow(t *testing.T) {
	m := asdf.DefaultMaterial()
	eyeV := geom.Vector{X: 0, Y: 0, Z: -1}
	normalV := geom.Vector{X: 0, Y: 0, Z: -1}
	light := PointLight{geom.Point{X: 0, Y: 0, Z: -10}, geom.White}

	r := Lighting(m, ddddf.NewSphere(geom.IdentityMatrix(), asdf.DefaultMaterial()), light, geom.Point{}, eyeV, normalV, true)

	assert.Equal(t, geom.Color{R: 0.1, G: 0.1, B: 0.1}, r)
}

func Test_shadow(t *testing.T) {
	tests := []struct {
		name     string
		point    geom.Point
		expected bool
	}{
		{"There is no shadow when nothing is collinear with point and light",
			geom.Point{X: 0, Y: 10, Z: 0}, false},
		{"The shadow when an object is between the point and the light",
			geom.Point{X: 10, Y: -10, Z: 10}, true},
		{"There is no shadow when an object is behind the light",
			geom.Point{X: -20, Y: 20, Z: -20}, false},
		{"There is no shadow when an object is behind the point",
			geom.Point{X: -2, Y: 2, Z: -2}, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := defaultWorld()

			r := w.IsShadowed(test.point)

			assert.Equal(t, test.expected, r)
		})
	}
}

func Test_Lighting_with_pattern_applied(t *testing.T) {
	m := asdf.MakeMaterialBuilder().
		SetAmbient(1).
		SetDiffuse(0).
		SetSpecular(0).
		SetPattern(asdf.MakeStripePattern(geom.White, geom.Black)).
		Build()

	eyeV := geom.Vector{X: 0, Y: 0, Z: -1}
	normalV := geom.Vector{X: 0, Y: 0, Z: -1}
	light := PointLight{geom.Point{X: 0, Y: 0, Z: -10}, geom.White}
	c1 := Lighting(m, ddddf.NewSphere(geom.IdentityMatrix(), asdf.DefaultMaterial()), light, geom.Point{X: 0.9, Y: 0, Z: 0}, eyeV, normalV, false)
	c2 := Lighting(m, ddddf.NewSphere(geom.IdentityMatrix(), asdf.DefaultMaterial()), light, geom.Point{X: 1.1, Y: 0, Z: 0}, eyeV, normalV, false)

	assert.Equal(t, geom.White, c1)
	assert.Equal(t, geom.Black, c2)
}

//util
func defaultWorld() World {
	s1 := ddddf.NewSphere(geom.IdentityMatrix(), testMaterialBuilder().Build())
	s2 := ddddf.NewSphere(geom.Scaling(0.5, 0.5, 0.5), asdf.DefaultMaterial())

	return World{
		pointLightSample(),
		[]ddddf.Shape{s1, s2},
	}
}

func pointLightSample() PointLight {
	return PointLight{geom.Point{X: -10, Y: 10, Z: -10}, geom.White}
}

func testMaterialBuilder() *asdf.MaterialBuilder {
	return asdf.MakeMaterialBuilder().
		SetColor(geom.Color{R: 0.8, G: 1.0, B: 0.6}).
		SetDiffuse(0.7).
		SetSpecular(0.2)
}

type TestPattern struct {
}

func (t TestPattern) PatternAt(point geom.Point) geom.Color {
	return geom.Color{R: point.X, G: point.Y, B: point.Z}
}
func (t TestPattern) Transformation() *geom.Matrix {
	return geom.IdentityMatrix()
}
