package figure

import "github.com/oleg/raytracer-go/geom"

type PointLight struct {
	Position  geom.Point
	Intensity geom.Color
}
