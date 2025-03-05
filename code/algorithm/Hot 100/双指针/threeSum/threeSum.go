package threeSum

import (
	"slices"
)

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	n := len(nums)
	ans := make([][]int, 0)
	for i, v := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 后续就是两数之和等于-v了
		// 尾指针
		k := n - 1
		j := i + 1
		for j < k {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			sum := nums[j] + nums[k] + v
			if sum < 0 {
				j++
			}
			if sum > 0 {
				k--
			}
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				for j++; j < k && nums[j] == nums[j-1]; j++ {
				} // 跳过重复数字
				for k--; k > j && nums[k] == nums[k+1]; k-- {
				} // 跳过重复数字
			}

		}
	}
	return ans
}
