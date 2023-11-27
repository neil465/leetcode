func findWordsContaining(words []string, x byte) []int {
  arr := []int{}
  for i := range words {
    if strings.Contains(words[i], string(x)) {
      arr = append(arr, i)
    }
  }
  return arr
}