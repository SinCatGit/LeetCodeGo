package leet00216

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

type para struct {
	k int
	n int
}

type ans struct {
	one [][]int
}

var qs = []question{
	question{
		para{2, 18},
		ans{[][]int{}},
	},
	question{
		para{3, 7},
		ans{[][]int{{1, 2, 4}}},
	},
	question{
		para{3, 9},
		ans{[][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}},
	},
}

func TestCombinationSum3(t *testing.T) {
	ast := assert.New(t)

	for _, q := range qs {
		a, p := q.ans, q.para
		ast.Equal(a.one, combinationSum3(p.k, p.n))
	}
}

func BenchmarkRemoveDuplicateLetters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			combinationSum3(q.para.k, q.para.n)
		}
	}
}
