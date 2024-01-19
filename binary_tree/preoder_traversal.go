/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	} else {
		result := []int{root.Val}
		left := preorderTraversal(root.Left)
		right := preorderTraversal(root.Right)

		result = append(result, left...)
		result = append(result, right...)

		return result
	}
}
