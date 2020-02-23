package leet00040

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := [][]int{}
	sol := []int{}
	backtrack(candidates, sol, target, &res)
	return res
}

func backtrack(cand []int, sol []int, target int, res *[][]int) {
	if target == 0 {
		copySol := make([]int, len(sol))
		copy(copySol, sol)
		*res = append(*res, sol)
	}
	if len(cand) == 0 || cand[0] > target {
		return
	}
	// sol = sol[:len(sol):len(sol)]
	backtrack(cand[1:], append(sol, cand[0]), target-cand[0], res)
	i := 1
	for i < len(cand) && cand[i] == cand[i-1] {
		i++
	}

	backtrack(cand[i:], sol, target, res)
}
