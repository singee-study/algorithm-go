package class44

import (
	"algorithm/utils"
	"math"
)

/*
有N个城市，任何两个城市之间的都有距离，任何一座城市到自己的距离都为0。所有点到点的距
离都存在一个N*N的二维数组matrix里，也就是整张图由邻接矩阵表示。现要求一旅行商从k城市
出发必须经过每一个城市且只在一个城市逗留一次，最后回到出发的k城，返回总距离最短的路的 距离。参数给定一个matrix，给定k。
*/

func TSP(matrix [][]int, k int) int {
	n := len(matrix)

	state := (1 << n) - 1
	dp := utils.New2dSliceV(1<<n, len(matrix), -1)

	return tsp(matrix, k, dp, state, k)
}

func tsp(matrix [][]int, k int, dp [][]int, state int, start int) int {
	if dp[state][start] != -1 {
		return dp[state][start]
	}

	xs := state - (1 << start)

	if xs == 0 { // 只剩下一个城市（这个城市必为 start）
		return matrix[start][k]
	}

	dp[state][start] = math.MaxInt32

	for i := 0; i < len(matrix); i++ {
		if xs&(1<<i) == 0 { // 这个城市已经被访问过了
			continue
		}

		dp[state][start] = utils.Min(
			dp[state][start],
			tsp(matrix, k, dp, xs, i)+matrix[start][i],
		)
	}

	return dp[state][start]
}
