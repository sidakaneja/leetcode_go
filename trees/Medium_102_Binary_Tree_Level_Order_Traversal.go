func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    curLevel, result := []*TreeNode{root}, [][]int{}

    for len(curLevel) > 0 {
        nextLevel := []*TreeNode{}
        curValues := make([]int, len(curLevel))

        for i, node := range curLevel {
            curValues[i] = node.Val
            
            if node.Left != nil {
                nextLevel = append(nextLevel, node.Left)
            }
            if node.Right != nil {
                nextLevel = append(nextLevel, node.Right)
            }
        }
        
        curLevel = nextLevel
        result = append(result, curValues)
    }

    return result
}
