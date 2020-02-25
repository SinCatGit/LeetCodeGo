package leet00996

import "math"

func isSquare(x int) bool {
	r := int(math.Sqrt(float64(x)))
	return r*r == x
}

func numSqurefulPerms(A []int) int {
	size := len(A)
	count := make(map[int]int, size)
	for _, a := range A {
		count[a]++
	}
	conds := make(map[int][]int, size)
	for x := range count {
		for y := range count {
			if isSquare(x + y) {
				conds[x] = append(conds[x], y)
			}
		}
	}

	cnt := 0
	var dfs func(int, int)

	dfs = func(x, left int) {
		if left == 0 {
			cnt++
			return
		}
		count[x]--
		for _, y := range conds[x] {
			if count[y] > 0 {
				dfs(y, left-1)
			}
		}
		count[x]++
	}

	for x := range count {
		dfs(x, size-1)
	}
	return cnt

}
