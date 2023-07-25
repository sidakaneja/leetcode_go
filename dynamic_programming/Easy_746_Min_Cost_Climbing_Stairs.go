// O(n) space
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func minCostClimbingStairs(cost []int) int {
    minCosts := append([]int{}, cost...)
    n := len(cost)

    for i := 2; i < n; i++ {
        minCosts[i] += min(minCosts[i-1], minCosts[i-2])
    }

    return min(minCosts[n-1], minCosts[n-2])
}

// Constant space

func minCostClimbingStairs(cost []int) int {
    n := len(cost)
    low, high := cost[0], cost[1]

    for i := 2; i < n; i++ {
        low, high = high, min(low, high) + cost[i]
    }
    return min(low, high)
}
