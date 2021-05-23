// https://leetcode.com/problems/maximum-depth-of-binary-tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    count := 0
    return getDepth(root, count)

}

func getDepth(root *TreeNode,count int) int {
    if root == nil {
        return count
    }

    return max(getDepth(root.Left, count+1), getDepth(root.Right, count+1))
}

func max(a int, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

