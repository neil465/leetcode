func maxLengthBetweenEqualCharacters(s string) int {
    m := map[byte]int{}
    max := -1
    for i := range s {
        if m[s[i]] == 0{
            m[s[i]] = i + 1
        } else if max < i - m[s[i]] {
            max = i - m[s[i]]
        }
    }
    return max
}