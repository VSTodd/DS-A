/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	return helper(root, 0)
}

func helper(root *TreeNode, total int) int {
	total = total*10 + root.Val

	if root.Left == nil && root.Right == nil {
		return total
	} else if root.Left == nil {
		return helper(root.Right, total)
	} else if root.Right == nil {
		return helper(root.Left, total)
	}

	return helper(root.Left, total) + helper(root.Right, total)
}

