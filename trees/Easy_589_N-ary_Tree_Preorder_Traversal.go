// Iterative
func preorder(root *Node) []int {
    if root == nil {
        return []int{}
    }

    stack, result := []*Node{root}, []int{}

    for len(stack) > 0 {
        l := len(stack)

        node := stack[l - 1]
        stack = stack[: l - 1]

        result = append(result, node.Val)

        for i := len(node.Children) - 1; i >= 0; i--  {
            stack = append(stack, node.Children[i])
        }
    }

    return result
}

// Recursive
func preorderHelper(root *Node, result *[]int){
    if root == nil {
        return
    }

    *result = append(*result, root.Val)
    for _, child := range root.Children {
        preorderHelper(child, result)
    }
}

func preorder(root *Node) []int {
    result := []int{}
    preorderHelper(root, &result)
    return result
}
