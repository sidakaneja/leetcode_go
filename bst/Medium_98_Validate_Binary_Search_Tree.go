// Recursive

func isValid(root *TreeNode, min, max int) bool {
    if root == nil {
        return true
    }
    return root.Val >= min && root.Val <= max && isValid(root.Left, min, root.Val - 1) && isValid(root.Right, root.Val + 1, max)
}

func isValidBST(root *TreeNode) bool {
    return isValid(root, math.MinInt, math.MaxInt)
}

// Iterative

func isValidBST(root *TreeNode) bool {
    stack := []*TreeNode{}
    var prev *TreeNode

    for len(stack) > 0  || root != nil {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        n := len(stack)
        root = stack[n - 1]
        stack = stack[:n - 1]

        if prev != nil && root.Val <= prev.Val {
            return false
        }

        prev = root
        root = root.Right
    }

    return true
}
