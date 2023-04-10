func hasGroupsSizeX(deck []int) bool {
    m := map[int]int{}
    for _, i := range deck {
        m[i]++
    }
    a := []int{}
    for _, i := range m {
        a = append(a, i)
    }
    r := a[0]
    for i := 1 ; i < len(a); i ++ {
        r = gcd(a[i], r)
        if r == 1 {
            return false
        }
    }
    return r != 1
}
func gcd(a, b int) int{
  if (a == 0) {
      return b
  }
  return gcd(b % a, a)
}