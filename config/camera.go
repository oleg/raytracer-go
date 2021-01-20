package config

import (
	"github.com/oleg/raytracer-go/geom"
	"github.com/oleg/raytracer-go/scene"
	"gopkg.in/yaml.v3"
	"io"
	"io/ioutil"
)

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

type Tuple struct {
	X float64 `yaml:"x"`
	Y float64 `yaml:"y"`
	Z float64 `yaml:"z"`
}

func ReadCamera(r io.Reader) (*scene.Camera, error) {
	conf, err := readUnmarshalConfig(r)
	if err != nil {
		return nil, err
	}
	camera := scene.NewCamera(
		conf.Size.Horizontal,
		conf.Size.Vertical,
		conf.Size.FieldOfView,
		scene.Sight{
			From: asPoint(conf.From),
			To:   asPoint(conf.To),
			Up:   asVector(conf.Up),
		})
	return camera, nil
}

func readUnmarshalConfig(r io.Reader) (Camera, error) {
	all, err := ioutil.ReadAll(r)
	if err != nil {
		return Camera{}, err
	}
	var camera Camera
	err = yaml.Unmarshal(all, &camera)
	if err != nil {
		return Camera{}, err
	}
	return camera, nil
}

func asPoint(t Tuple) geom.Point {
	return geom.Point{
		X: t.X,
		Y: t.Y,
		Z: t.Z,
	}
}
func asVector(t Tuple) geom.Vector {
	return geom.Vector{
		X: t.X,
		Y: t.Y,
		Z: t.Z,
	}
}
