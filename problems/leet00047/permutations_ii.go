package leet00047

func permuteUnique(nums []int) [][]int {
	answer := [][]int{}
	counter := map[int]int{}
	numsLen := len(nums)

	for i := 0; i < numsLen; i++ {
		counter[nums[i]]++
	}

	dfs(&answer, 0, &nums, counter)
	return answer
}

func dfs(answer *[][]int, idx int, path *[]int, counter map[int]int) {
	if len(*path) == idx {
		sol := make([]int, len(*path))
		copy(sol, *path)
		*answer = append(*answer, sol)
	}
	for k, v := range counter {
		if v > 0 {
			counter[k]--
			(*path)[idx] = k
			dfs(answer, idx+1, path, counter)
			counter[k]++
		}
	}
}
