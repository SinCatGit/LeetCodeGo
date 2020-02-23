package leet00039

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
		para{[]int{2, 3, 7}, 7},
		ans{[][]int{{2, 2, 3}, {7}}},
	},
	question{
		para{[]int{2, 3, 5}, 8},
		ans{[][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
	},
}

func TestCombinationSum(t *testing.T) {
	ast := assert.New(t)

	for _, q := range qs {
		a, p := q.ans, q.para
		ast.Equal(a.one, combinationSum(p.candidates, p.target))
	}
}

func BenchmarkRemoveDuplicateLetters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			combinationSum(q.para.candidates, q.para.target)
		}
	}
}
