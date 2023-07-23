func flood(image [][]int, sr, sc, newColor, oldColor int) {
    m, n := len(image), len(image[0])

    if sr < 0 || sr >= m || sc < 0 || sc >= n || oldColor == newColor {
        return
    }
    if image[sr][sc] == oldColor  {
        image[sr][sc] = newColor
        flood(image, sr - 1, sc, newColor, oldColor)
        flood(image, sr + 1, sc, newColor, oldColor)
        flood(image, sr, sc - 1, newColor, oldColor)
        flood(image, sr, sc + 1, newColor, oldColor)
    }
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    flood(image, sr, sc, color, image[sr][sc])
    return image
}
