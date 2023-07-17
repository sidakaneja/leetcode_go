func isSubsequence(s string, t string) bool {
    found := 0

    for i := 0; i < len(t) && found < len(s); i++ {
        if t[i] == s[found] {
            found++
        }
    }

    return found == len(s)
}
