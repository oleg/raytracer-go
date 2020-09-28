package multid

import (
	"github.com/oleg/graytracer/oned"
	"log"
	"strconv"
	"strings"
)

//todo: !!! decide if I want to return a pointers or structs !!!
//todo create packages matrix2,matrix3,matrix4
//todo add iterate function that accept function
const L4 = 4

type Matrix [L4][L4]float64

var IdentityMatrix = Matrix{
	{1, 0, 0, 0},
	{0, 1, 0, 0},
	{0, 0, 1, 0},
	{0, 0, 0, 1},
}

func IdentityMatrixF() Matrix {
	return Matrix{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

//todo must?
func NewMatrix(str string) Matrix {
	m := Matrix{}
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

func (m Matrix) Multiply(o Matrix) Matrix {
	r := Matrix{}
	for i := 0; i < L4; i++ {
		for j := 0; j < L4; j++ {
			for k := 0; k < L4; k++ {
				r[i][j] += m[i][k] * o[k][j]
			}
		}
	}
	return r
}

func (m Matrix) MultiplyPoint(o oned.Point) oned.Point {
	return oned.Point(m.multiplyTuple(oned.Tuple(o), 1.))
}

//todo: remove duplication
func (m Matrix) MultiplyVector(o oned.Vector) oned.Vector {
	return oned.Vector(m.multiplyTuple(oned.Tuple(o), 0.))
}

func (m Matrix) multiplyTuple(t oned.Tuple, w float64) oned.Tuple {
	return oned.Tuple{
		m[0][0]*t.X + m[0][1]*t.Y + m[0][2]*t.Z + m[0][3]*w,
		m[1][0]*t.X + m[1][1]*t.Y + m[1][2]*t.Z + m[1][3]*w,
		m[2][0]*t.X + m[2][1]*t.Y + m[2][2]*t.Z + m[2][3]*w,
	}
}

func (m Matrix) Transpose() Matrix {
	//todo or implement as loops?
	return Matrix{
		{m[0][0], m[1][0], m[2][0], m[3][0]},
		{m[0][1], m[1][1], m[2][1], m[3][1]},
		{m[0][2], m[1][2], m[2][2], m[3][2]},
		{m[0][3], m[1][3], m[2][3], m[3][3]},
	}
}

//todo: quick fix gives 10x improvements
var cache = make(map[Matrix]Matrix)

func (m Matrix) Inverse() Matrix {
	if cached, ok := cache[m]; ok {
		return cached
	}
	determinant := determinant4x4(m)
	inverse := Matrix{}
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
