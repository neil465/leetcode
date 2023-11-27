func maximizeSquareHoleArea(n int, m int, hBars []int, vBars []int) int {
    sort.Ints(hBars)
    sort.Ints(vBars)

    maxH := 1
    curH := 1
    for i := 1; i < len(hBars); i++ {
      if hBars[i] == n + 2 {
        break
      }
      if hBars[i] == hBars[i - 1] + 1{
        curH ++
      } else {
        curH = 1
      }
      maxH = max(curH, maxH)
    }
    maxV := 1
    curV := 1

    for i := 1; i < len(vBars); i++ {
      if vBars[i] == m + 2 {
        break
      }
      if vBars[i] == vBars[i - 1] + 1 {
        curV ++
      } else {
        curV = 1
      }
      maxV = max(maxV, curV)
    }
    return (min(maxV, maxH) + 1) * (min(maxV, maxH) + 1)

}