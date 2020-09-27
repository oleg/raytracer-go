package figure

import (
	"github.com/stretchr/testify/assert"
	"gray/multid"
	"gray/oned"
	"testing"
)

func Test_transformation_matrix_for_default_orientation(t *testing.T) {
	from := oned.Point{0, 0, 0}
	to := oned.Point{0, 0, -1}
	up := oned.Vector{0, 1, 0}

	tr := ViewTransform(from, to, up)

	assert.Equal(t, multid.IdentityMatrix, tr)
}

func Test_view_transformation_matrix_looking_in_positive_z_direction(t *testing.T) {
	from := oned.Point{0, 0, 0}
	to := oned.Point{0, 0, 1}
	up := oned.Vector{0, 1, 0}

	tr := ViewTransform(from, to, up)

	assert.Equal(t, multid.Scaling(-1, 1, -1), tr)
}

func Test_view_transformation_moves_the_world(t *testing.T) {
	from := oned.Point{0, 0, 8}
	to := oned.Point{0, 0, 0}
	up := oned.Vector{0, 1, 0}

	tr := ViewTransform(from, to, up)

	assert.Equal(t, multid.Translation(0, 0, -8), tr)
}

func Test_arbitrary_view_transformation(t *testing.T) {
	from := oned.Point{1, 3, 2}
	to := oned.Point{4, -2, 8}
	up := oned.Vector{1, 1, 0}

	tr := ViewTransform(from, to, up)

	expected := multid.NewMatrix4(
		`| -0.50709 | 0.50709 |  0.67612 | -2.36643 |  
		 |  0.76772 | 0.60609 |  0.12122 | -2.82843 |
		 | -0.35857 | 0.59761 | -0.71714 |  0.00000 |
		 |  0.00000 | 0.00000 |  0.00000 |  1.00000 |`)
	multid.AssertMatrixEqualInDelta(t, expected, tr)

}
