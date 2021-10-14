/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	result := 0
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	for start, end := 0, len(height)-1; start < end; {
		tmp := (end - start) * min(height[end], height[start])
		if tmp > result {
			result = tmp
		}
		//哪边短换哪边
		if height[end] > height[start] {
			start++
		} else {
			end--
		}
	}
	return result
}

// @lc code=end

