package leetcode94

/*
	Challenge Description: Given the root of a binary tree, return the inorder
	traversal of its nodes' values.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Iterative Solution
func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{}
	curr := root

	for curr != nil || len(stack) != 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		last := len(stack) - 1
		curr, stack = stack[last], stack[:last]
		result = append(result, curr.Val)
		curr = curr.Right
	}
	return result
}

/*
// Recursive Solution
func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	dfs(root, &result)
	return result
}

func dfs(root *TreeNode, result *[]int) {
	if root != nil {
		dfs(root.Left, result)
		*result = append(*result, root.Val)
		dfs(root.Right, result)
	}
}
*/
