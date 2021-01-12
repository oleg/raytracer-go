package scene

import (
	"github.com/oleg/raytracer-go/shapes"
	"github.com/oleg/raytracer-go/geom"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_constructing_camera(t *testing.T) {
	hsize := 160
	vsize := 120
	fieldOfView := math.Pi / 2

	c := NewCameraDefault(hsize, vsize, fieldOfView)

	assert.Equal(t, 160, c.HSize)
	assert.Equal(t, 120, c.VSize)
	assert.Equal(t, math.Pi/2, c.FieldOfView)
	assert.Equal(t, geom.IdentityMatrix(), c.Transform)
}

func Test_pixel_size_for_horizontal_canvas(t *testing.T) {
	c := NewCameraDefault(200, 125, math.Pi/2)

	assert.Equal(t, 0.01, c.PixelSize)
}

func Test_pixel_size_for_vertical_canvas(t *testing.T) {
	c := NewCameraDefault(125, 200, math.Pi/2)

	assert.Equal(t, 0.01, c.PixelSize)
}

func Test_constructing_ray_with_camera(t *testing.T) {
	tests := []struct {
		name     string
		camera   *Camera
		x, y     int
		expected shapes.Ray
	}{
		{"Constructing a ray through the center of the canvas",
			NewCameraDefault(201, 101, math.Pi/2),
			100, 50,
			shapes.Ray{Origin: geom.Point{X: 0, Y: 0, Z: 0}, Direction: geom.Vector{X: 0, Y: 0, Z: -1}},
		},
		{"Constructing a ray through a corner of the canvas",
			NewCameraDefault(201, 101, math.Pi/2),
			0, 0,
			shapes.Ray{Origin: geom.Point{X: 0, Y: 0, Z: 0}, Direction: geom.Vector{X: 0.66519, Y: 0.33259, Z: -0.66851}},
		},
		{"Constructing a ray when the camera is transformed",
			NewCamera(201, 101, math.Pi/2, geom.RotationY(math.Pi/4).Multiply(geom.Translation(0, -2, 5))),
			100, 50,
			shapes.Ray{Origin: geom.Point{X: 0, Y: 2, Z: -5}, Direction: geom.Vector{X: math.Sqrt2 / 2, Y: 0, Z: -math.Sqrt2 / 2}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := test.camera.RayForPixel(test.x, test.y)

			geom.AssertPointEqualInDelta(t, test.expected.Origin, r.Origin)
			geom.AssertVectorEqualInDelta(t, test.expected.Direction, r.Direction)
		})
	}
}

func Test_rendering_world_with_camera(t *testing.T) {
	w := defaultWorld()
	from := geom.Point{X: 0, Y: 0, Z: -5}
	to := geom.Point{X: 0, Y: 0, Z: 0}
	up := geom.Vector{X: 0, Y: 1, Z: 0}
	c := NewCamera(11, 11, math.Pi/2, ViewTransform(from, to, up))

	image := c.Render(w)

	geom.AssertColorEqualInDelta(t, geom.Color{R: 0.38066, G: 0.47583, B: 0.2855}, image.Pixels[5][5])
}
