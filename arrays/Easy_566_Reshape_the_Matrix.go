func matrixReshape(mat [][]int, r int, c int) [][]int {
    m, n, count := len(mat), len(mat[0]), 0
    
    if m * n != r * c {
        return mat
    }

    result := make([][]int, r)
    for i := 0; i < r; i++ {result[i] = make([]int, c)}

    for _, row := range mat {
        for _, cell := range row {
            result[count / c][count % c] = cell
            // could use count = (i * n) + j
            count++
        }
    }
    return result
}
