package multid

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gray/oned"
	"testing"
)

const delta = 0.000009

func AssertMatrixEqualInDelta(t *testing.T, expected, actual Matrix4) {
	assert.InDeltaMapValues(t, matrixToMap(expected), matrixToMap(actual), delta)
}
func matrixToMap(m Matrix4) map[string]float64 {
	r := map[string]float64{}
	for i, col := range m {
		for j, e := range col {
			k := fmt.Sprintf("%d:%d", i, j)
			r[k] = e
		}
	}
	return r
}

//todo remove duplication
func AssertPointEqualInDelta(t *testing.T, expected, actual oned.Point) {
	assert.InDeltaMapValues(t, pointToMap(expected), pointToMap(actual), delta)
}
func pointToMap(p oned.Point) map[string]float64 {
	return map[string]float64{"X": p.X, "Y": p.Y, "Z": p.Z}
}
