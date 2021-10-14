/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	sLen := int(len(s))
	result := ""
	maxLen := 1
	for i, _ := range s {
		start := i - 1
		end := i + 1
		for start >= 0 && s[start] == s[i] {
			start--
		}
		for end < sLen && s[end] == s[i] {
			end++
		}
		for start >= 0 && end < sLen {
			if s[start] != s[end] {
				break
			}
			start--
			end++
		}
		curLen := end - start + 1
		if maxLen < curLen {
			maxLen = end - start + 1
			result = s[start+1 : end]
		}
	}
	return result
}

// @lc code=end

