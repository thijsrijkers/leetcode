# Intuition
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
<!-- Describe your approach to solving the problem. -->

# Complexity
- Time complexity:
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity:
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```golang []
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }

    numberOfNodes := 1
    goToNextNode(root.Left, &numberOfNodes)
    goToNextNode(root.Right, &numberOfNodes)
    return numberOfNodes
}

func goToNextNode(node *TreeNode, numberOfNodes *int) {
    if node == nil {
        return
    }
    *numberOfNodes++
    goToNextNode(node.Left, numberOfNodes)
    goToNextNode(node.Right, numberOfNodes)
}

```