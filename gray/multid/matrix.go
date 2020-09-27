package multid

import (
	"gray/oned"
	"log"
	"strconv"
	"strings"
)

//todo: !!! decide if I want to return a pointers or structs !!!
//todo create packages matrix2,matrix3,matrix4
//todo add iterate function that accept function
const L4 = 4

type Matrix4 [L4][L4]float64

var IdentityMatrix = Matrix4{
	{1, 0, 0, 0},
	{0, 1, 0, 0},
	{0, 0, 1, 0},
	{0, 0, 0, 1},
}

func IdentityMatrixF() Matrix4 {
	return Matrix4{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

//todo must?
func NewMatrix4(str string) Matrix4 {
	m := Matrix4{}
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
			m[i][j] = trimAndParseFloat(s)
		}
	}
	return m
}

func (m Matrix4) Multiply(o Matrix4) Matrix4 {
	r := Matrix4{}
	for i := 0; i < L4; i++ {
		for j := 0; j < L4; j++ {
			for k := 0; k < L4; k++ {
				r[i][j] += m[i][k] * o[k][j]
			}
		}
	}
	return r
}

func (m Matrix4) MultiplyPoint(o oned.Point) oned.Point {
	return oned.Point(m.multiplyTuple(oned.Tuple(o), 1.))
}

//todo: remove duplication
func (m Matrix4) MultiplyVector(o oned.Vector) oned.Vector {
	return oned.Vector(m.multiplyTuple(oned.Tuple(o), 0.))
}

func (m Matrix4) multiplyTuple(t oned.Tuple, x float64) oned.Tuple {
	//todo: refactor
	r := [4]float64{}
	o := [4]float64{t.X, t.Y, t.Z, x}
	for i := 0; i < L4; i++ {
		for k := 0; k < L4; k++ {
			r[i] += m[i][k] * o[k]
		}
	}
	return oned.Tuple{r[0], r[1], r[2]}
}

func (m Matrix4) Transpose() Matrix4 {
	//todo or implement as loops?
	return Matrix4{
		{m[0][0], m[1][0], m[2][0], m[3][0]},
		{m[0][1], m[1][1], m[2][1], m[3][1]},
		{m[0][2], m[1][2], m[2][2], m[3][2]},
		{m[0][3], m[1][3], m[2][3], m[3][3]},
	}
}

func (m Matrix4) determinant() float64 {
	return determinant4x4(m)
}

//todo: quick fix gives 10x improvements
var cache = make(map[Matrix4]Matrix4)

func (m Matrix4) Inverse() Matrix4 {
	if cached, ok := cache[m]; ok {
		return cached
	}
	determinant := m.determinant()
	inverse := Matrix4{}
	for i := 0; i < L4; i++ {
		for j := 0; j < L4; j++ {
			inverse[j][i] = cofactor4x4(m, i, j) / determinant
		}
	}
	cache[m] = inverse
	return inverse
}

func trimAndParseFloat(s string) float64 {
	s = strings.Trim(s, " ")
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	return val
}

//is it copying?
func determinant4x4(m Matrix4) float64 {
	r := 0.
	for i, v := range m[0] {
		r += v * cofactor4x4(m, 0, i)
	}
	return r
}

//todo:test
func cofactor4x4(m Matrix4, row, column int) float64 {
	return minor4x4(m, row, column) * sign(row, column)
}

func minor4x4(m Matrix4, row, column int) float64 {
	sm := submatrix4x4(m, row, column)
	return determinant3x3(&sm)
}

func submatrix4x4(m Matrix4, row, column int) [3][3]float64 {
	r := [3][3]float64{}
	for ri, mi := 0, 0; mi < L4; mi++ {
		if mi == row {
			continue
		}
		for rj, mj := 0, 0; mj < L4; mj++ {
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

func determinant3x3(m *[3][3]float64) float64 {
	r := 0.
	for i, v := range m[0] {
		r += v * cofactor3x3(m, 0, i)
	}
	return r
}

func cofactor3x3(m *[3][3]float64, row, column int) float64 {
	return minor3x3(m, row, column) * sign(row, column)
}

func minor3x3(m *[3][3]float64, row, column int) float64 {
	sm := submatrix3x3(m, row, column)
	return determinant2x2(&sm)
}

//todo: how to reuse submatrix code?
func submatrix3x3(m *[3][3]float64, row, column int) [2][2]float64 {
	r := [2][2]float64{}
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

func determinant2x2(m *[2][2]float64) float64 {
	return m[0][0]*m[1][1] - m[0][1]*m[1][0]
}

func sign(row int, column int) float64 {
	if (row+column)%2 == 0 {
		return 1
	} else {
		return -1
	}
}
