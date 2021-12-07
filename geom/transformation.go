package geom

import "math"

func Translation(x, y, z float64) *Matrix {
	return &Matrix{
		Data: [4][4]float64{
			{1, 0, 0, x},
			{0, 1, 0, y},
			{0, 0, 1, z},
			{0, 0, 0, 1},
		},
	}
}

func Scaling(x, y, z float64) *Matrix {
	return &Matrix{
		Data: [4][4]float64{
			{x, 0, 0, 0},
			{0, y, 0, 0},
			{0, 0, z, 0},
			{0, 0, 0, 1},
		},
	}
}

func RotationX(r float64) *Matrix {
	return &Matrix{
		Data: [4][4]float64{
			{1, 0, 0, 0},
			{0, math.Cos(r), -math.Sin(r), 0},
			{0, math.Sin(r), math.Cos(r), 0},
			{0, 0, 0, 1},
		},
	}
}

func RotationY(r float64) *Matrix {
	return &Matrix{
		Data: [4][4]float64{
			{math.Cos(r), 0, math.Sin(r), 0},
			{0, 1, 0, 0},
			{-math.Sin(r), 0, math.Cos(r), 0},
			{0, 0, 0, 1},
		},
	}
}

func RotationZ(r float64) *Matrix {
	return &Matrix{
		Data: [4][4]float64{
			{math.Cos(r), -math.Sin(r), 0, 0},
			{math.Sin(r), math.Cos(r), 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 1},
		},
	}
}

func Shearing(xy, xz, yx, yz, zx, zy float64) *Matrix {
	return &Matrix{
		Data: [4][4]float64{
			{1, xy, xz, 0},
			{yx, 1, yz, 0},
			{zx, zy, 1, 0},
			{0, 0, 0, 1},
		},
	}
}
