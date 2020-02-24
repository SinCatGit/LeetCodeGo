package leet00377

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

type para struct {
	nums   []int
	target int
}

type ans struct {
	one int
}

var qs = []question{
	question{
		para{[]int{1, 2, 3}, 4},
		ans{7},
	},
	question{
		para{[]int{3, 33, 333}, 10000},
		ans{0},
	},
}

func TestCombinationSum3(t *testing.T) {
	ast := assert.New(t)

	for _, q := range qs {
		a, p := q.ans, q.para
		ast.Equal(a.one, combinationSum4(p.nums, p.target))
	}
}

func BenchmarkRemoveDuplicateLetters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			combinationSum4(q.para.nums, q.para.target)
		}
	}
}
