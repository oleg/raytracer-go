package geom

import (
	"log"
	"strconv"
	"strings"
)

//todo: decide if I want to return a pointers or structs?
//todo create packages matrix2,matrix3,matrix4?
//todo add iterate function that accept function?
const L4 = 4

//todo implement HasTransformation?
type Matrix struct {
	Data    [L4][L4]float64
	inverse *Matrix
}

var identityMatrixData = [L4][L4]float64{
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
	for i := 0; i < L4; i++ {
		for j := 0; j < L4; j++ {
			for k := 0; k < L4; k++ {
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
	determinant := m.determinant()
	inverse := &Matrix{}
	for i := 0; i < L4; i++ {
		for j := 0; j < L4; j++ {
			inverse.Data[j][i] = m.cofactor(i, j) / determinant
		}
	}
	inverse.inverse = m
	m.inverse = inverse
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

func (m *Matrix) determinant() float64 {
	r := 0.
	for i, v := range m.Data[0] {
		r += v * m.cofactor(0, i)
	}
	return r
}

func (m *Matrix) cofactor(row, column int) float64 {
	return m.minor(row, column) * sign(row, column)
}

func (m *Matrix) minor(row, column int) float64 {
	return m.submatrix(row, column).determinant()
}

func (m *Matrix) submatrix(row, column int) *matrix3x3 {
	r := &matrix3x3{}
	for ri, mi := 0, 0; mi < L4; mi++ {
		if mi == row {
			continue
		}
		for rj, mj := 0, 0; mj < L4; mj++ {
			if mj == column {
				continue
			}
			r[ri][rj] = m.Data[mi][mj]
			rj++
		}
		ri++
	}
	return r
}

type matrix3x3 [3][3]float64

func (m *matrix3x3) determinant() float64 {
	r := 0.
	for i, v := range m[0] {
		r += v * m.cofactor(0, i)
	}
	return r
}

func (m *matrix3x3) cofactor(row, column int) float64 {
	return m.minor(row, column) * sign(row, column)
}

func (m *matrix3x3) minor(row, column int) float64 {
	return m.submatrix(row, column).determinant()
}

func (m *matrix3x3) submatrix(row, column int) *matrix2x2 {
	//todo: how to reuse submatrix code?
	r := &matrix2x2{}
	for ri, mi := 0, 0; mi < 3; mi++ {
		if mi == row {
			continue
		}
		for rj, mj := 0, 0; mj < 3; mj++ {
			if mj == column {
				continue
			}
			r[ri][rj] = m[mi][mj]
			rj++
		}
		ri++
	}
	return r
}

type matrix2x2 [2][2]float64

func (m *matrix2x2) determinant() float64 {
	return m[0][0]*m[1][1] - m[0][1]*m[1][0]
}

func sign(row int, column int) float64 {
	if (row+column)%2 == 0 {
		return 1
	} else {
		return -1
	}
}
