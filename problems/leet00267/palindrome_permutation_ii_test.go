package leet00267

import (
	"sort"
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
	one []string
}

var qs = []question{
	question{
		para{"aabb"},
		ans{[]string{"abba", "baab"}},
	},
	question{
		para{"abc"},
		ans{[]string{}},
	},
	question{
		para{"aabbc"},
		ans{[]string{"abcba", "bacab"}},
	},
}

func TestPermuteUnique(t *testing.T) {
	ast := assert.New(t)

	for _, q := range qs {
		a, p := q.ans, q.para
		r := generatePalindromes(p.s)
		sort.Strings(r)
		ast.Equal(a.one, r)
	}
}

func BenchmarkPermuteUnique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			generatePalindromes(q.para.s)
		}
	}
}
