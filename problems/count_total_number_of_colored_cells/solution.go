func coloredCells(n int) int64 {
  if n == 1 {
    return int64(1)
  }
  b := int64(1)
  for i := 1; i < n ; i ++ {
    b += int64(4 * (i))
  }
  return b
}