package leet00311

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	inA [][]int
	inB [][]int
	out [][]int
}{
	{
		[][]int{{1, 0, 0}, {-1, 0, 3}},
		[][]int{{7, 0, 0}, {0, 0, 0}, {0, 0, 1}},
		[][]int{{7, 0, 0}, {-7, 0, 3}},
	},
}

func TestMultiply(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.out, multiply(tc.inA, tc.inB))
	}
}

func BenchmarkRemoveDuplicateLetters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			multiply(tc.inA, tc.inB)
		}
	}
}
