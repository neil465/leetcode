func maxProfit(prices []int) int {
    pointer := 0 
    max := 0
    for i := 1 ; i < len(prices) ; i++{
        if prices[i]-prices[pointer] < 0 {
            pointer = i
            continue
        }
        max += prices[i]-prices[pointer]
        pointer = i

    }
    return max
}