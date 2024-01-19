/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	traversed := preorderTraversal(root)
	currentNode := root
	for i := 0; i < len(traversed)-1; i++ {
		currentNode.Left = nil
		var newNode TreeNode
		newNode.Val = traversed[i+1]
		currentNode.Right = &newNode
		currentNode = &newNode
	}
}

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