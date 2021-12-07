package geom

//go:generate go run ./gen-matrix/generator.go
import (
	"log"
	"strconv"
	"strings"
)

//todo implement HasTransformation?
type Matrix struct {
	data    matrix4x4 //todo make private
	inverse *Matrix
}

var identity4x4 = matrix4x4{
	{1, 0, 0, 0},
	{0, 1, 0, 0},
	{0, 0, 1, 0},
	{0, 0, 0, 1},
}

var identityMatrixRef = NewMatrix(identity4x4).Inverse()

func IdentityMatrix() *Matrix {
	return identityMatrixRef
}

func NewMatrix(data [4][4]float64) *Matrix {
	return &Matrix{data: data}
}

func ParseMatrix(str string) *Matrix {
	m := &Matrix{}
	rows := strings.Split(str, "\n")
	if len(rows) != 4 {
		log.Fatal("must have 4 rows")
	}
	for i, row := range rows {
		columns := strings.Split(row, "|")
		if len(columns) != 6 {
			log.Fatal("must have 6 separators for 4 columns")
		}
		for j, s := range columns[1:5] {
			m.data[i][j] = trimAndParseFloat(s)
		}
	}
	return m
}

func trimAndParseFloat(s string) float64 {
	s = strings.Trim(s, " ")
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	return val
}

func (m *Matrix) Data() [4][4]float64 { //todo oleg test it's copy
	return m.data
}

func (m *Matrix) Multiply(o *Matrix) *Matrix {
	result := &Matrix{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				result.data[i][j] += m.data[i][k] * o.data[k][j]
			}
		}
	}
	return result
}

func (m *Matrix) MultiplyPoint(t Point) Point {
	x, y, z := m.multiply(t.X, t.Y, t.Z, 1.)
	return Point{X: x, Y: y, Z: z}
}

func (m *Matrix) MultiplyVector(t Vector) Vector {
	x, y, z := m.multiply(t.X, t.Y, t.Z, 0.)
	return Vector{X: x, Y: y, Z: z}
}

func (m *Matrix) multiply(x, y, z, w float64) (float64, float64, float64) {
	d := m.data
	return d[0][0]*x + d[0][1]*y + d[0][2]*z + d[0][3]*w,
		d[1][0]*x + d[1][1]*y + d[1][2]*z + d[1][3]*w,
		d[2][0]*x + d[2][1]*y + d[2][2]*z + d[2][3]*w
}

func (m *Matrix) Transpose() *Matrix {
	//todo also cache?
	//todo or implement as loops?
	d := m.data
	return &Matrix{
		data: [4][4]float64{
			{d[0][0], d[1][0], d[2][0], d[3][0]},
			{d[0][1], d[1][1], d[2][1], d[3][1]},
			{d[0][2], d[1][2], d[2][2], d[3][2]},
			{d[0][3], d[1][3], d[2][3], d[3][3]},
		},
	}
}

func (m *Matrix) Inverse() *Matrix {
	if m.inverse != nil {
		return m.inverse
	}
	m.inverse = &Matrix{
		data:    *m.data.inverse(),
		inverse: m,
	}
	return m.inverse
}
