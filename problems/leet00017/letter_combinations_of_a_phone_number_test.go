package leet00017

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	in  string
	out []string
}{
	{
		"23",
		[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
	},
}

func TestMultiply(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		res := letterCombinations(tc.in)
		sort.Strings(res)
		ast.Equal(tc.out, res)
	}
}

func BenchmarkRemoveDuplicateLetters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			letterCombinations(tc.in)
		}
	}
}
