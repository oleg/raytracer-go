package geom

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
