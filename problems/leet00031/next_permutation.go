package leet00031

import "sort"

func nextPermutation(nums []int) {
	i, j := len(nums)-1, len(nums)-1
	for j > 0 && nums[j-1] >= nums[j] {
		j--
	}
	if j == 0 {
		sort.Ints(nums)
		return
	}
	for i > 0 && nums[j-1] >= nums[i] {
		i--
	}

	nums[j-1], nums[i] = nums[i], nums[j-1]
	l, r := j, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
