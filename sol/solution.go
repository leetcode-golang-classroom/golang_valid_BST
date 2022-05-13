package sol

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return CheckBST(root, math.MaxFloat64, -math.MaxFloat64)
}

func CheckBST(root *TreeNode, upperBound float64, lowerBound float64) bool {
	if root == nil {
		return true
	}
	cur := float64(root.Val)
	if cur <= lowerBound || cur >= upperBound {
		return false
	}
	return CheckBST(root.Left, cur, lowerBound) && CheckBST(root.Right, upperBound, cur)
}
