package leetcode102

/*
	Challenge Description: Given the root of a binary tree, return the level
	order traversal of its nodes' values. (i.e., from left to right, level by
	level).
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BFS solution
// 520 ns/op
func levelOrder(root *TreeNode) (result [][]int) {
	if root == nil {
		return result
	}

	// Create a queue
	q := []*TreeNode{root}

	// Iterate until queue is fully consumed
	for len(q) > 0 {

		// This zero-valued slice that will hold this level's values
		levelVals := make([]int, len(q))

		for i := range levelVals {

			// Pop the front of the queue (FIFO-style)
			node := q[0]
			q = q[1:]

			// Replace zero-value
			levelVals[i] = node.Val

			// If there are children then add them to the queue
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		result = append(result, levelVals)
	}
	return result
}

// Recursive DFS Solution
// 471.9 ns/op
/*
func levelOrder(root *TreeNode) (results [][]int) {
	if root != nil {
		helper(root, 0, &results)
	}
	return results
}

func helper(node *TreeNode, level int, results *[][]int) {
	if len(*results) == level {
		*results = append(*results, []int{})
	}

	(*results)[level] = append((*results)[level], node.Val)

	if node.Left != nil {
		helper(node.Left, level+1, results)
	}
	if node.Right != nil {
		helper(node.Right, level+1, results)
	}
}
*/
