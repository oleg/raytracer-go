package figure

import (
	"github.com/oleg/raytracer-go/multid"
	"github.com/oleg/raytracer-go/oned"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_inverse_matrix(t *testing.T) {
	m := multid.IdentityMatrix()

	mInverted := m.Inverse()

	assert.Equal(t, multid.IdentityMatrix(), mInverted)
}

func Test_transformation_matrix_for_default_orientation(t *testing.T) {
	from := oned.Point{X: 0, Y: 0, Z: 0}
	to := oned.Point{X: 0, Y: 0, Z: -1}
	up := oned.Vector{X: 0, Y: 1, Z: 0}

	tr := ViewTransform(from, to, up)

	assert.Equal(t, multid.IdentityMatrix().Data, tr.Data)
}

func Test_view_transformation_matrix_looking_in_positive_z_direction(t *testing.T) {
	from := oned.Point{X: 0, Y: 0, Z: 0}
	to := oned.Point{X: 0, Y: 0, Z: 1}
	up := oned.Vector{X: 0, Y: 1, Z: 0}

	tr := ViewTransform(from, to, up)

	assert.Equal(t, multid.Scaling(-1, 1, -1), tr)
}

func Test_view_transformation_moves_the_world(t *testing.T) {
	from := oned.Point{X: 0, Y: 0, Z: 8}
	to := oned.Point{X: 0, Y: 0, Z: 0}
	up := oned.Vector{X: 0, Y: 1, Z: 0}

	tr := ViewTransform(from, to, up)

	assert.Equal(t, multid.Translation(0, 0, -8), tr)
}

func Test_arbitrary_view_transformation(t *testing.T) {
	from := oned.Point{X: 1, Y: 3, Z: 2}
	to := oned.Point{X: 4, Y: -2, Z: 8}
	up := oned.Vector{X: 1, Y: 1, Z: 0}

	tr := ViewTransform(from, to, up)

	expected := multid.NewMatrix4(
		`| -0.50709 | 0.50709 |  0.67612 | -2.36643 |  
		 |  0.76772 | 0.60609 |  0.12122 | -2.82843 |
		 | -0.35857 | 0.59761 | -0.71714 |  0.00000 |
		 |  0.00000 | 0.00000 |  0.00000 |  1.00000 |`)
	multid.AssertMatrixEqualInDelta(t, expected, tr)

}
