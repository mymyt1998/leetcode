/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	numsLen := len(nums)
	var result [][]int
	numMap := map[int][][]int{}
	countMap := map[int]int{}
	for _, n := range nums {
		countMap[n]++
	}
	aMap := map[int]bool{}
	for i, a := range nums {
		if aMap[a] {
			continue
		}
		aMap[a] = true
		countMap[a]--
		bMap := map[int]bool{}
		for j := i + 1; j < numsLen; j++ {
			b := nums[j]
			if !bMap[b] && !aMap[b] && countMap[b] > 0 {
				numMap[0-a-b] = append(numMap[0-a-b], []int{a, b})
			}
			bMap[b] = true
		}
		if countMap[a] > 0 {
			numMap[0-a-a] = append(numMap[0-a-a], []int{a, a})
		}
		countMap[a]++
	}
	cMap := map[int]bool{}
	for _, c := range nums {
		if cMap[c] {
			continue
		}
		for _, ans := range numMap[c] {
			a, b := ans[0], ans[1]
			countMap[c]--
			countMap[a]--
			countMap[b]--
			if countMap[c] >= 0 && countMap[a] >= 0 && countMap[b] >= 0 {
				if !cMap[a] && !cMap[b] {
					result = append(result, []int{c, a, b})
				}
			}
			countMap[c]++
			countMap[a]++
			countMap[b]++
		}
		cMap[c] = true
	}
	return result
}

//自己想的三重循环->两重循环超时了
//先排序就不需要做这么多检验了
// @lc code=end

