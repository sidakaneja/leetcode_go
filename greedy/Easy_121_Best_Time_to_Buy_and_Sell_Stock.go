func maxProfit(prices []int) int {
    maxProfit, start := 0, prices[0]

    for _, price := range prices {
        if (price - start) > maxProfit {
            maxProfit = price - start
        }
        if price < start {
            start = price
        }
    }
    return maxProfit
}
