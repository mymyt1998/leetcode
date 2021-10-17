/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums) //用排序来保证不重复计算
	var result [][]int
	for i := 0; i < n; i++ {
		//不用重复的a
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		a := 0 - nums[i]
		k := n - 1
		for j := i + 1; j < n; j++ {
			//不用重复的b
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			//找到等于a-b的c
			for k > j && nums[k]+nums[j] > a {
				k--
			}
			//没找到合适的c
			if j == k {
				break
			}
			if nums[k]+nums[j] == a {
				result = append(result, []int{nums[i], nums[j], nums[k]})
			}
		}
	}

	return result
}

//自己想的三重循环->两重循环超时了
//先排序就不需要做这么多检验了
// @lc code=end

