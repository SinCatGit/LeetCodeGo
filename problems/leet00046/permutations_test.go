package leet00046

import (
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
		para{[]int{1, 2, 3}},
		ans{[][]int{
			{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2},
		}},
	},
}

func TestCombinationSum3(t *testing.T) {
	ast := assert.New(t)

	for _, q := range qs {
		a, p := q.ans, q.para
		ast.Equal(a.one, permute(p.nums))
	}
}

func BenchmarkRemoveDuplicateLetters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			permute(q.para.nums)
		}
	}
}
