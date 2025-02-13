package maxArea

func maxArea(height []int) (ans int) {
	left, right := 0, len(height)-1
	for left < right {
		area := (right - left) * min(height[left], height[right])
		ans = max(ans, area)
		// 谁短，谁动
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return
}
