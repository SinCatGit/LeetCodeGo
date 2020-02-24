package leet00046

func permute(nums []int) [][]int {
	answer := [][]int{}
	dfs(&answer, 0, nums)
	return answer
}

func dfs(answer *[][]int, idx int, nums []int) {
	if len(nums) == idx {
		sol := make([]int, len(nums))
		copy(sol, nums)
		*answer = append(*answer, sol)
	}

	for i := idx; i < len(nums); i++ {
		nums[i], nums[idx] = nums[idx], nums[i]
		dfs(answer, idx+1, nums)
		nums[i], nums[idx] = nums[idx], nums[i]
	}
}
