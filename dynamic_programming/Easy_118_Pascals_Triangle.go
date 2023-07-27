func generate(numRows int) [][]int {
    result := [][]int{
        {1},
    }

    for i := 2; i <= numRows; i++ {
        row := make([]int, i)
        row[0], row[i-1] = 1, 1

        for j := 1; j < i - 1; j++ {
            row[j] = result[i-2][j-1] + result[i-2][j]
        }
        result = append(result, row)
    }
    return result
}
