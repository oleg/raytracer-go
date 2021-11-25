package geom

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const Delta = 0.000009

func AssertVectorEqualInDelta(t *testing.T, expected, actual Vector) {
	assert.InDeltaMapValues(t, vectorToMap(expected), vectorToMap(actual), Delta)
}

func vectorToMap(v Vector) map[string]float64 {
	return map[string]float64{"X": v.X, "Y": v.Y, "Z": v.Z}
}

func AssertColorEqualInDelta(t *testing.T, expected, actual Color) {
	assert.InDeltaMapValues(t, colorToMap(expected), colorToMap(actual), Delta)
}
func colorToMap(v Color) map[string]float64 {
	return map[string]float64{"R": v.R, "G": v.G, "B": v.B}
}

func AssertMatrixEqualInDelta(t *testing.T, expected, actual *Matrix) {
	assert.InDeltaMapValues(t, matrixToMap(expected), matrixToMap(actual), Delta)
}

func matrixToMap(m *Matrix) map[string]float64 {
	r := map[string]float64{}
	for i, col := range m.Data {
		for j, e := range col {
			k := fmt.Sprintf("%d:%d", i, j)
			r[k] = e
		}
	}
	return r
}

func AssertPointEqualInDelta(t *testing.T, expected, actual Point) {
	assert.InDeltaMapValues(t, pointToMap(expected), pointToMap(actual), Delta)
}

func pointToMap(p Point) map[string]float64 {
	return map[string]float64{"X": p.X, "Y": p.Y, "Z": p.Z}
}
