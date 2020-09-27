package figure

import (
	"github.com/stretchr/testify/assert"
	"gray/multid"
	"gray/oned"
	"math"
	"testing"
)

func Test_constructing_camera(t *testing.T) {
	hsize := 160
	vsize := 120
	fieldOfView := math.Pi / 2

	c := MakeCameraD(hsize, vsize, fieldOfView)

	assert.Equal(t, 160, c.HSize)
	assert.Equal(t, 120, c.VSize)
	assert.Equal(t, math.Pi/2, c.FieldOfView)
	assert.Equal(t, multid.IdentityMatrix, c.Transform)
}

func Test_pixel_size_for_horizontal_canvas(t *testing.T) {
	c := MakeCameraD(200, 125, math.Pi/2)

	assert.Equal(t, 0.01, c.PixelSize)
}

func Test_pixel_size_for_vertical_canvas(t *testing.T) {
	c := MakeCameraD(125, 200, math.Pi/2)

	assert.Equal(t, 0.01, c.PixelSize)
}

func Test_constructing_ray_with_camera(t *testing.T) {
	tests := []struct {
		name     string
		camera   Camera
		x, y     int
		expected Ray
	}{
		{"Constructing a ray through the center of the canvas",
			MakeCameraD(201, 101, math.Pi/2),
			100, 50,
			Ray{oned.Point{0, 0, 0}, oned.Vector{0, 0, -1}},
		},
		{"Constructing a ray through a corner of the canvas",
			MakeCameraD(201, 101, math.Pi/2),
			0, 0,
			Ray{oned.Point{0, 0, 0}, oned.Vector{0.66519, 0.33259, -0.66851}},
		},
		{"Constructing a ray when the camera is transformed",
			MakeCamera(201, 101, math.Pi/2, multid.RotationY(math.Pi/4).Multiply(multid.Translation(0, -2, 5))),
			100, 50,
			Ray{oned.Point{0, 2, -5}, oned.Vector{math.Sqrt2 / 2, 0, -math.Sqrt2 / 2}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := test.camera.RayForPixel(test.x, test.y)

			multid.AssertPointEqualInDelta(t, test.expected.Origin, r.Origin)
			oned.AssertVectorEqualInDelta(t, test.expected.Direction, r.Direction)
		})
	}
}

func Test_rendering_world_with_camera(t *testing.T) {
	w := defaultWorld()
	from := oned.Point{0, 0, -5}
	to := oned.Point{0, 0, 0}
	up := oned.Vector{0, 1, 0}
	c := MakeCamera(11, 11, math.Pi/2, ViewTransform(from, to, up))

	image := c.Render(w)

	oned.AssertColorEqualInDelta(t, oned.Color{0.38066, 0.47583, 0.2855}, image.Pixels[5][5])
}
