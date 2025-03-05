package twoSum

func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	r := n - 1
	l := 0
	ans := []int{}
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum < target {
			l++
		}
		if sum > target {
			r--
		}
		if sum == target {
			return []int{l + 1, r + 1}
		}
	}
	return ans
}
