func areSimilar(mat [][]int, k int) bool {
    for _, row := range mat {
      nr := append(row[k % len(row):], row[:k % len(row)]...)
      for j := range row {
        if nr[j] != row[j] {
          return false
        }
      }
    }
    return true
}