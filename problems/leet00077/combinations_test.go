package leet00077

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

type para struct {
	n int
	k int
}

type ans struct {
	one [][]int
}

var qs = []question{
	question{
		para{4, 2},
		ans{[][]int{
			{1, 2}, {1, 3}, {2, 3}, {1, 4}, {2, 4}, {3, 4},
		}},
	},
	question{
		para{4, 1},
		ans{[][]int{
			{1}, {2}, {3}, {4},
		}},
	},
	question{
		para{4, 4},
		ans{[][]int{
			{1, 2, 3, 4},
		}},
	},
}

func TestCombinationSum3(t *testing.T) {
	ast := assert.New(t)

	for _, q := range qs {
		a, p := q.ans, q.para
		ast.Equal(a.one, combine(p.n, p.k))
	}
}

func BenchmarkRemoveDuplicateLetters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			combine(q.para.n, q.para.k)
		}
	}
}
