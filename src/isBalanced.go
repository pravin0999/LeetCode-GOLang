// https://leetcode.com/problems/balanced-binary-tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {

    if root == nil { return true }

    dummyStack := []*TreeNode{root}

    for len(dummyStack) > 0 {
        node := dummyStack[(len(dummyStack)-1)]
        dummyStack = dummyStack[:(len(dummyStack)-1)]
        if node.Left != nil { dummyStack = append(dummyStack, node.Left) }
        if node.Right != nil { dummyStack = append(dummyStack, node.Right) }

        if abs(getDepth(node.Left) - getDepth(node.Right)) > 1 {
            return false
        }
    }
    return true
}

func getDepth(node *TreeNode) int {
    if node == nil {
        return 0
    }
    return 1 + max(getDepth(node.Left), getDepth(node.Right))
}

func max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

