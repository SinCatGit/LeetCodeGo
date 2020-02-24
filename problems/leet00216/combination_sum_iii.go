package leet00216

func combinationSum3(k int, n int) [][]int {

	res := [][]int{}
	sol := []int{}

	backtrack(1, sol, k, n, &res)
	return res
}

func backtrack(start int, sol []int, k int, n int, res *[][]int) {
	if k == 0 && n == 0 {
		cpSol := make([]int, len(sol))
		copy(cpSol, sol)
		*res = append(*res, cpSol)
	}
	if start > 9 || start > n || k <= 0 {
		return
	}
	backtrack(start+1, append(sol, start), k-1, n-start, res)
	backtrack(start+1, sol, k, n, res)
}
