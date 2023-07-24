func climbStairs(n int) int {
    // f(n) = f(n - 1) + f(n - 2)
    a, b := 1, 1

    for i := 1; i <= n; i++ {
        a, b = b, a + b
    }

    return a
}
