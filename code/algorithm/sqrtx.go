package algorithm

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	left, right := 1, x
	// 中间的位数平方 大于 x，就减一
	for left < right {
		mid := left + (right-left+1)/2
		if mid*mid > x {
			right = mid - 1
		} else {
			left = mid
		}
	}
	return left
}
