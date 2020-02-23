package leet00040

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

type para struct {
	candidates []int
	target     int
}

type ans struct {
	one [][]int
}

var qs = []question{
	question{
		para{[]int{10, 1, 2, 7, 6, 1, 5}, 8},
		ans{[][]int{
			{1, 1, 6},
			{1, 2, 5},
			{1, 7},
			{2, 6},
		}},
	},
	question{
		para{[]int{2, 5, 2, 1, 2}, 5},
		ans{[][]int{{1, 2, 2}, {5}}},
	},
}

func TestCombinationSum(t *testing.T) {
	ast := assert.New(t)

	for _, q := range qs {
		a, p := q.ans, q.para
		ast.Equal(a.one, combinationSum2(p.candidates, p.target))
	}
}

func BenchmarkRemoveDuplicateLetters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			combinationSum2(q.para.candidates, q.para.target)
		}
	}
}
