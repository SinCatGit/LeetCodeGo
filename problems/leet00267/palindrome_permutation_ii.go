package leet00267

func generatePalindromes(s string) []string {
	// write your code here
	res := make([]string, 0)
	counter := make(map[rune]int)
	mid := ""
	for _, c := range s {
		counter[c]++
	}
	for c, v := range counter {
		if v%2 == 1 && mid != "" {
			return []string{}
		}

		if v%2 == 1 {
			mid = string(c)
		}

		counter[c] = int(v / 2)
	}

	var dfs func(string, int)
	dfs = func(path string, left int) {
		if left == 0 {
			res = append(res, path+mid+reverse(path))
		}
		for c, v := range counter {
			if v > 0 {
				counter[c]--
				dfs(path+string(c), left-1)
				counter[c]++
			}
		}
	}
	dfs("", int(len(s)/2))

	return res

}

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}
