package leet00060

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
	one string
}

var qs = []question{
	question{
		para{3, 3},
		ans{"213"},
	},
	// question{
	// 	para{4, 9},
	// 	ans{"2314"},
	// },
}

func TestPermuteUnique(t *testing.T) {
	ast := assert.New(t)

	for _, q := range qs {
		a, p := q.ans, q.para
		ast.Equal(a.one, getPermutation(p.n, p.k))
	}
}

func BenchmarkPermuteUnique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			getPermutation(q.para.n, q.para.k)
		}
	}
}
