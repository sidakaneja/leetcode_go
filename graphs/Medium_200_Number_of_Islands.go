func dfs(grid [][]byte, i, j, m, n int) {
    if i < 0 || i >= m || j < 0 || j >= n {
        return
    }
    if grid[i][j] == '1' {
        grid[i][j] = '0'
        dfs(grid, i - 1, j, m, n)
        dfs(grid, i + 1, j, m, n)
        dfs(grid, i, j + 1, m, n)
        dfs(grid, i, j - 1, m, n)
    }
}

func numIslands(grid [][]byte) int {
    // modifying grid
    m, n := len(grid), len(grid[0])
    count := 0

    for i, row := range grid {
        for j, cell := range row {
            if cell == '1' {
                count++
                dfs(grid, i, j, m, n)
            }
        }
    }

    return count
}
