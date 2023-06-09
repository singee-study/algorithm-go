package class44

import (
	"algorithm/utils"
)

/*
你有无限的1*2的砖块，要铺满M*N的区域，
不同的铺法有多少种?

*/

func PavingTail(N, M int) int {
	// N 行 M 列，保证 M 较小且小于 31
	if N < M {
		M, N = N, M
	}

	state := (1 << M) - 1 // 默认 (-1 行) 全铺满
	dp := utils.New2dSlice(1<<M, N)

	return pavingTail(dp, state, 0, N, M)
}

func pavingTail(dp [][]int, pre int, cur int, N, M int) (v int) {
	if cur == N { // 末行+1，要求上一行必须全铺满
		if pre == (1<<M)-1 {
			return 1
		} else {
			return 0
		}
	}

	if dp[pre][cur] != 0 {
		return dp[pre][cur]
	}
	defer func() {
		dp[pre][cur] = v
	}()

	cl := ^pre & (1<<M - 1) // 本行情况
	// 上一行为 0 的本行必须向上铺，为 0 的可选

	return pavingTailDfs(dp, cl, cur, 0, N, M)
}

func pavingTailDfs(dp [][]int, cl int, cur int, col int, N, M int) (v int) {
	if col == M { // 本行结束，下一行
		return pavingTail(dp, cl, cur+1, N, M)
	}
	//if (cl & (1 << col)) != 0 { // 本行已铺
	//	return pavingTailDfs(dp, cl, cur, col+1, N, M)
	//}

	count := pavingTailDfs(dp, cl, cur, col+1, N, M) // 不铺

	if col+1 < M && (cl&(1<<col)) == 0 && (cl&(1<<(col+1))) == 0 { // 可以铺
		count += pavingTailDfs(dp, cl|(1<<col)|(1<<(col+1)), cur, col+2, N, M)
	}

	return count
}
