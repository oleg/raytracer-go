package scene

import (
	"github.com/oleg/raytracer-go/geom"
	"github.com/oleg/raytracer-go/physic"
	"github.com/oleg/raytracer-go/shapes"
	"math"
)

type Camera struct {
	HSize, VSize          int
	FieldOfView           float64
	HalfWidth, HalfHeight float64
	PixelSize             float64
	Transform             *geom.Matrix
}

func NewCamera(hSize, vSize int, fieldOfView float64, sight physic.HasTransformation) *Camera {
	halfWidth, halfHeight := calcHalfWidthAndHeight(hSize, vSize, fieldOfView)
	pixelSize := halfWidth * 2 / float64(hSize)
	return &Camera{hSize, vSize, fieldOfView, halfWidth, halfHeight, pixelSize, sight.Transformation()}
}

func calcHalfWidthAndHeight(hSize int, vSize int, fieldOfView float64) (float64, float64) {
	halfView := math.Tan(fieldOfView / 2.)
	aspect := float64(hSize) / float64(vSize)
	if aspect >= 1 {
		return halfView, halfView / aspect
	}
	return halfView * aspect, halfView
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
