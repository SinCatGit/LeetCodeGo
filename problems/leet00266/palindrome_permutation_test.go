package leet00266

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

type para struct {
	s string
}

type ans struct {
	one bool
}

var qs = []question{
	question{
		para{"aab"},
		ans{true},
	},
	question{
		para{"code"},
		ans{false},
	},
}

func TestPermuteUnique(t *testing.T) {
	ast := assert.New(t)

	for _, q := range qs {
		a, p := q.ans, q.para
		ast.Equal(a.one, canPermutePalindrome(p.s))
	}
}

func BenchmarkPermuteUnique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			canPermutePalindrome(q.para.s)
		}
	}
}
