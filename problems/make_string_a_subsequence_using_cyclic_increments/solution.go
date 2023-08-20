func canMakeSubsequence(str1 string, str2 string) bool {
    str1p, str2p := 0, 0
    for str1p < len(str1) && str2p < len(str2) {
      if (int(str2[str2p])- int(str1[str1p])  > -1 && int(str2[str2p])- int(str1[str1p]) <= 1) || (int(str2[str2p]) == 'a' && str1[str1p] == 'z'){
        str2p ++
      } 
      str1p ++
    }
    return str2p == len(str2) 
}
func abs(i int) int {
  if i > 0 {
    return i
  }
  return -i
}