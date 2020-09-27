package multid

//is it copying?
func determinant4x4(m Matrix) float64 {
	r := 0.
	for i, v := range m[0] {
		r += v * cofactor4x4(m, 0, i)
	}
	return r
}

//todo:test
func cofactor4x4(m Matrix, row, column int) float64 {
	return minor4x4(m, row, column) * sign(row, column)
}

func minor4x4(m Matrix, row, column int) float64 {
	sm := submatrix4x4(m, row, column)
	return determinant3x3(&sm)
}

func submatrix4x4(m Matrix, row, column int) [3][3]float64 {
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
