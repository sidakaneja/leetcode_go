// Memoization
func helper(costGrid [][]int, i, j int) int {
    if i < 0 || j < 0 {
        return 0
    } else if costGrid[i][j] == 0 {
        costGrid[i][j] = helper(costGrid, i-1, j) + helper(costGrid, i, j-1)
    }
    return costGrid[i][j]
}
func uniquePaths(m int, n int) int {
    costGrid := make([][]int, m)
    for i := range costGrid {costGrid[i] = make([]int, n)}
    costGrid[0][0] = 1

    return helper(costGrid, m-1, n-1)
}

// Dynamic programming
func uniquePaths(m int, n int) int {
    costGrid := make([][]int, m)
    for i := range costGrid {costGrid[i] = make([]int, n)}
    costGrid[0][0] = 1

    for i, row := range costGrid {
        for j, _ := range row {
            if i == 0 || j == 0 {
                costGrid[i][j] = 1
            } else {
                costGrid[i][j] = costGrid[i-1][j] + costGrid[i][j-1]
            }
        }
    }
    return costGrid[m-1][n-1]
}

// O(n) space instead of O(m * n)
func uniquePaths(m int, n int) int {
    prev, cur := make([]int, n), make([]int, n)
    prev[0], cur[0] = 1, 1

    for i := 0; i < m; i++ {
        for j := 1; j < n; j++ {
            cur[j] = cur[j-1] + prev[j]
        }
        prev, cur = cur, prev
    }
    return prev[n-1]
}
