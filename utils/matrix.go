package utils

func New2dSlice(m, n int) [][]int {
	return new2dSliceV(m, n, 0)
}

func New2dSliceV(m, n int, v int) [][]int {
	return new2dSliceV(m, n, v)
}

// generate a m*n grid slice
func new2dSliceV(m, n, v int) [][]int {
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
		for j := range matrix[i] {
			matrix[i][j] = v
		}
	}
	return matrix
}
