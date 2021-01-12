package figure

import (
	"github.com/oleg/raytracer-go/geom"
	"math"
	"sort"
)

const MaxDepth = 4

type World struct {
	Light   PointLight
	Objects []Shape
}

func (w *World) ColorAt(r Ray, remaining uint8) geom.Color {
	xs := w.Intersect(r)
	if ok, hit := xs.Hit(); ok {
		return w.ShadeHit(hit.PrepareComputations(r, xs), remaining)
	}
	return geom.Black
}

func (w *World) Intersect(ray Ray) Inters {
	r := make([]Inter, 0, 10)
	for _, shape := range w.Objects {
		r = append(r, Intersect(shape, ray)...)
	}
	sort.Slice(r, func(i, j int) bool {
		return r[i].Distance < r[j].Distance
	})
	return r
}

func (w *World) ShadeHit(comps Computations, remaining uint8) geom.Color {
	shadowed := w.IsShadowed(comps.OverPoint)
	surface := Lighting(
		comps.Object.Material(),
		comps.Object,
		w.Light,
		comps.OverPoint,
		comps.EyeV,
		comps.NormalV,
		shadowed)
	reflected := w.ReflectedColor(comps, remaining)
	refracted := w.RefractedColor(comps, remaining)
	material := comps.Object.Material()
	if material.Reflective > 0 && material.Transparency > 0 {
		reflectance := Schlick(comps)
		return surface.
			Add(reflected.MultiplyByScalar(reflectance)).
			Add(refracted.MultiplyByScalar(1 - reflectance))
	}
	return surface.
		Add(reflected).
		Add(refracted)
}

func (w *World) IsShadowed(point geom.Point) bool {
	v := w.Light.Position.SubtractPoint(point)
	distance := v.Magnitude()
	direction := v.Normalize()
	intersections := w.Intersect(Ray{point, direction})
	hit, inter := intersections.Hit()
	return hit && inter.Distance < distance
}

func (w *World) ReflectedColor(comps Computations, remaining uint8) geom.Color {
	if remaining <= 0 {
		return geom.Black
	}
	reflective := comps.Object.Material().Reflective
	if reflective == 0 {
		return geom.Black
	}
	reflectRay := Ray{comps.OverPoint, comps.ReflectV}
	color := w.ColorAt(reflectRay, remaining-1)
	return color.MultiplyByScalar(reflective)
}

func (w *World) RefractedColor(comps Computations, remaining uint8) geom.Color {
	if remaining <= 0 {
		return geom.Black
	}
	transparency := comps.Object.Material().Transparency
	if transparency == 0 {
		return geom.Black
	}
	nRatio := comps.N1 / comps.N2
	cosI := comps.EyeV.Dot(comps.NormalV)
	sin2t := math.Pow(nRatio, 2) * (1 - math.Pow(cosI, 2))
	if sin2t > 1 {
		return geom.Black
	}

	cosT := math.Sqrt(1.0 - sin2t)
	direction := comps.NormalV.MultiplyScalar(nRatio*cosI - cosT).
		SubtractVector(comps.EyeV.MultiplyScalar(nRatio))

	refractRay := Ray{comps.UnderPoint, direction}
	return w.ColorAt(refractRay, remaining-1).MultiplyByScalar(transparency)
}
