package scene

import (
	"github.com/oleg/raytracer-go/geom"
	"github.com/oleg/raytracer-go/shapes"
	"math"
	"sort"
)

type Computations struct {
	Distance   float64
	Object     shapes.Shape
	Point      geom.Point
	OverPoint  geom.Point
	UnderPoint geom.Point
	EyeV       geom.Vector
	NormalV    geom.Vector
	ReflectV   geom.Vector
	Inside     bool
	N1         float64
	N2         float64
}

func NewComputations(i shapes.Inter, r shapes.Ray, xs shapes.Inters) Computations {
	comps := Computations{}
	comps.Distance = i.Distance
	comps.Object = i.Object
	comps.Point = r.Position(i.Distance)
	comps.EyeV = r.Direction.Negate()

	normalV := NormalAt(comps.Object, comps.Point)
	comps.Inside = normalV.Dot(comps.EyeV) < 0
	if comps.Inside {
		comps.NormalV = normalV.Negate()
	} else {
		comps.NormalV = normalV
	}
	comps.ReflectV = r.Direction.Reflect(comps.NormalV)
	comps.OverPoint = comps.Point.AddVector(comps.NormalV.MultiplyScalar(geom.Delta))
	comps.UnderPoint = comps.Point.SubtractVector(comps.NormalV.MultiplyScalar(geom.Delta))
	comps.N1, comps.N2 = calcNs(i, xs)
	return comps
}

func NormalAt(shape shapes.Shape, worldPoint geom.Point) geom.Vector {
	localPoint := shape.Transformation().Inverse().MultiplyPoint(worldPoint)
	localNormal := shape.NormalAt(localPoint)
	worldNormal := shape.Transformation().Inverse().Transpose().MultiplyVector(localNormal)
	return worldNormal.Normalize()
}

func Schlick(comps Computations) float64 {
	cos := comps.EyeV.Dot(comps.NormalV)
	if comps.N1 > comps.N2 {
		n := comps.N1 / comps.N2
		sin2t := n * n * (1.0 - cos*cos)
		if sin2t > 1.0 {
			return 1.0
		}
		cos = math.Sqrt(1.0 - sin2t)
	}
	r0 := math.Pow((comps.N1-comps.N2)/(comps.N1+comps.N2), 2)
	return r0 + (1-r0)*math.Pow(1-cos, 5)
}

//todo does this method belongs here?
func calcNs(hit shapes.Inter, xs shapes.Inters) (n1 float64, n2 float64) {
	var shapes = make([]shapes.Shape, 0, 10)
	for _, i := range xs {
		if i == hit {
			n1 = refractiveIndex(shapes)
		}
		shapes = updateShapes(shapes, i.Object)
		if i == hit {
			n2 = refractiveIndex(shapes)
		}
	}
	return
}

func updateShapes(shapes []shapes.Shape, shape shapes.Shape) []shapes.Shape {
	if ok, pos := includes(shapes, shape); ok {
		return remove(shapes, pos)
	} else {
		return append(shapes, shape)
	}
}

func refractiveIndex(shapes []shapes.Shape) float64 {
	if len(shapes) != 0 {
		return shapes[len(shapes)-1].Material().RefractiveIndex
	}
	return 1.0
}

func includes(containers []shapes.Shape, object shapes.Shape) (bool, int) {
	for i, o := range containers {
		if o == object {
			return true, i
		}
	}
	return false, -1
}

func remove(s []shapes.Shape, i int) []shapes.Shape {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func Hit(xs shapes.Inters) (bool, shapes.Inter) {
	//todo move to constructor
	sort.Slice(xs, func(i, j int) bool {
		return xs[i].Distance < xs[j].Distance
	})
	for _, e := range xs {
		if e.Distance > 0 {
			return true, e
		}
	}
	return false, shapes.Inter{}
}
