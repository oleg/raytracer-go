package scene

import (
	"github.com/oleg/raytracer-go/geom"
	"github.com/oleg/raytracer-go/shapes"
	"math"
)

type Camera struct {
	HSize, VSize          int
	HalfWidth, HalfHeight float64
	FieldOfView           float64
	PixelSize             float64
	Transform             *geom.Matrix
}

func NewCamera(hSize, vSize int, fieldOfView float64, transform *geom.Matrix) *Camera {
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

func (camera *Camera) RayForPixel(x, y int) shapes.Ray {
	xOffset := (float64(x) + 0.5) * camera.PixelSize
	yOffset := (float64(y) + 0.5) * camera.PixelSize

	worldX := camera.HalfWidth - xOffset
	worldY := camera.HalfHeight - yOffset

	pixel := camera.Transform.Inverse().MultiplyPoint(geom.Point{X: worldX, Y: worldY, Z: -1})
	origin := camera.Transform.Inverse().MultiplyPoint(geom.Point{X: 0, Y: 0, Z: 0})
	direction := pixel.SubtractPoint(origin).Normalize()
	return shapes.Ray{Origin: origin, Direction: direction}
}

func (camera *Camera) Render(w World) *Canvas {
	canvas := NewCanvas(camera.HSize, camera.VSize)
	for y := 0; y < camera.VSize; y++ {
		for x := 0; x < camera.HSize; x++ {
			ray := camera.RayForPixel(x, y)
			color := w.ColorAt(ray, MaxDepth)
			canvas.Pixels[x][y] = color
		}
	}
	return canvas
}