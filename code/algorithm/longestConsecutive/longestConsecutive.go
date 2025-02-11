package longestConsecutive

func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}
	res := 0
	for v := range m {
		if m[v-1] {
			continue
		}
		// 不存在则开始计算爹值
		conut := 0
		w := v
		for m[w] {
			w++
			conut++
		}
		res = max(res, conut)
	}
	return res
}
