func finalPrices(prices []int) []int {
    ms := []int{0}
    for i := 1 ; i < len(prices) ; i++{
        
        for len(ms) > 0 && prices[i] <= prices[ms[len(ms) - 1]] {
            prices[ms[len(ms) - 1]] = prices[ms[len(ms) - 1]] - prices[i]
            ms = ms[:len(ms) - 1]
        }

        ms = append(ms, i)
    }
    return prices
}