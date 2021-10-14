/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var counter int
	var result, tmp *ListNode
	for l1 != nil || l2 != nil {
		a, b := 0, 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		sum := a + b + counter
		sum, counter = sum%10, sum/10
		//initial head
		if result == nil {
			result = &ListNode{Val: sum}
			tmp = result
		} else {
			tmp.Next = &ListNode{Val: sum}
			tmp = tmp.Next
		}
	}
	if counter > 0 {
		tmp.Next = &ListNode{Val: counter}
	}
	return result
}

// @lc code=end

