func letterCasePermutation(s string) []string {
    return help(0, s)
}
func help(i int, s string) []string {
    if i == len(s) {
        return []string{s}
    }
    if unicode.IsLetter(rune(s[i])) {
        b := []rune(s)
        b[i] = rune(strings.ToUpper(string(b[i]))[0])
        a := help(i + 1, string(b))
        b[i] = rune(strings.ToLower(string(b[i]))[0])
        a = append(a, help(i + 1, string(b))...)
        return a
    }
    return help(i + 1, s)
}