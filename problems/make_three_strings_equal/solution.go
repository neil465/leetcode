func findMinimumOperations(s1 string, s2 string, s3 string) int {

  for i := 0 ; i < min(min(len(s1), len(s2)), len(s3)); i++ {
    if s1[i] != s2[i] || s2[i] != s3[i] || s3[i] != s1[i] {
      if i == 0 {
        return -1
      }
      return len(s1) - i + len(s2) - i + len(s3) - i
    }
  }
  return len(s1) - min(min(len(s1), len(s2)), len(s3)) + len(s2) - min(min(len(s1), len(s2)), len(s3)) + len(s3) - min(min(len(s1), len(s2)), len(s3))
}