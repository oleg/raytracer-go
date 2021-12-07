package config

import (
	"errors"
	"github.com/oleg/raytracer-go/geom"
	"github.com/oleg/raytracer-go/physic"
	"github.com/oleg/raytracer-go/scene"
	"github.com/oleg/raytracer-go/shapes"
	"gopkg.in/yaml.v3"
	"io"
	"io/ioutil"
)

type Scene struct {
	Camera `yaml:"camera"`
	World  `yaml:"world"`
}

type Camera struct {
	Size  `yaml:"size"`
	Sight `yaml:"sight"`
}

type Size struct {
	Horizontal  int     `yaml:"horizontal"`
	Vertical    int     `yaml:"vertical"`
	FieldOfView float64 `yaml:"field-of-view"`
}

type Sight struct {
	From Tuple `yaml:"from"`
	To   Tuple `yaml:"to"`
	Up   Tuple `yaml:"up"`
}

type World struct {
	Light   `yaml:"light"`
	Objects []Object `yaml:"objects"`
}

type Light struct {
	Position  Tuple `yaml:"position"`
	Intensity Color `yaml:"intensity"`
}

type Object struct {
	Type      string        `yaml:"type"`
	Transform [4][4]float64 `yaml:"transform"`
	Material  Material      `yaml:"material"`
}

type Material struct {
	Color           Color   `yaml:"color"`
	Ambient         float64 `yaml:"ambient"`
	Diffuse         float64 `yaml:"diffuse"`
	Specular        float64 `yaml:"specular"`
	Reflective      float64 `yaml:"reflective"`
	Transparency    float64 `yaml:"transparency"`
	RefractiveIndex float64 `yaml:"refractive-index"`
	Shininess       float64 `yaml:"shininess"`
	Pattern         Pattern `yaml:"pattern"`
}
type Pattern struct {
	Type      string        `yaml:"type"`
	A         Color         `yaml:"color-a"`
	B         Color         `yaml:"color-b"`
	Transform [4][4]float64 `yaml:"transform"`
}
type Tuple struct {
	X float64 `yaml:"x"`
	Y float64 `yaml:"y"`
	Z float64 `yaml:"z"`
}

type Color struct {
	R float64 `yaml:"r"`
	G float64 `yaml:"g"`
	B float64 `yaml:"b"`
}

func ReadScene(r io.Reader) (camera *scene.Camera, world scene.World, err error) {
	conf, err := unmarshalConf(r)
	if err != nil {
		return
	}
	objects, err := toObjects(conf.World.Objects)
	if err != nil {
		return
	}
	world = scene.World{
		Light:   toLight(conf.World.Light),
		Objects: objects,
	}
	camera = scene.NewCamera(
		conf.Size.Horizontal,
		conf.Size.Vertical,
		conf.Size.FieldOfView,
		scene.Sight{
			From: toPoint(conf.From),
			To:   toPoint(conf.To),
			Up:   toVector(conf.Up),
		})
	return
}

func unmarshalConf(r io.Reader) (Scene, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return Scene{}, err
	}
	var s Scene
	err = yaml.Unmarshal(data, &s)
	if err != nil {
		return Scene{}, err
	}
	return s, nil
}

func toLight(conf Light) scene.PointLight {
	return scene.PointLight{
		Position:  toPoint(conf.Position),
		Intensity: toColor(conf.Intensity),
	}
}

func toObjects(objects []Object) ([]shapes.Shape, error) {
	objets := make([]shapes.Shape, 0, len(objects))
	for _, v := range objects {
		object, err := parseObject(v)
		if err != nil {
			return nil, err
		}
		objets = append(objets, object)
	}
	return objets, nil
}

func parseObject(v Object) (shapes.Shape, error) {
	tr := &geom.Matrix{Data: v.Transform}
	mt, err := toMaterial(v)
	if err != nil {
		return nil, err
	}
	switch v.Type {
	case "Plane":
		return shapes.NewPlane(tr, mt), nil
	case "Sphere":
		return shapes.NewSphere(tr, mt), nil
	default:
		return nil, errors.New("unknown type " + v.Type)
	}
}

func toMaterial(v Object) (*physic.Material, error) {
	patter, err := toPatter(v.Material.Pattern)
	if err != nil {
		return nil, err
	}
	material := &physic.Material{
		Color:           toColor(v.Material.Color),
		Ambient:         v.Material.Ambient,
		Diffuse:         v.Material.Diffuse,
		Specular:        v.Material.Specular,
		Shininess:       v.Material.Shininess,
		Reflective:      v.Material.Reflective,
		Transparency:    v.Material.Transparency,
		RefractiveIndex: v.Material.RefractiveIndex,
		Pattern:         patter,
	}
	return material, nil
}

func toPatter(pattern Pattern) (physic.Pattern, error) {
	if pattern.Type == "" {
		return nil, nil
	}
	switch pattern.Type {
	case "Checkers":
		checkersPattern := physic.CheckersPattern{
			A:             toColor(pattern.A),
			B:             toColor(pattern.B),
			Transformable: physic.Transformable{Rule: &geom.Matrix{Data: pattern.Transform}},
		}
		return checkersPattern, nil
	default:
		return nil, errors.New("Unknown pattern type " + pattern.Type)
	}
}

func toPoint(t Tuple) geom.Point {
	return geom.Point{
		X: t.X,
		Y: t.Y,
		Z: t.Z,
	}
}
func toVector(t Tuple) geom.Vector {
	return geom.Vector{
		X: t.X,
		Y: t.Y,
		Z: t.Z,
	}
}

func toColor(color Color) geom.Color {
	return geom.Color{
		R: color.R,
		G: color.G,
		B: color.B,
	}
}
