package leet00017

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	kv := map[byte][]string{
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}
	res := []string{""}

	for i := range digits {
		d := digits[i]
		resLen := len(res)
		for j := 0; j < resLen; j++ {
			for _, c := range kv[d] {
				res = append(res, res[j]+c)
			}
		}
		res = res[resLen:]
	}
	return res
}
