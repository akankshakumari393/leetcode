func maxProfit(prices []int) int {
    minPrice := math.MaxInt32
    maxProfit := 0
    for _, price := range prices {
        minPrice = min(minPrice, price)
        maxProfit = max(maxProfit, (price - minPrice))
    }
    return maxProfit
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}