func countSubstrings(s string) int {
    res := len(s)
    chocies := [][]int{}
    for i := 0 ; i < len(s); i ++ {
      chocies = append(chocies, []int{i, i})
    }
    sk := []byte(s)
    visited := map[[2]int]bool{}
    for len(chocies) > 0 {
      nChocies := [][]int{}

      for _, choice := range chocies {
        
        if choice[0] - 1 >= 0 && isPalindrome(sk[choice[0] - 1:  choice[1] + 1]) && !visited[[2]int{choice[0] - 1,  choice[1] }]{
          res ++
          nChocies = append(nChocies, []int{ choice[0] - 1,  choice[1] })
          visited[[2]int{choice[0] - 1,  choice[1] }] = true
        }
        if choice[1] + 1 < len(s) && isPalindrome(sk[choice[0]:  choice[1] + 2]) && !visited[[2]int{choice[0] ,  choice[1] + 1}]{
          res ++
          nChocies = append(nChocies, []int{choice[0] ,  choice[1] +1})
          visited[[2]int{choice[0] ,  choice[1] + 1}] = true
        }
        if choice[1] + 1 < len(s) &&  choice[0] - 1 >= 0 && isPalindrome(sk[choice[0] - 1:  choice[1] + 2]) && ! visited[[2]int{choice[0] - 1,  choice[1] + 1 }]{
          res ++
          nChocies = append(nChocies, []int{ choice[0] - 1,  choice[1] +1})
          visited[[2]int{choice[0] - 1,  choice[1] + 1 }] = true
        }
      }
      chocies = nChocies
      
    }
    return res
}
func isPalindrome(a []byte) bool {
  for i, j := 0, len(a) - 1; i < j ; i, j = i + 1, j - 1 {
    if a[i] != a[j] {
      return false
    }
  }
  return true
}