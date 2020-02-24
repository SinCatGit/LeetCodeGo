package leet00047

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

type para struct {
	nums []int
}

type ans struct {
	one [][]int
}

var qs = []question{
	question{
		para{[]int{1, 1, 2}},
		ans{[][]int{
			{1, 1, 2}, {1, 2, 1}, {2, 1, 1},
		}},
	},
}

func TestPermuteUnique(t *testing.T) {
	ast := assert.New(t)

	for _, q := range qs {
		a, p := q.ans, q.para
		result := permuteUnique(p.nums)
		sort.Slice(result, func(i, j int) bool {
			iLen := len(result[i])
			jLen := len(result[j])
			k := 0
			for ; k < iLen && k < jLen; k++ {
				if result[i][k] != result[j][k] {
					return result[i][k] < result[j][k]
				}
			}
			if k == iLen {
				return true
			}
			return false
		})
		ast.Equal(a.one, result)
	}
}

func BenchmarkPermuteUnique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			permuteUnique(q.para.nums)
		}
	}
}
