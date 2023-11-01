

func minCut(s string) int {
    return re( s, 0, 0, map[[2]int]int{}) - 1
}
func re( s string, i, j int, dp map[[2]int]int) int {
  if j > len(s) {
    if i >= len(s) || isPalindrome(s[i:len(s)]) {
      return 0
    }
    return math.MaxInt32
  }
  if val, ok := dp[[2]int{i, j}]; ok {
    return val
  }
  if !isPalindrome(s[i:j]) { 
    return re(s, i, j + 1, dp) 
  } 
  v, v2 := re(s, j, j + 1, dp) + 1 , math.MaxInt32
  if j - i != 0 {
    v2 = re(s, i, j + 1 , dp) 
  }
  if v2 < v {
    dp[[2]int{i, j}] = v2
    return v2 
  } else {
    dp[[2]int{i, j}] = v
  }

  return v 
}
func isPalindrome(s string) bool {
  for i, j := 0, len(s) -1 ; i < j; i, j = i + 1, j - 1{
    if s[i] != s[j] {
      return false
    }
  }
  return true
}
func max(i, j int) int {
  if i > j {
    return i
  }
  return j
}