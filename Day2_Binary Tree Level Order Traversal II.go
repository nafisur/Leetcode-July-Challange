package leetcode

/*
Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its bottom-up level order traversal as:
[
  [15,7],
  [9,20],
  [3]
]
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	result := [][]int{}
	helper([]*TreeNode{root}, &result)
	return result[1:]
}

func helper(level []*TreeNode, result *[][]int) {
	if len(level) == 0 {
		return
	}
	nextLevel := []*TreeNode{}
	for _, v := range level {
		if v != nil {
			nextLevel = append(nextLevel, v.Left)
			nextLevel = append(nextLevel, v.Right)
		}
	}
	helper(nextLevel, result)
	list := []int{}
	for _, v := range level {
		if v != nil {
			list = append(list, v.Val)
		}
	}
	*result = append(*result, list)
}
