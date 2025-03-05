package lengthOfLongestSubstring

/*
**
不定长滑动窗口，实际就是双指针移动
*/
func lengthOfLongestSubstring(s string) int {
	ans := 0
	cnt := map[int32]int{}
	left := 0
	for right, v := range s {
		cnt[v]++
		for cnt[v] > 1 {
			cnt[int32(s[left])]--
			left++
		}
		ans = max(ans, right-left+1) // 更新窗口长度最大值
	}
	return ans
}
