func minimumSteps(s string) int64 {
  left := len(s) -1 
  for left > 0 && s[left] == '1' {
    left --
  }
  a := 0 
  for i := left - 1 ; i >= 0  ; i-- {
    if s[i] == '1' {
      a += pos(left - i)
      left --
    }
  }

  return int64(a)
}
func pos (i int) int {
  if i < 0 {
    return 0
  }
  return i
}