package leet00377

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 0; i < target+1; i++ {
		for _, n := range nums {
			if n > i {
				break
			}
			dp[i] += dp[i-n]
		}
	}
	return dp[target]
}
