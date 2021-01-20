package config

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
