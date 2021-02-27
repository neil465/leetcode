func lengthOfLastWord(s string) int {
    s = strings.TrimRight(s," ")
    h := strings.Split(s, " ")
	return len(string(h[len(h)-1]))
}