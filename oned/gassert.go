package oned

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//todo move to test file?
func AssertVectorEqualInDelta(t *testing.T, expected, actual Vector) {
	assert.InDeltaMapValues(t, vectorToMap(expected), vectorToMap(actual), Delta)
}
func vectorToMap(v Vector) map[string]float64 {
	return map[string]float64{"X": v.X, "Y": v.Y, "Z": v.Z}
}

//todo move to test file?
func AssertColorEqualInDelta(t *testing.T, expected, actual Color) {
	assert.InDeltaMapValues(t, colorToMap(expected), colorToMap(actual), Delta)
}
func colorToMap(v Color) map[string]float64 {
	return map[string]float64{"R": v.R(), "G": v.G(), "B": v.B()}
}
