/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	maxLen := int(0) //result
	startIndex := 0
	check := map[rune]int{} //record char lastest index
	for i, c := range s {
		if check[c] != 0 && check[c] >= startIndex { //repeated and in current substring
			startIndex = check[c] //move substring start index
		}
		curLen := i + 1 - startIndex
		if maxLen < curLen {
			maxLen = curLen
		}
		check[c] = i + 1 //mark c location
	}
	return maxLen
}

// @lc code=end

