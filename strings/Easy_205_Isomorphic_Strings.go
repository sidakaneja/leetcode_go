func isIsomorphic(s string, t string) bool {
    mapsTo, mappedBy := make(map[byte] int), make(map[byte] int)

    for i := range s {
        p, q := s[i], t[i]

        if mapsTo[p] != mappedBy[q]{
            return false
        }

        // Need + 1 otherwise not present in maps is the same as element at index 0 due to zero values
        mapsTo[p] = i + 1
        mappedBy[q] = i + 1
    }

    return true
}
