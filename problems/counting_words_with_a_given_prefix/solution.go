func prefixCount(words []string, pref string) int {
    c := 0
    for _, i := range words {
      if strings.HasPrefix(i, pref) {
        c++
      }
    }
    return c
}