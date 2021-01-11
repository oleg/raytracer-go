package figure

import (
	"github.com/oleg/raytracer-go/multid"
	"github.com/oleg/raytracer-go/oned"
	"math"
)

type Camera struct {
	HSize, VSize          int
	HalfWidth, HalfHeight float64
	FieldOfView           float64
	PixelSize             float64
	Transform             *multid.Matrix4
}

func NewCameraDefault(hSize, vSize int, fieldOfView float64) *Camera {
	return NewCamera(hSize, vSize, fieldOfView, multid.IdentityMatrix())
}

func NewCamera(hSize, vSize int, fieldOfView float64, transform *multid.Matrix4) *Camera {
	halfView := math.Tan(fieldOfView / 2.)
	aspect := float64(hSize) / float64(vSize)
	var halfWidth, halfHeight float64
	if aspect >= 1 {
		halfWidth, halfHeight = halfView, halfView/aspect
	} else {
		halfWidth, halfHeight = halfView*aspect, halfView
	}
	pixelSize := halfWidth * 2 / float64(hSize)
	return &Camera{
		hSize, vSize,
		halfWidth, halfHeight,
		fieldOfView, pixelSize, transform}
}

func (camera *Camera) RayForPixel(x, y int) Ray {
	xOffset := (float64(x) + 0.5) * camera.PixelSize
	yOffset := (float64(y) + 0.5) * camera.PixelSize

	worldX := camera.HalfWidth - xOffset
	worldY := camera.HalfHeight - yOffset

	pixel := camera.Transform.Inverse().MultiplyPoint(oned.Point{worldX, worldY, -1})
	origin := camera.Transform.Inverse().MultiplyPoint(oned.Point{0, 0, 0})
	direction := pixel.SubtractPoint(origin).Normalize()
	return Ray{origin, direction}
}

func (camera *Camera) Render(w World) *multid.Canvas {
	canvas := multid.NewCanvas(camera.HSize, camera.VSize)
	for y := 0; y < camera.VSize; y++ {
		for x := 0; x < camera.HSize; x++ {
			ray := camera.RayForPixel(x, y)
			color := w.ColorAt(ray, MaxDepth)
			canvas.Pixels[x][y] = color
		}
	}
	return canvas
}
