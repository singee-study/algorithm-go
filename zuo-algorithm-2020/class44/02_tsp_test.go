package class44

import (
	"github.com/ImSingee/tt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestTSP(t *testing.T) {
	testMatrix := func(t *testing.T, matrix [][]int) {
		t.Helper()

		for i := 0; i < len(matrix); i++ {
			tt.AssertEqual(t, correctTSP(matrix, i), TSP(matrix, i))
		}
	}

	t.Run("matrix2", func(t *testing.T) {
		testMatrix(t, matrix2())
	})

	t.Run("matrix3", func(t *testing.T) {
		testMatrix(t, matrix3())
	})

	t.Run("matrix5", func(t *testing.T) {
		testMatrix(t, matrix5())
	})

	t.Run("random", func(t *testing.T) {
		matrices := make([][][]int, 0, 30)

		for i := 0; i < 30; i++ {
			matrices = append(matrices, generateMatrix(10, 125))
		}

		for _, matrix := range matrices {
			testMatrix(t, matrix)
		}
	})
}

func matrix2() [][]int {
	return [][]int{
		{0, 2},
		{2, 0},
	}
}

func matrix3() [][]int {
	return [][]int{
		{0, 3, 6},
		{3, 0, 9},
		{6, 9, 0},
	}
}

func matrix5() [][]int {
	return [][]int{
		{0, 3, 6, 7, 9},
		{3, 0, 2, 1, 3},
		{6, 2, 0, 4, 6},
		{7, 1, 4, 0, 2},
		{9, 3, 6, 2, 0},
	}
}

func generateMatrix(maxSize int, maxValue int) [][]int {
	rand.Seed(time.Now().UnixNano())
	l := rand.Intn(maxSize) + 1
	matrix := make([][]int, l)
	for i := 0; i < l; i++ {
		matrix[i] = make([]int, l)
		for j := 0; j < l; j++ {
			matrix[i][j] = rand.Intn(maxValue) + 1
		}
	}
	for i := 0; i < l; i++ {
		matrix[i][i] = 0
	}
	return matrix
}

func correctTSP(matrix [][]int, k int) int {
	N := len(matrix)
	set := make([]int, N)
	for i := 0; i < N; i++ {
		set[i] = 1
	}
	return func1(matrix, set, k, k)
}
func func1(matrix [][]int, set []int, k int, start int) int {
	cityNum := 0
	for i := 0; i < len(set); i++ {
		if set[i] != 0 {
			cityNum++
		}
	}

	if cityNum == 1 {
		return matrix[start][k]
	}

	set[start] = 0
	min := math.MaxInt32
	for i := 0; i < len(set); i++ {
		if set[i] != 0 {
			cur := matrix[start][i] + func1(matrix, set, k, i)
			if cur < min {
				min = cur
			}
		}
	}
	set[start] = 1
	return min
}
