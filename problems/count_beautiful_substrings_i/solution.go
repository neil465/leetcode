func beautifulSubstrings(s string, k int) int {
  curv := make([]int, len(s))
  curc := make([]int, len(s))
  v := 0
  for i := range s {
    if strings.Contains("aeiou", string(s[i])) {
      curv[i] = 1
    } else {
      curc[i] = 1
    }
    if i != 0 {
      curv[i] += curv[i - 1]
      curc[i] += curc[i - 1]
    } 
    
    

  }
  for i := range s {
    subv, subc := curv[i] , curc[i]
      
    if subv == subc && ( subv * subc ) % k == 0  {

      v ++
    }
    for j := 0 ; j < i; j ++ {
      subv, subc = curv[i] - curv[j], curc[i] - curc[j]
      
      if subv == subc && ( subv * subc ) % k == 0  {

        v ++
      }
    }
  }
  
  return v
}