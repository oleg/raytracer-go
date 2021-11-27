package geom

//go:generate go run ./gen-matrix/generator.go
import (
	"log"
	"strconv"
	"strings"
)

//todo: decide if I want to return a pointers or structs?
//todo add iterate function that accept function?
//todo implement HasTransformation?
type Matrix struct {
	Data    matrix4x4
	inverse *Matrix
}

var identityMatrixData = matrix4x4{
	{1, 0, 0, 0},
	{0, 1, 0, 0},
	{0, 0, 1, 0},
	{0, 0, 0, 1},
}

var identityMatrixRef = createInvertedMatrix()

func createInvertedMatrix() *Matrix {
	m := &Matrix{Data: identityMatrixData}
	m.inverse = m
	return m
}

func IdentityMatrix() *Matrix {
	return identityMatrixRef
}

//todo must?
func NewMatrix(str string) *Matrix {
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
			m.Data[i][j] = trimAndParseFloat(s)
		}
	}
	return m
}

func (m *Matrix) Multiply(o *Matrix) *Matrix {
	res := &Matrix{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				res.Data[i][j] += m.Data[i][k] * o.Data[k][j]
			}
		}
	}
	return res
}

func (m *Matrix) MultiplyPoint(t Point) Point {
	d := m.Data
	return Point{
		X: d[0][0]*t.X + d[0][1]*t.Y + d[0][2]*t.Z + d[0][3]*1.,
		Y: d[1][0]*t.X + d[1][1]*t.Y + d[1][2]*t.Z + d[1][3]*1.,
		Z: d[2][0]*t.X + d[2][1]*t.Y + d[2][2]*t.Z + d[2][3]*1.,
	}
}

//todo: remove duplication
func (m *Matrix) MultiplyVector(t Vector) Vector {
	d := m.Data
	return Vector{
		X: d[0][0]*t.X + d[0][1]*t.Y + d[0][2]*t.Z + d[0][3]*0.,
		Y: d[1][0]*t.X + d[1][1]*t.Y + d[1][2]*t.Z + d[1][3]*0.,
		Z: d[2][0]*t.X + d[2][1]*t.Y + d[2][2]*t.Z + d[2][3]*0.,
	}

}

func (m *Matrix) Transpose() *Matrix {
	//todo also cache?
	//todo or implement as loops?
	d := m.Data
	return &Matrix{
		Data: [4][4]float64{
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
		Data:    *m.Data.inverse(),
		inverse: m,
	}
	return m.inverse
}

func trimAndParseFloat(s string) float64 {
	s = strings.Trim(s, " ")
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	return val
}
