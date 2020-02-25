package leet00996

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

type para struct {
	A []int
}

type ans struct {
	one int
}

var qs = []question{
	question{
		para{[]int{1, 8, 17}},
		ans{2},
	},
	question{
		para{[]int{2, 2, 2}},
		ans{1},
	},
}

func TestPermuteUnique(t *testing.T) {
	ast := assert.New(t)

	for _, q := range qs {
		a, p := q.ans, q.para
		ast.Equal(a.one, numSqurefulPerms(p.A))
	}
}

func BenchmarkPermuteUnique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			numSqurefulPerms(q.para.A)
		}
	}
}
