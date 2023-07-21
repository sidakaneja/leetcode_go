func firstBadVersion(n int) int {
    lo, hi := 0, n
    
    for lo < hi {
        mid := lo + ((hi - lo) / 2)
        if isBadVersion(mid) {
            hi = mid
        } else {
            lo = mid + 1
        }
    }
    return lo
}
