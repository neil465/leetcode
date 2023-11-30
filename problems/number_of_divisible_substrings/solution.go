var m = map[string]int{
  "a" : 1, "b" : 1,
  "c" : 2, "d" : 2, "e" : 2,
  "f" : 3, "g" : 3, "h" : 3,
  "i" : 4, "j" : 4, "k" : 4,
  "l" : 5, "m" : 5, "n" : 5,
  "o" : 6, "p" : 6, "q" : 6,
  "r" : 7, "s" : 7, "t" : 7,
  "u" : 8, "v" : 8, "w" : 8,
  "x" : 9, "y" : 9, "z" : 9,

}


func countDivisibleSubstrings(word string) int {
  pre := make([]int, len(word) )
  cur := 0 
  for i := range word {
    cur += m[string(word[i])]
    pre[i] = cur
  }

  a := 0 
  for i := 0 ; i < len(word); i++ {
    if pre[i] % (i + 1) == 0 {
      a++
    }
    for j := 0 ; j < i ; j++ {
      if (pre[i] - pre[j]) % (i - j) == 0 {
        a++
      }
    }
  }
  return a
}