package main

import (
	"fmt"
	"math"
)

func ways3(N, M int) int {
	if N < 1 || M < 1 || (N*M)&1 != 0 {
		return 0
	}
	if N == 1 || M == 1 {
		return 1
	}
	max := int(math.Max(float64(N), float64(M)))
	min := int(math.Min(float64(N), float64(M)))
	pre := (1 << min) - 1
	dp := make([][]int, 1<<min)
	for i := range dp {
		dp[i] = make([]int, max+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return process3(pre, 0, max, min, dp)
}

func process3(pre, i, N, M int, dp [][]int) int {
	if dp[pre][i] != -1 {
		return dp[pre][i]
	}
	var ans int
	if i == N {
		if pre == (1<<M)-1 {
			ans = 1
		} else {
			ans = 0
		}
	} else {
		op := (^pre) & ((1 << M) - 1)
		ans = dfs3(op, M-1, i, N, M, dp)
	}
	dp[pre][i] = ans
	return ans
}

func dfs3(op, col, level, N, M int, dp [][]int) int {
	if col == -1 {
		return process3(op, level+1, N, M, dp)
	}
	ans := 0
	ans += dfs3(op, col-1, level, N, M, dp)
	if col > 0 && (op&(3<<(col-1))) == 0 {
		ans += dfs3(op|(3<<(col-1)), col-2, level, N, M, dp)
	}
	return ans
}

func main() {
	N := 10
	M := 10
	result := ways3(N, M)
	fmt.Println(result)
}
