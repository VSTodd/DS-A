/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	} else {
		left := inorderTraversal(root.Left)
		right := inorderTraversal(root.Right)

		left = append(left, root.Val)
		left = append(left, right...)

		return left
	}
}
