package leet00311

func multiply(A [][]int, B [][]int) [][]int {
	m := len(A)
	n := len(B[0])
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	noneZeroB := make([][]int, n)

	for i := 0; i < n; i++ {
		col := make([]int, 0)
		for j := range B {
			if B[j][i] != 0 {
				col = append(col, j, B[j][i])
			}
		}
		noneZeroB[i] = col
	}

	for i := range A {
		for k := range noneZeroB {
			t := 0
			for j := 0; j < len(noneZeroB[k]); j += 2 {
				t += A[i][noneZeroB[k][j]] * noneZeroB[k][j+1]
			}
			res[i][k] = t
		}
	}
	return res
}
