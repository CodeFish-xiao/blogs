package maxVowels

// 长度恰好为 k 的子串中，最多可以包含多少个元音字母
/***
入：下标为 i 的元素进入窗口，更新相关统计量。如果 i<k−1 则重复第一步。
更新：更新答案。一般是更新最大值/最小值。
出：下标为 i−k+1 的元素离开窗口，更新相关统计量。
*/
func maxVowels(s string, k int) int {
	res := 0
	vowel := 0
	for i, in := range s {
		if in == 'a' || in == 'e' || in == 'i' || in == 'o' || in == 'u' {
			vowel++
		}
		// 窗口大小不足 k
		if i < k-1 {
			continue
		}
		res = max(res, vowel)
		// 离开窗口
		out := s[i-k+1]
		if out == 'a' || out == 'e' || out == 'i' || out == 'o' || out == 'u' {
			vowel--
		}
	}
	return res
}
