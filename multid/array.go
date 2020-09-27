package multid

//how to reuse this?
func multiply(a, b [][]float64) [][]float64 {
	if len(a) == 0 && len(b) == 0 {
		return nil
	}

	aw := len(a)
	ah := len(a[0])

	bw := len(b)
	bh := len(b[0])

	if ah != bw {
		panic("not possible 2") //todo test
	}

	c := make([][]float64, aw)
	for i := 0; i < aw; i++ {
		c[i] = make([]float64, bh)
	}

	for i := 0; i < aw; i++ {
		for j := 0; j < bh; j++ {
			for k := 0; k < ah; /*&& k < bw*/ k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return c
}
