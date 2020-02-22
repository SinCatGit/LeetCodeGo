package leet00316

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	in  string
	out string
}{
	{
		"cbacdcbc",
		"acdb",
	},
	{
		"bcabc",
		"abc",
	},
}

func TestRemoveDuplicateLetters(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.out, removeDuplicateLetters(tc.in))
	}
}

func BenchmarkRemoveDuplicateLetters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			removeDuplicateLetters(tc.in)
		}
	}
}
