// Recursive
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if p.Val < root.Val && q.Val < root.Val {
        return lowestCommonAncestor(root.Left, p, q)
    } else if p.Val > root.Val && q.Val > root.Val{
        return lowestCommonAncestor(root.Right, p, q)
    } else {
        return root
    }
}

// Iterative
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    node := root
    
    for node != nil {
        if p.Val < node.Val && q.Val < node.Val {
            node = node.Left
        } else if p.Val > node.Val && q.Val > node.Val{
            node = node.Right
        } else {
            return node
        }
    }

    return node
}
