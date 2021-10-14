/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	result := int(0)
	Positive := true
	leading := true
	for _, c := range s {
		if leading {
			if c == ' ' {
				continue
			}
			leading = false
			if c == '-' {
				Positive = false
				continue
			}
			if c == '+' {
				continue
			}
		}
		if c > '9' || c < '0' {
			break
		} else {
			k := int(c - '0')
			if Positive {
				if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && k >= 7) {
					result = math.MaxInt32
					break
				}
				result = result*10 + k
			} else {
				if result < math.MinInt32/10 || (result == math.MinInt32/10 && k > 7) {
					result = math.MinInt32
					break
				}
				result = result*10 - k
			}
		}
	}
	return result
}

// @lc code=end

