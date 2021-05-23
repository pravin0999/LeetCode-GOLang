// This code is the solution to the LeetCode problem - https://leetcode.com/problems/path-sum/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
    if (root == nil){
        return false
    }

    return getSum(root, targetSum);

}

func getSum(root *TreeNode, targetSum int) bool {
    if(root == nil) {
        return false;
    }
    if (root.Left == nil && root.Right == nil){
        return 0 == targetSum - root.Val;
    }

    return getSum(root.Left, targetSum - root.Val) || getSum(root.Right, targetSum - root.Val);
}

