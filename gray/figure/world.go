package figure

import (
	"gray/oned"
	"math"
	"sort"
)

const MaxDepth = 4

type World struct {
	Light   PointLight
	Objects []Shape
}

func (w World) Intersect(ray Ray) Inters {
	r := Inters{}
	for _, shape := range w.Objects {
		r = append(r, Intersect(shape, ray)...)
	}
	sort.Slice(r, func(i, j int) bool {
		return r[i].Distance < r[j].Distance
	})
	return r
}

func (w World) ColorAt(r Ray, remaining uint8) oned.Color {
	xs := w.Intersect(r)
	if ok, hit := xs.Hit(); ok {
		return w.ShadeHit(hit.PrepareComputationsEx(r, xs), remaining)
	}
	return oned.Black
}

func (w World) ShadeHit(comps Computations, remaining uint8) oned.Color {
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

func (w World) IsShadowed(point oned.Point) bool {
	v := w.Light.Position.SubtractPoint(point)
	distance := v.Magnitude()
	direction := v.Normalize()
	intersections := w.Intersect(Ray{point, direction})
	hit, inter := intersections.Hit()
	return hit && inter.Distance < distance
}

func (w World) ReflectedColor(comps Computations, remaining uint8) oned.Color {
	if remaining <= 0 {
		return oned.Black
	}
	reflective := comps.Object.Material().Reflective
	if reflective == 0 {
		return oned.Black
	}
	reflectRay := Ray{comps.OverPoint, comps.ReflectV}
	color := w.ColorAt(reflectRay, remaining-1)
	return color.MultiplyByScalar(reflective)
}

func (w World) RefractedColor(comps Computations, remaining uint8) oned.Color {
	if remaining <= 0 {
		return oned.Black
	}
	transparency := comps.Object.Material().Transparency
	if transparency == 0 {
		return oned.Black
	}
	nRatio := comps.N1 / comps.N2
	cosI := comps.EyeV.Dot(comps.NormalV)
	sin2t := math.Pow(nRatio, 2) * (1 - math.Pow(cosI, 2))
	if sin2t > 1 {
		return oned.Black
	}

	cosT := math.Sqrt(1.0 - sin2t)
	direction := comps.NormalV.MultiplyScalar(nRatio*cosI - cosT).
		SubtractVector(comps.EyeV.MultiplyScalar(nRatio))

	refractRay := Ray{comps.UnderPoint, direction}
	return w.ColorAt(refractRay, remaining-1).MultiplyByScalar(transparency)
}
