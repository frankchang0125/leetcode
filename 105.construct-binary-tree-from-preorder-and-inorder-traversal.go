/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
 *
 * https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
 *
 * algorithms
 * Medium (38.95%)
 * Likes:    1668
 * Dislikes: 46
 * Total Accepted:    220.5K
 * Total Submissions: 538.9K
 * Testcase Example:  '[3,9,20,15,7]\n[9,3,15,20,7]'
 *
 * Given preorder and inorder traversal of a tree, construct the binary tree.
 * 
 * Note:
 * You may assume that duplicates do not exist in the tree.
 * 
 * For example, given
 * 
 * 
 * preorder = [3,9,20,15,7]
 * inorder = [9,3,15,20,7]
 * 
 * Return the following binary tree:
 * 
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// * The first node of Preorder Traversal is the root of binary tree.
// * The nodes in the left part of root node in Inorder array must in the
//   left subtree of the root node.
// * The nodes in the right part of root node in Inorder array must in the
//   right subtree of the root node.
//
// ex: preorder = [3,9,20,15,7]
//     inorder = [9,3,15,20,7]
//     * The root node of binary tree must be 3.
//     * [9] must in the left subtree of root node 9.
//     * [29, 15, 7] must in the right subtree of root node 9.
//
// Reference: http://bit.ly/2LRcqLS
func buildTree(preorder []int, inorder []int) *TreeNode {
    pLen := len(preorder)
    iLen := len(inorder)

    if pLen == 0 || iLen == 0 || pLen != iLen {
        return nil
    }

    return constructTree(preorder, inorder, 0, len(preorder) - 1, 0, len(inorder) - 1)
}

func constructTree(preorder []int, inorder []int,
    preStart, preEnd, inStart, inEnd int) *TreeNode {
    if preStart > preEnd || inStart > inEnd {
        return nil
    }

    root := &TreeNode{
        Val: preorder[preStart],
    }

    // Find root node index: i in inorder array.
    i := 0
    for _, num := range inorder {
        if num == root.Val {
            break
        }
        i++
    }

    // (i - inStart): The number of nodes in left subtree.
    //  => We can then build left subtree and right subtree
    //     by their preorder arrays.
    // ex: preorder = [1,2,4,5,3,7,6,8]
    //     inorder = [4,2,5,1,6,7,3,8]
    //                  1
    //              /       \
    //             2         3
    //           /   \     /   \
    //          4     5   7     8
    //                   /
    //                  6
    //     Root node = 1, num nodes in left subtree = 3.
    //      => Left subtree's preorder array: [2,4,5]
    //      => Right subtree's preorder array: [3,7,6,8]
    numLeftNodes := i - inStart

    // Left subtree.
    root.Left = constructTree(preorder, inorder,
        preStart+1, preStart+numLeftNodes,
        inStart, i-1)

    // Right subtree.
    root.Right = constructTree(preorder, inorder,
        preStart+numLeftNodes+1, preEnd,
        i+1, inEnd)

    return root
}
