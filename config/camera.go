package config

type Camera struct {
	Size        CameraSize `yaml:"size"`
	FieldOfView float64    `yaml:"field-of-view"`
	//transform *geom.Matrix
}

type CameraSize struct {
	Horizontal int `yaml:"horizontal"`
	Vertical   int `yaml:"vertical"`
}
