package multid

//todo: where to put this methods?
func Translation(x, y, z float64) Matrix4 {
	return Matrix4{
		{1, 0, 0, x},
		{0, 1, 0, y},
		{0, 0, 1, z},
		{0, 0, 0, 1},
	}
}