package leet00316

import "bytes"

func removeDuplicateLetters(s string) string {
	b := []byte(s)
	var r []byte
	for i, c := range b {
		if len(r) == 0 {
			r = append(r, c)
			continue
		}
		if bytes.IndexByte(r, c) != -1 {
			continue
		}

		for len(r) > 0 && c < r[len(r)-1] && bytes.IndexByte(b[i+1:], r[len(r)-1]) != -1 {
			r = r[:len(r)-1]
		}
		r = append(r, c)
	}
	return string(r)
}
