package leetcode103

/*
	Challenge Description: Given the root of a binary tree, return the zigzag
	level order traversal of its nodes' values. (i.e., from left to right, then
	right to left for the next level and alternate between).
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) (result [][]int) {
	if root == nil {
		return result
	}
	q := []*TreeNode{root}
	ltr := true //left to right
	for len(q) > 0 {
		levelVals := make([]int, len(q))
		for i := range levelVals {
			node := q[0]
			q = q[1:]

			if ltr {
				levelVals[i] = node.Val
			} else {
				k := len(levelVals) - 1
				levelVals[k-i] = node.Val
			}

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		ltr = !ltr
		result = append(result, levelVals)
	}
	return result
}
