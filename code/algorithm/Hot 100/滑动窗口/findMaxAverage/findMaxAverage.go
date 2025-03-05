package findMaxAverage

import "math"

func findMaxAverage(nums []int, k int) float64 {
	res := math.MinInt
	sum := 0
	for i, in := range nums {
		sum += in
		if i < k-1 {
			continue
		}
		res = max(res, sum)
		// 离开窗口
		sum -= nums[i-k+1]
	}
	return float64(res) / float64(k)
}
