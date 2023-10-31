func findAllConcatenatedWordsInADict(words []string) []string {
  m := map[string]bool{}
  for _, word := range words{
    m[word] = true
  }
  k := []string{}
  
  for _, word := range words {
    dp := map[string]int{}
    if re(word, m, 0, dp) > 0{
      k = append(k, word)
    }
  } 
  return k
}
func re(word string, words map[string]bool, depth int, dp map[string]int) int {
  if words[word] && depth > 0{
    return depth 
  }
  if val, ok := dp[word]; ok {

    return val
  }

  for i := 0; i < len(word)  ; i ++ {
    if re(word[:i], words, depth + 1, dp) > 0 && re(word[i:], words, depth + 1, dp) > 0{
      dp[word] = depth + 1
      return depth + 1
    }
  }
  dp[word] = -1
  return -1
}