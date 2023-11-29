func bestClosingTime(customers string) int {
  missed := make([]int, len(customers) + 1)
  rsum := 0 
  for i := len(customers) -1 ; i >= 0 ; i -- {
    if customers[i] == 'Y' {
      rsum ++
    }
    missed[i] = rsum
  }

  acur := 0 
  if customers[0] == 'N'  {
      acur ++
  }
  minInd, minsum := -1, missed[0] + 0
  if missed[1] + acur < minsum {
    minsum = missed[1] + acur
    minInd ++
  }
  cur := 0
  for i := 0 ; i < len(customers) ; i ++ {
    if customers[i] == 'N'  {
      cur ++
    }
    if missed[i + 1] + cur < minsum {
      minsum = missed[i + 1] + cur
      minInd = i
    }
  }

  return minInd + 1
}