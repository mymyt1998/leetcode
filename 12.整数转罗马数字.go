/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */

// @lc code=start
func intToRoman(num int) string {
	charList := map[int]string{0: "", 1: "I", 5: "V", 10: "X", 50: "L", 100: "C", 500: "D", 1000: "M"}
	numList := []int{1000, 500, 100, 50, 10, 5, 1, 0}
	result := ""
	for num > 0 {
		if _, ok := charList[num]; ok {
			result = result + charList[num]
			break
		}
		maxFactor, nearlyFactor := 1, -1
		for i, key := range numList {
			maxFactor = key
			if num >= key {
				nearlyFactor = 0
				break
			}
			j := i + 2 - i%2
			if numList[j]+num >= key {
				nearlyFactor = numList[j]
				break
			}
		}
		result = result + charList[nearlyFactor] + charList[maxFactor]
		num = num - maxFactor + nearlyFactor
	}
	return result
}

//老是在10进制和5进制计算时出错
// @lc code=end

