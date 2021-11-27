func numberOfSubstrings(s string) int64 {
    count := int64(len(s))
    m := make(map[rune]int)
    for _,i := range s {if m[i] > 0{count += int64(m[i])}
        m[i]++}
    return count
}