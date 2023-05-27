func buyChoco(prices []int, money int) int {
    sort.Ints(prices)
    x := prices[0] + prices[1]
    if money - x >= 0 {
        return money - x
    }
    return money
}