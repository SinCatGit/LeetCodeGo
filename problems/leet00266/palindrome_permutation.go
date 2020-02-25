package leet00266

func canPermutePalindrome(s string) bool {
	cnt := 0
	counter := make(map[rune]int)
	for _, b := range s {
		counter[b]++
	}
	for _, v := range counter {
		cnt += v % 2
	}
	return cnt < 2
}
