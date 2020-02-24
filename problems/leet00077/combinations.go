package leet00077

func combine(n int, k int) [][]int {
	if k == 1 {
		res := make([][]int, n)
		for i := 0; i < n; i++ {
			res[i] = []int{i + 1}
		}
		return res
	}
	if k == n {
		ele := make([]int, n)
		for i := 0; i < n; i++ {
			ele[i] = i + 1
		}
		return [][]int{ele}
	}

	if k < 1 || k > n {
		return [][]int{}
	}

	unused := combine(n-1, k)
	part := combine(n-1, k-1)
	for i := 0; i < len(part); i++ {
		part[i] = append(part[i], n)
	}
	return append(unused, part...)
}
