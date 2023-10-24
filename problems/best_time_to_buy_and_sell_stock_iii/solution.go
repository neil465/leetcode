func maxProfit(prices []int) int {
    var dp = make([]int, len(prices))
    pnt := 0
    res := 0
    for i := 1; i < len(prices); i ++ {
      if prices[i] - prices[pnt] < 0 {
        pnt = i
      }
      res = max(res, prices[i] - prices[pnt])
      dp[i] = res
    }
    sellPtr := len(prices) -1 
    res2 := 0 
    for i := len(prices) - 1; i > 0 ; i -- {
      if prices[sellPtr] - prices[i]  < 0 {
        sellPtr = i
      } else {
        res2 = max(res2, prices[sellPtr] - prices[i] + dp[i - 1] ,dp[i])
      }
    }
    return res2
}
func max(a ...int) int {
  m := math.MinInt32
  for _, v := range a {
    if v > m {
      m = v
    }
  }
  return m
}