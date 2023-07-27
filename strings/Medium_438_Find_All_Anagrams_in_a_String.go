func findAnagrams(s string, p string) []int {
    var pCounts, sCounts [26]int
    for _, c := range p {pCounts[c - 'a']++}

    result := []int{}

    for i, c := range s {
        sCounts[c - 'a']++
        startIndex := i - len(p)

        if startIndex >= 0 {
            startChar := s[startIndex]
            sCounts[startChar - 'a']--
        }
        if pCounts == sCounts {
            result = append(result, startIndex + 1)
        }
    }

    return result
}
