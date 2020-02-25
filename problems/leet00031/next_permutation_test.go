package leet00031

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
	one []int
}

var qs = []question{
	question{
		para{[]int{5, 4, 3, 6, 4, 3, 2}},
		ans{[]int{5, 4, 4, 2, 3, 3, 6}},
	},
	question{
		para{[]int{3, 2, 1}},
		ans{[]int{1, 2, 3}},
	},
	question{
		para{[]int{1, 2, 3}},
		ans{[]int{1, 3, 2}},
	},
	question{
		para{[]int{1, 1, 5}},
		ans{[]int{1, 5, 1}},
	},
	question{
		para{[]int{1, 1}},
		ans{[]int{1, 1}},
	},
	question{
		para{[]int{5, 1, 1}},
		ans{[]int{1, 1, 5}},
	},
}

func TestPermuteUnique(t *testing.T) {
	ast := assert.New(t)

	for _, q := range qs {
		a, p := q.ans, q.para
		nextPermutation(p.nums)
		ast.Equal(a.one, p.nums)
	}
}

func BenchmarkPermuteUnique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			nextPermutation(q.para.nums)
		}
	}
}
