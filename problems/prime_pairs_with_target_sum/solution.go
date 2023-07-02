func findPrimePairs(n int) [][]int {
    z := [][]int{}
    m := se(n)
    for i := range m {
        if m[n - i] && n - i >= i {
            z = append(z, []int{i, n - i})

        }
    }
    sort.Slice(z, func(i, j int) bool {
        return z[i][0] < z[j][0]
    })
    return z
    
}
func se(n int) map[int]bool {
  integers := make([]bool, n+1)
  for i := 2; i < n+1; i++ {
      integers[i] = true
  }

  for p := 2; p*p <= n; p++ {
      if integers[p] == true {
          for i := p * 2; i <= n; i += p {
              integers[i] = false
          }
      }
  }

    primes := map[int] bool {}
  for p := 2; p <= n; p++ {
      if integers[p] == true {
          primes[p] = true
      }
  }

  return primes
}
