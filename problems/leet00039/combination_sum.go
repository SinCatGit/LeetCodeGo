package leet00039

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := [][]int{}
	sol := []int{}
	backtrack(candidates, sol, target, &res)
	return res
}

func backtrack(cand []int, sol []int, target int, res *[][]int) {
	if target == 0 {
		*res = append(*res, sol)
	}
	if len(cand) == 0 || cand[0] > target {
		return
	}
	sol = sol[:len(sol):len(sol)]
	backtrack(cand, append(sol, cand[0]), target-cand[0], res)
	backtrack(cand[1:], sol, target, res)
}
