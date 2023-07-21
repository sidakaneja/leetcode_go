func longestPalindrome(s string) int {
    counts := make(map[rune] int)
    foundOdd := 0
    result := 0

    for _, c := range s {
        counts[c]++
    }

    for _, value := range counts {
        result += (value / 2) * 2
        if value % 2 == 1 {
            foundOdd = 1
        }
    }

    return result + foundOdd
}
