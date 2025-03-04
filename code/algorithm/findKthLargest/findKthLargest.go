package findKthLargest

import "slices"

func findKthLargest(nums []int, k int) int {
	slices.Sort(nums)
	return nums[k-1]
}
