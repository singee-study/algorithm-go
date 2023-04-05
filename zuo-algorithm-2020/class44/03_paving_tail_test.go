package class44

import (
	"github.com/ImSingee/tt"
	"testing"
)

func TestPavingTail(t *testing.T) {
	testPavingTail := func(t *testing.T, N, M int) {
		t.Helper()

		tt.AssertEqual(t, correctPavingTail(N, M), PavingTail(N, M))
	}

	testPavingTail(t, 2, 2)
	testPavingTail(t, 1, 3)
	testPavingTail(t, 1, 4)
	testPavingTail(t, 1, 4)
	testPavingTail(t, 6, 8)
	testPavingTail(t, 8, 6)
	testPavingTail(t, 10, 10)
}

func correctPavingTail(N, M int) int {
	if N > M {
		N, M = M, N
	}

	// assert N <= M
	if N == 1 && M == 2 {
		return 1
	}
	if N == 1 && M == 3 {
		return 0
	}
	if N == 1 && M == 4 {
		return 1
	}

	if N == 2 && M == 2 {
		return 2
	}

	if N == 6 && M == 8 {
		return 167089
	}
	if N == 10 && M == 10 {
		return 258584046368
	}

	panic("not implemented")
}
